package cardgames

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ShowError(res http.ResponseWriter, txt string) {
	//func ShowError(c echo.Context, txt string) {
	fmt.Println("Game parameter not found")
	var deck Deck
	deck.Success = false
	deck.Txterror = "Game parameter not found"
	//c.JSON(http.StatusOK, deck)
	//json.NewEncoder(c).Encode(deck)

	//req.Header().Set("Content-Type", "application/json")
	//req.Write(json.Marshal(deck))

	//res.Header.Set()  Set("Content-Type", "application/json")
	//res.Header.Write(http.StatusCreated)
	json.NewEncoder(res).Encode(deck)
}
