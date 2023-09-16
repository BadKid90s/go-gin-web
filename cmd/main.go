package main

import (
	"fmt"
	"go-gin-demo/cmd/wire"
	"go-gin-demo/pkg/config"
	"go-gin-demo/pkg/http"
	"log"
)

func main() {
	conf := config.NewConfig()

	handler := wire.NewServer(conf)
	if handler == nil {
		log.Panicln("Starting Server Failed")
	}

	http.Run(handler, fmt.Sprintf(":%d", conf.GetInt("server.port")))
}
