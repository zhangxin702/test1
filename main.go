package main

import (
	"fmt"
	"net/http"
	"gotest/pkg/service"
)

func main() {
	server := service.NewService()
	endpoints := service.NewEndPointServer(server)
	httpHandler := service.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
