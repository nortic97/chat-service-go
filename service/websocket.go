package service

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var client []websocket.Conn

func HandlerMessage(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)
	client = append(client, *conn)
	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Show message in log
		log.Printf("%s sent: %s", conn.RemoteAddr(), string(msg))

		for _, client := range client {
			// Write message back to browser
			if err = client.WriteMessage(messageType, msg); err != nil {
				return
			}
		}

	}
}
