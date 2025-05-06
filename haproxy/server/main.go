package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/ws/01", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("query", r.URL.Query())
		err := m.HandleRequest(w, r)
		if err != nil {
			log.Println(err)
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("Client connected:", s.Keys) // 可以查看 Session 的 Keys，預設是空的
		m.Broadcast([]byte("New client connected!"))
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Client disconnected:", s.Keys)
		m.Broadcast([]byte("Client disconnected."))
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("Received message from client: %s\n", msg)
		// 將收到的訊息廣播給所有其他連接的客戶端 (包括發送者)
		m.Broadcast(msg)
	})

	fmt.Println("Melody WebSocket server is starting to listen on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
