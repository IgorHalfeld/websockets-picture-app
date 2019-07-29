package main

import (
	"fmt"
	"github.com/igorhalfeld/websocketsapi/helpers"
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

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
