package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("websocket upgrade err", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read messsage error,", err)
			break
		}
		log.Printf("received: %s", msg)

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("Writer error:", err)
			break
		}
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ws", handleWebsocket)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "ws.html", gin.H{})
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "hello",
			"date": time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	log.Println("websocket server run at :9002")
	r.Run(":9002")
}
