package main

// WebSocket contains a map with all registered
// web socket connections. A map with a bool value
// is used that values can be deleted from the map
// on unregistration. The boolean value is only an
// unused dummy value.
type WebSocket struct {
	conns map[*WebSocketConn]bool
}

// NewWebSocket creates a new instacne of web socket
// and initilaizes the ws connections map.
func NewWebSocket() *WebSocket {
	return &WebSocket{
		conns: make(map[*WebSocketConn]bool),
	}
}

// Register registers a web socket connection in the
// ws connections list of the manager.
func (ws *WebSocket) Register(wsc *WebSocketConn) {
	ws.conns[wsc] = true
}

// Unregister removes a web socket connection from the
// ws connections list of the manager.
func (ws *WebSocket) Unregister(wsc *WebSocketConn) {
	delete(ws.conns, wsc)
}

// Broadcast sends an event with the passed name and data
// to all registered web socket connections. If an error
// occures, the onErr function will be called and will
// get passed the ws connection and the error object.
func (ws *WebSocket) Broadcast(name string, data interface{}, onErr func(*WebSocketConn, error)) {
	for conn := range ws.conns {
		if err := conn.Send(name, data); err != nil && onErr != nil {
			onErr(conn, err)
		}
	}
}
