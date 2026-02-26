package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	http   *http.Server
}

func StartServer(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HTML)
	router.HandleFunc("/upload", handlers.Upload)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	return &Server{
		logger: logger,
		http:   httpServer,
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}
