// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
// }

// func main() {
//     r := gin.Default()

//     r.GET("/ws", func(c *gin.Context) {
//         // Upgrade the HTTP connection to a WebSocket connection
//         conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//         if err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         defer conn.Close()

//         for {
//             messageType, p, err := conn.ReadMessage()
//             if err != nil {
//                 return
//             }
//             if err := conn.WriteMessage(messageType, p); err != nil {
//                 return
//             }
//         }
//     })

//     r.Run(":8080")
// }
