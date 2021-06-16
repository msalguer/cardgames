package cardgames

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//************************** DRAW DECK *******************************************
//func DrawDeck(c echo.Context, res http.ResponseWriter, req *http.Request) { ///api/deck/:deck_id/draw"
func DrawDeck(res http.ResponseWriter, req *http.Request) { ///api/deck/:deck_id/draw"
	//	func DrawDeck(c echo, res http.ResponseWriter, req *http.Request) error {
	//defer req.Body.Close()
	fmt.Println("Endpoint Hit: drawDeck")
	//deck_id := c.AcquireContext().Param("deck_id") // Param("deck_id")
	//deck_id := c.Param("deck_id")
	//count, err := strconv.Atoi(c.QueryParam("count"))
	urlPart := strings.Split(req.URL.Path, "/")
	deck_id := urlPart[3]
	count, err := strconv.Atoi(req.URL.Query().Get("count"))
	print(count)

	var retopendeck Opendeck

	if err != nil || deck_id == "" {
		retopendeck.Success = false
		retopendeck.Txterror = "Deck not found, or the card counter parameter does not have a value greater than 0 (or does not exist)"
		//return c.JSON(http.StatusOK, retopendeck)
		json.NewEncoder(res).Encode(retopendeck)
	}

	enc := false

	var newopendecks OpenDecks
	var newopendeck Opendeck

	for _, opendeck := range Opendecks.Decks {

		if opendeck.Deck_id == deck_id {
			enc = true
			retopendeck.Deck_id = opendeck.Deck_id

			if count > opendeck.Remaining {
				retopendeck.Success = false
				retopendeck.Txterror = "The card to draw counter is greater than the number of available cards"
				retopendeck.Remaining = opendeck.Remaining
				//return c.JSON(http.StatusOK, retopendeck)
				json.NewEncoder(res).Encode(retopendeck)
			} else {
				retopendeck.Remaining = opendeck.Remaining - count
				retopendeck.Success = true

				for j := 0; j < count; j++ {
					//for j := 0; j < count; j++ {
					retopendeck.Cards = append(retopendeck.Cards, opendeck.Cards[j])
				}

				//Add new opendeck (replace/modify opendecks)
				newopendeck.Deck_id = opendeck.Deck_id
				newopendeck.Remaining = retopendeck.Remaining
				newopendeck.Success = true
				for i := count; i < len(opendeck.Cards); i++ {
					newopendeck.Cards = append(newopendeck.Cards, opendeck.Cards[i])
				}
				newopendecks.Decks = append(newopendecks.Decks, newopendeck)

			}

		} else {
			//Add rest opendecks

			newopendecks.Decks = append(newopendecks.Decks, opendeck)
		}

	}

	if enc {
		//Replace opendecks modify
		Opendecks = newopendecks
	} else {
		retopendeck.Success = false
		retopendeck.Txterror = "Deck not found"
	}
	//return c.JSON(http.StatusOK, retopendeck)
	json.NewEncoder(res).Encode(retopendeck)
}
