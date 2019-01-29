package main

type WebSocket struct {
	conns map[*WebSocketConn]bool
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		conns: make(map[*WebSocketConn]bool),
	}
}

func (ws *WebSocket) Register(wsc *WebSocketConn) {
	ws.conns[wsc] = true

}

func (ws *WebSocket) Unregister(wsc *WebSocketConn) {
	delete(ws.conns, wsc)
}

func (ws *WebSocket) Broadcast(name string, data interface{}, onErr func(*WebSocketConn, error)) {
	for conn := range ws.conns {
		if err := conn.Send(name, data); err != nil && onErr != nil {
			onErr(conn, err)
		}
	}
}
