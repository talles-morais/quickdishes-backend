package services

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexões de qualquer origem. Ajuste para produção.
	},
}

var clients = make(map[*websocket.Conn]bool) // Conexões ativas
var broadcast = make(chan interface{})       // Canal para mensagens
var mutex = sync.Mutex{}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
	}
}

func StartBroadcast() {
	for {
		message := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
