package websocket_Room

import "github.com/gorilla/websocket"

func ShutdownAllWebSocketClients() {
	mutex.Lock()
	defer mutex.Unlock()
	for _, room := range game_rooms {
		for _, client := range room.clients {
			client.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
				websocket.CloseServiceRestart,
				"Server is shutting down",
			))
			client.conn.Close()
		}
		room.clients = make(map[string]*Client)
	}
	conn_clients = make(map[*Client]bool)
}
