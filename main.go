package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/my-route", mainPage)

	port := ":8000"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type Message struct {
	Message string `json:"msg"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	message := Message{"Hello World"}
	js,_ := json.Marshal(message)
    w.Write(js)
}