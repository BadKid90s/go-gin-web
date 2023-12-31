package main

import (
	"fmt"
	"go-gin-demo/cmd/server/wire"
	"go-gin-demo/pkg/config"
	"go-gin-demo/pkg/http"
	"go-gin-demo/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	logger.Info("Starting Server")

	handler := wire.NewServer(conf, logger)

	err := http.Run(handler, fmt.Sprintf(":%d", conf.GetInt("server.port")))
	if err != nil {
		logger.Error("Server starting failed")
	}
}
