package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	router := gin.New()

	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		server.OnEvent("/control", "video-control", func(s socketio.Conn, msg string) {
			server.BroadcastToNamespace("/control", "execute", msg)
		})
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/video", http.Dir("./asset"))
	if err := router.Run(":8888"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
