package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type client struct {
	conn *websocket.Conn
	send chan []byte
}

type chatRoom struct {
	clients    map[*client]bool
	register   chan *client
	unregister chan *client
	broadcast  chan []byte
	mutex      sync.Mutex
}

func (cr *chatRoom) run() {
	for {
		select {
		case client := <-cr.register:
			cr.mutex.Lock()
			cr.clients[client] = true
			cr.mutex.Unlock()
		case client := <-cr.unregister:
			cr.mutex.Lock()
			delete(cr.clients, client)
			close(client.send)
			cr.mutex.Unlock()
		case message := <-cr.broadcast:
			cr.mutex.Lock()
			for client := range cr.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(cr.clients, client)
				}
			}
			cr.mutex.Unlock()
		}
	}
}

func (c *client) readPump(chat *chatRoom) {
	defer func() {
		chat.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		chat.broadcast <- message
	}
}

func (c *client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &client{conn: conn, send: make(chan []byte, 256)}
	
}

func main() {

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handleChat(w, r)
	})

	fmt.Println("WebSocket chat server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}