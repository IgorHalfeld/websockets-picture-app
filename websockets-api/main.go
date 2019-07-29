package main

import (
	"fmt"
	"github.com/igorhalfeld/websockets-api/helpers"
	"log"
	"net/http"
	"os"
)

func main() {
	hub := helpers.NewHub()
	go hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "It's works!")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		helpers.ServeWs(hub, w, r)
	})


	var port string
	if os.Getenv("PORT") == "" {
		port = "3000"
	} else {
		port = os.Getenv("PORT")
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
