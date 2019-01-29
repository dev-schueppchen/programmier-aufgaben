package main

import (
	"net/http"
)

type HTTPServer struct {
	ws *WebSocket
}

func StartHTTPServer(port string) {
	httpServer := &HTTPServer{
		ws: NewWebSocket(),
	}

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsc, err := NewWebSocketConn(w, r, func(wsc *WebSocketConn, reason error) {
			httpServer.ws.Unregister(wsc)
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wsc.On("message", func(e *Event) {
			msg, err := NewMessageFromEvent(e)
			if err != nil {
				LogDbg.Println("[ WS ] Could not parse ws message as message object: ", err)
				return
			}

			httpServer.ws.Broadcast("message", msg, func(wsc *WebSocketConn, err error) {
				LogErr.Printf("[ WS ] Failed sending message event to %s: %v\n", wsc.conn.LocalAddr().String(), err)
			})
		})

		httpServer.ws.Register(wsc)
	})

	http.ListenAndServe(":"+port, nil)
}
