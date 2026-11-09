package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	loger := log.New(os.Stdout, "server: ", log.LstdFlags)
	srv := server.NewServer(loger)
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		loger.Fatalf("failed to start server: %v", err)
	}
}
