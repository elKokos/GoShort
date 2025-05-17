package main

import (
	"GoShort/configs"
	"GoShort/internal/auth"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthConfig{conf})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("server is running on port:", server.Addr)
	log.Fatal(server.ListenAndServe())
}
