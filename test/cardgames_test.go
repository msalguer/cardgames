package cardgames

import (
	c "main/cardgames"
	"os"
	"testing"
)

var Deck_id string

//var newdk newdecks

type Newdecks struct {
	deck_id string `json:"deck_id"`
	newdeck []Newp `json:"newdecks"`
}
type Newp struct {
	game      string   `json:"game"`
	shuffle   string   `json:"shuffle"`
	query     string   `json:"query"`
	remaining int      `json:"remaining"`
	maxcards  int      `json:"maxcards"`
	shuffled  bool     `json:"shuffled"`
	Cards     []c.Card `json:"cards"`
}

type Deckid struct {
	deck_id string `json:"deck_id"`
}
type Ids struct {
	did []Deckid `json:"decksids"`
}

var Ds Newdecks //Decks GET petitions for tests (TDD)
var Dis Ids     //Deck ids list (capture by petitions)

func preparetests() {

	var np Newp
	np.game = "poker"
	np.shuffle = "withoutshuffle"
	np.query = ""
	np.remaining = 52
	np.maxcards = 52
	np.shuffled = false
	Ds.newdeck = append(Ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "withoutshuffle"
	np.query = "?cards=AS,KD,AC,2C,KH"
	np.remaining = 5
	np.maxcards = 52
	np.shuffled = false
	Ds.newdeck = append(Ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "shuffle"
	np.query = ""
	np.remaining = 52
	np.maxcards = 52
	np.shuffled = true
	Ds.newdeck = append(Ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "shuffle"
	np.remaining = 5
	np.maxcards = 52
	np.shuffled = true
	np.query = "?cards=AS,KD,AC,2C,KH"
	Ds.newdeck = append(Ds.newdeck, np)

}

func init() {

	os.Setenv("TEST", "true")
	//Prepare TDD
	preparetests()
	TestNewDeck(&testing.T{})
	TestOpenDeck(&testing.T{})
	TestDrawDeck(&testing.T{})

}
