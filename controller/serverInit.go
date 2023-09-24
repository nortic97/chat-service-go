package controller

import (
	"chat-server/service"
	"log"
	"net/http"
)

func ServerInit() {
	http.HandleFunc("/reverse", service.HandlerMessage)
	log.Print("Server started in: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
