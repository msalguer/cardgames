package main

import (
	"fmt"
	"log"
	"main/cardgames"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//Routes
	r := mux.NewRouter()
	r.HandleFunc("/api", cardgames.HomePage).Methods("GET")
	r.HandleFunc("/api/deck/{game}/new/{shuffle}", cardgames.NewDeck).Methods("GET")
	r.HandleFunc("/api/deck/{deck_id}", cardgames.OpenDeck).Methods("GET")
	r.HandleFunc("/api/deck/{deck_id}/draw", cardgames.DrawDeck).Methods("GET")

	fmt.Println("Server listen on 8000 port")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
