// utils/websocket.go
package utils

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketClient struct {
	Conn *websocket.Conn
	Send chan []byte
}

var clients = make(map[*WebSocketClient]bool)
var broadcast = make(chan []byte)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Upgrade error: %v", err)
	}
	defer ws.Close()

	client := &WebSocketClient{Conn: ws, Send: make(chan []byte)}
	clients[client] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, client)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			client.Send <- msg
		}
	}
}

func StartWebSocketServer() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
