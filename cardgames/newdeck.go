package cardgames

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//************************** CREATE DECK *******************************************
func NewDeck(res http.ResponseWriter, req *http.Request) { //"/api/deck/"+d.game+"/new/"+d.shuffle+d.query

	urlPart := strings.Split(req.URL.Path, "/")

	fmt.Println("Endpoint Hit: createDeck")
	game := urlPart[3]
	shuffle := urlPart[5]
	cards := req.URL.Query().Get("cards")

	print(cards)

	shuffled := false
	if string(shuffle) == "shuffle" || shuffle == "1" {
		shuffled = true
	} else {
		shuffled = false
	}

	filename := fmt.Sprintf("%s%s", game, ".json")
	byteValue := ReadDataFile(res, filename)
	fmt.Println("Successfully Opened pocker.json")

	// we initialize our Users array
	var deck Deck
	var opendeck Opendeck
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &opendeck)

	//Filter by cards
	if len(cards) > 0 {
		var filterdeck Opendeck
		cards := strings.Split(cards, ",")
		for _, onecard := range opendeck.Cards {
			for _, card := range cards {
				fmt.Println(card, onecard.Code)
				if card == onecard.Code {
					var newcard Card
					newcard.Code = onecard.Code
					newcard.Suit = onecard.Suit
					newcard.Value = onecard.Value
					filterdeck.Cards = append(filterdeck.Cards, newcard)
				}
			}
		}

		opendeck = filterdeck
	}

	if shuffled == true {
		rand.Seed(time.Now().UnixNano())
		for i := len(opendeck.Cards) - 1; i > 0; i-- { // Fisher–Yates shuffle
			j := rand.Intn(i + 1)
			opendeck.Cards[i], opendeck.Cards[j] = opendeck.Cards[j], opendeck.Cards[i]
		}
	}

	opendeck.Deck_id = GenerateSecureToken(MinSecretKeySize)
	opendeck.Shuffled = shuffled
	opendeck.Remaining = len(opendeck.Cards)
	opendeck.Success = true

	deck.Deck_id = opendeck.Deck_id
	deck.Shuffled = opendeck.Shuffled
	deck.Remaining = opendeck.Remaining
	deck.Success = opendeck.Success

	Opendecks.Decks = append(Opendecks.Decks, opendeck)

	json.NewEncoder(res).Encode(deck)

}
