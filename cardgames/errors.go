package cardgames

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ShowError(res http.ResponseWriter, txt string) {
	fmt.Println("Game parameter not found")
	var deck Deck
	deck.Success = false
	deck.Txterror = "Game parameter not found"
	json.NewEncoder(res).Encode(deck)
}
