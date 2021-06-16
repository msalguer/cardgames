package main

import (
	"fmt"
	"log"
	"main/cardgames"
	"net/http"

	"github.com/gorilla/mux"
)

/*
import (
	c "main/cardgames"
)
*/

func main() {
	//cardgames.SetupServer()

	//http.HandleFunc("/", NewDeck) //utils.go
	//E.GET("/api/deck/:game/new/:shuffle", NewDeck(E.AcquireContext().Response().Writer, E.AcquireContext().Request().Response.Request)) //With cards parameter
	//E.GET("/api/deck/:game/new/:shuffle", NewDeck)
	//E.GET("/api/deck/:deck_id", OpenDeck)
	//E.GET("/api/deck/:deck_id/draw", DrawDeck) //With count parameter
	//E.GET("/a", myHandler)
	//.HandleFunc("/status", myHandler)
	//E.Router().Add("/api/deck/:game/new/:shuffle", NewDeck(E.AcquireContext().Response().Writer, E.AcquireContext().Request().Response.Request))

	//Routes
	r := mux.NewRouter()
	r.HandleFunc("/api", cardgames.HomePage).Methods("GET")
	r.HandleFunc("/api/deck/{game}/new/{shuffle}", cardgames.NewDeck).Methods("GET") //(E.NewContext(E.AcquireContext()))
	r.HandleFunc("/api/deck/{deck_id}", cardgames.OpenDeck).Methods("GET")
	r.HandleFunc("/api/deck/{deck_id}/draw", cardgames.DrawDeck).Methods("GET")

	/*
		http.HandleFunc("/api/deck/{game}/new/{shuffle}", cardgames.NewDeck) //(E.NewContext(E.AcquireContext()))
		http.HandleFunc("/api/deck/{deck_id}/draw", cardgames.DrawDeck)
		http.HandleFunc("/api/deck/{deck_id}", cardgames.OpenDeck)
		http.HandleFunc("/api", cardgames.HomePage)
	*/

	fmt.Println("Server listen on 8000 port")
	//http.ListenAndServe(":8000", nil)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

/*
func Cardgames() {
	//c.Loadreadme()
	c.SetupServer()
}
*/
