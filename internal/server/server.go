package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	LOG  *log.Logger
	HTTP *http.Server
}

func NewServer(log *log.Logger) Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleEndpoint)
	mux.HandleFunc("/upload", handlers.HandleUpload)
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return Server{LOG: log, HTTP:httpServer}
}
