package main

import (
	"fmt"
	"net/http"

	handler "handlers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("This is My chat")

	server := mux.NewRouter()
	server.HandleFunc("/login", handler.Login).Methods("POST")
	server.HandleFunc("/getchat", handler.Chat).Methods("GET")
	server.HandleFunc("/postchat", handler.TalkChat).Methods("POST")
	server.HandleFunc("/allchat", handler.AllChat).Methods("GET")
	server.HandleFunc("/signup", handler.SignUp).Methods("POST")
	http.ListenAndServe(":8089", server)
}
