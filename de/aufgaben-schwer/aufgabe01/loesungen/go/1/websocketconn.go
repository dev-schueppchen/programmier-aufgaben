package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// The upgrader will upgrade a HTTP connection
// to a web socket connection by using the
// set configuration.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketConn contains the actual connection,
// a channel used to send messages out, a map
// with all registered event handlers by name and
// a function, which will be called, when the web
// socket connection was closed.
type WebSocketConn struct {
	conn     *websocket.Conn
	out      chan []byte
	events   map[string]EventHandler
	onClosed func(*WebSocketConn, error)
}

// NewWebSocketConn creates a new instance of WebSocketConn. By passing the ResponseWriter and the
// Reuquest, the HTTP connection will be upgraded to a web socket connection by the upgrader.
// Also, the reader and writer loops are initialized in two seperate go routines.
func NewWebSocketConn(w http.ResponseWriter, r *http.Request, onClosed func(*WebSocketConn, error)) (*WebSocketConn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	wsc := &WebSocketConn{
		conn:     conn,
		out:      make(chan []byte),
		events:   make(map[string]EventHandler),
		onClosed: onClosed,
	}

	// Reader and writer needs to be started in
	// seperate go routines ("threads"), because they
	// block their current thread until receiving
	// or writing a message.
	go wsc.reader()
	go wsc.writer()

	return wsc, nil
}

// The reader initializes a loop, which will block
// until a web socket message was received. The message
// wil be parsed to an event and if the event name is
// registered in the event handler map, the appertaining
// handler function will be called and the event will be
// passed to this function.
func (wsc *WebSocketConn) reader() {
	defer wsc.conn.Close()

	// Creating "infinite" loop.
	for {
		// This blocks until ReadMessage() receives a message
		// on the web socket input.
		_, msg, err := wsc.conn.ReadMessage()
		if err != nil {
			// If any close message was send, the onClosed handler will be called, if defined
			// and the loop will break.
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				LogInf.Println("[ WS ] closed: ", err)
				if wsc.onClosed != nil {
					wsc.onClosed(wsc, err)
				}
			}
			break
		}

		// Creating message event object from raw message data.
		event, err := NewEventFromRaw(msg)
		if err != nil {
			LogErr.Println("[ WS ] Error parsing received message: ", err)
			continue
		}

		LogDbg.Printf("[ WS ] Received message: {name: %s, data: %v}\n", event.Name, event.Data)

		// Searching for event registration of the event name
		// and execute the corresponding handler, if defined.
		if handler, ok := wsc.events[event.Name]; ok {
			handler(event)
		}
	}
}

// The writer initializes a loop, which will block
// until a new message enters the out channel. After,
// the message will be send to the connected client.
func (wsc *WebSocketConn) writer() {
	// Creating "infinite" loop.
	for {
		// Block and listen for incomming messages
		// in web socket connections out channel,
		// which are dedicated to be sent to the
		// web socket client.
		msg, ok := <-wsc.out
		if !ok {
			wsc.conn.WriteMessage(websocket.CloseMessage, nil)
			return
		}

		// Getting the writer and specify the message type
		// as text message.
		w, err := wsc.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			LogErr.Println("[ WS ] Failed getting next writer for sending message: ", err)
			return
		}

		// Write the message data to the writer and close after.
		w.Write(msg)
		w.Close()
	}
}

// On registers a new event handler function for an event name and
// returns a function which will unregister the event after execution.
func (wsc *WebSocketConn) On(name string, handler EventHandler) func() {
	wsc.events[name] = handler
	return func() {
		delete(wsc.events, name)
	}
}

// Send creates an event object from the passed name and data,
// encodes the event to JSON formatted raw data, which will be
// send to the connections out channel after.
func (wsc *WebSocketConn) Send(name string, data interface{}) error {
	event := &Event{
		Name: name,
		Data: data,
	}

	// Get raw JSON data from event object.
	raw, err := event.Raw()
	if err != nil {
		return err
	}

	// Send data into web socket connections out channel.
	wsc.out <- raw

	return nil
}
