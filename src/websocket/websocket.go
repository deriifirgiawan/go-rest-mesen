package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var mutex = sync.Mutex{}

func HandleWebSocket(context *gin.Context) {
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Println("Failed Connect to WebSocket: ", err)
		return
	}

	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, _, err := conn.ReadMessage()

		if err != nil {
			log.Println("Lost Connection:", err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
	}
}