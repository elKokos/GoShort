package main

import (
	"GoShort/configs"
	"GoShort/internal/auth"
	"GoShort/internal/link"
	"GoShort/pkg/db"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := configs.LoadConfig()
	db, _ := db.New(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)

	//auth
	auth.NewAuthHandler(router, auth.AuthConfig{conf})

	//link
	link.NewLinkHandler(router, link.LinkConfig{LinkRepository: linkRepository})

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
