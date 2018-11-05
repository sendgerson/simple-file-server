package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("dir")))
	log.Println("Listening on 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
