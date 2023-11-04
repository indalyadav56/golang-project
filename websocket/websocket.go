package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

var pool = sync.Pool{
    New: func() interface{} {
        return &WsConnection{
            conn: nil,
        }
    },
}

type WsConnection struct {
    conn *websocket.Conn
}

func Websocket(c *gin.Context) {
    upgrader.CheckOrigin= func(*http.Request) bool { return true}
    conn := pool.Get().(*WsConnection)
    defer pool.Put(conn)
    
    var err error
    conn.conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }

    // Use connection
    conn.conn.WriteMessage(1, []byte("ping"))
}