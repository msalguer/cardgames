package cardgames

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//************************** OPEN DECK *******************************************
func OpenDeck(res http.ResponseWriter, req *http.Request) { ///api/deck/:deck_id"
	fmt.Println("Endpoint Hit: openDeck")
	urlPart := strings.Split(req.URL.Path, "/")
	deck_id := urlPart[3]
	print(deck_id)

	var retopendeck Opendeck
	enc := false
	for _, opendeck := range Opendecks.Decks {
		if opendeck.Deck_id == deck_id {
			retopendeck = opendeck
			enc = true
			break
		}
	}
	if !enc {
		var opendeck Opendeck
		opendeck.Success = false
		opendeck.Txterror = "Deck not found"
		retopendeck = opendeck
	}

	json.NewEncoder(res).Encode(retopendeck)
}
