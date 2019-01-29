package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketConn struct {
	conn    *websocket.Conn
	out     chan []byte
	events  map[string]EventHandler
	onClose func(*WebSocketConn, error)
}

func NewWebSocketConn(w http.ResponseWriter, r *http.Request, onClose func(*WebSocketConn, error)) (*WebSocketConn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	wsc := &WebSocketConn{
		conn:    conn,
		out:     make(chan []byte),
		events:  make(map[string]EventHandler),
		onClose: onClose,
	}

	go wsc.reader()
	go wsc.writer()

	return wsc, nil
}

func (wsc *WebSocketConn) reader() {
	defer wsc.conn.Close()

	for {
		_, msg, err := wsc.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				LogInf.Println("[ WS ] closed: ", err)
				if wsc.onClose != nil {
					wsc.onClose(wsc, err)
				}
			}
			break
		}

		event, err := NewEventFromRaw(msg)
		if err != nil {
			LogErr.Println("[ WS ] Error parsing received message: ", err)
			continue
		}

		LogDbg.Printf("[ WS ] Received message: {name: %s, data: %v}\n", event.Name, event.Data)

		if handler, ok := wsc.events[event.Name]; ok {
			handler(event)
		}
	}
}

func (wsc *WebSocketConn) writer() {
	for {
		select {
		case msg, ok := <-wsc.out:
			if !ok {
				wsc.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			w, err := wsc.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				LogErr.Println("[ WS ] Failed getting next writer for sending message: ", err)
				return
			}

			w.Write(msg)
			w.Close()
		}
	}
}

func (wsc *WebSocketConn) On(name string, handler EventHandler) func() {
	wsc.events[name] = handler
	return func() {
		delete(wsc.events, name)
	}
}

func (wsc *WebSocketConn) Send(name string, data interface{}) error {
	event := &Event{
		Name: name,
		Data: data,
	}

	raw, err := event.Raw()
	if err != nil {
		return err
	}

	LogDbg.Println(string(raw))

	go func() {
		wsc.out <- raw
	}()

	return nil
}
