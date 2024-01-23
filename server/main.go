package main

import (
	"gommunity/server/handlers"
	"log"
	"net/http"
)

func main() {
	handler := handlers.HostRouteHandler{}
	handler.RegisterHost("127.0.0.1", handlers.HTTPHandler{})
	if err := http.ListenAndServe(":80", &handler); err != nil {
		log.Fatalln("HTTP 수신 소켓을 초기화하는 과정에서 예외가 발생하였습니다.", err.Error())
	}
}
