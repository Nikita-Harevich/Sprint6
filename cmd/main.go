package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	srv := server.StartServer(logger)
	logger.Println("сервер работает на порту :8080")
	log.Fatal(srv.Start())
}
