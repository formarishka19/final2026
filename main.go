package main

import (
	"log"
	"net/http"

	"final2026/pkg/db"
	"final2026/pkg/server"
)

func main() {
	if db.Init("scheduler.db") != nil {
		log.Println("Error db init, exiting")
		return
	}
	httpLogger := new(log.Logger)
	httpServer := server.NewHttpServer(httpLogger)

	err := httpServer.Srv.ListenAndServe()

	// Проверяем, что это не обычное закрытие сервера (через Shutdown/Close)
	if err != nil && err != http.ErrServerClosed {
		httpLogger.Fatalf("Ошибка сервера: %v", err)
	}
}
