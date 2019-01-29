package main

import (
	"net/http"
)

// HTTPServer contains the webscoket
// manager.
type HTTPServer struct {
	ws *WebSocket
}

// StartHTTPServer opens a new HTTP server and initializes
// the websocket manager and all handlers for the URL endpoints.
func StartHTTPServer(port string) error {
	httpServer := &HTTPServer{
		ws: NewWebSocket(),
	}

	// Expose frontend asset files on root path.
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	// The frontend will request this endpoint to upgrade the HTTP
	// connection to a websocket connection.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// The connection will be upgraded to a WS connection. Also, a handler
		// will be set which unregisters the ws connection after disconnect event.
		wsc, err := NewWebSocketConn(w, r, func(wsc *WebSocketConn, reason error) {
			httpServer.ws.Unregister(wsc)
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// If a "message" event calls, the handler will tyr to create a
		// message object from the event. If this was successful, the message
		// will be broadcasted to all connected and registered clients.
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

		// Now, the ws connection will be registered in the
		// ws manager.
		httpServer.ws.Register(wsc)
	})

	// The HTTP server will be exposed to the passed
	// port number on the local address.
	return http.ListenAndServe(":"+port, nil)
}
