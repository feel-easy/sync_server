package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), s.RemoteAddr(), s.Rooms(), s.Namespace())
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("video-control", msg)
	})

	server.OnEvent("/", "video-control", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("video-control", msg)
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	// connection
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
