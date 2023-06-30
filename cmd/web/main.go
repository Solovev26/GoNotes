package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
