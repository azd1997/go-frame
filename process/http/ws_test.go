package http

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
	"time"
)

func TestWs(t *testing.T) {
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		fmt.Println("Conn: ", err)
		return
	}

	go timeWriter(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read: ", err)
			return
		}

		fmt.Printf("Received: %s\n", msg)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(2 *time.Second)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
		conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		conn.WriteMessage(websocket.TextMessage, []byte("server_time"))
	}
}
