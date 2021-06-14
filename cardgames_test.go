package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var deck_id string

//var newdk newdecks

type newdecks struct {
	deck_id string `json:"deck_id"`
	newdeck []newp `json:"newdecks"`
}
type newp struct {
	game      string `json:"game"`
	shuffle   string `json:"shuffle"`
	query     string `json:"query"`
	remaining int    `json:"remaining"`
	maxcards  int    `json:"maxcards"`
	shuffled  bool   `json:"shuffled"`
	Cards     []Card `json:"cards"`
}

type deckid struct {
	deck_id string `json:"deck_id"`
}
type ids struct {
	did []deckid `json:"decksids"`
}

var ds newdecks //Decks GET petitions for tests (TDD)
var dis ids     //Deck ids list (capture by petitions)

func preparetests() {

	var np newp
	np.game = "poker"
	np.shuffle = "withoutshuffle"
	np.query = ""
	np.remaining = 52
	np.maxcards = 52
	np.shuffled = false
	ds.newdeck = append(ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "withoutshuffle"
	np.query = "?cards=AS,KD,AC,2C,KH"
	np.remaining = 5
	np.maxcards = 52
	np.shuffled = false
	ds.newdeck = append(ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "shuffle"
	np.query = ""
	np.remaining = 52
	np.maxcards = 52
	np.shuffled = true
	ds.newdeck = append(ds.newdeck, np)

	np.game = "poker"
	np.shuffle = "shuffle"
	np.remaining = 5
	np.maxcards = 52
	np.shuffled = true
	np.query = "?cards=AS,KD,AC,2C,KH"
	ds.newdeck = append(ds.newdeck, np)

}

func init() {

	//Prepare TDD
	preparetests()

	go SetupServer()                   //Start the HTTP server, concurrence mode
	time.Sleep(200 * time.Millisecond) //wait to start server

}

//********************* TEST NEW DECK ***********************
func TestNewDeck(t *testing.T) {

	assert := assert.New(t)

	for _, d := range ds.newdeck {

		if resp, err := http.Get(
			"http://127.0.0.1:8000/api/deck/" + d.game + "/new/" + d.shuffle + d.query); err == nil {

			assert.Equal(http.StatusCreated, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)

			bodyString := string(bodyBytes)
			log.Info(bodyString)

			var deck Deck
			json.Unmarshal(bodyBytes, &deck)

			assert.Equal(true, deck.Success)
			assert.Equal("", deck.Txterror)
			assert.Equal(d.remaining, deck.Remaining)
			assert.Equal(d.shuffled, deck.Shuffled)
			assert.NotEmptyf(deck.Deck_id, "", "Deck_id is Empty.")

			var newid deckid
			newid.deck_id = deck.Deck_id
			dis.did = append(dis.did, newid)

		} else {
			assert.Fail(err.Error())
		}
	}

}

//***************** TEST OPEN DECK ***********************
func TestOpenDeck(t *testing.T) {

	assert := assert.New(t)

	for i := 0; i < len(dis.did); i++ {
		if resp, err := http.Get(
			"http://127.0.0.1:8000/api/deck/" + dis.did[i].deck_id); err == nil {

			assert.Equal(http.StatusOK, resp.StatusCode)
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			log.Info(bodyString)
			var deck Deck
			json.Unmarshal(bodyBytes, &deck)

			assert.Equal(true, deck.Success)
			assert.Equal("", deck.Txterror)
			assert.Equal(ds.newdeck[i].remaining, deck.Remaining)
			assert.Equal(ds.newdeck[i].shuffled, deck.Shuffled)
			assert.Equal(dis.did[i].deck_id, deck.Deck_id)
			assert.Equal(ds.newdeck[i].remaining, len(Opendecks.Decks[i].Cards))

			log.Print(Opendecks.Decks[i])

			// -------- Check all deck --------
			var odeck Opendeck
			filename := fmt.Sprintf("%s%s", ds.newdeck[i].game, ".json")
			jsonFile, err := os.Open(filename)
			// defer the closing of our jsonFile so that we can parse it later on
			defer jsonFile.Close()
			// if we os.Open returns an error then handle it
			if err != nil {
				assert.Fail("File game not found:" + filename)
			} else {
				log.Print("Successfully Opened: " + filename)
				// read our opened xmlFile as a byte array.
				byteValue, _ := ioutil.ReadAll(jsonFile)

				// we unmarshal our byteArray which contains our
				// jsonFile's content into 'users' which we defined above
				json.Unmarshal(byteValue, &odeck)
			}
			if ds.newdeck[i].maxcards == Opendecks.Decks[i].Remaining {
				if deck.Shuffled {
					//Shuffle
					assert.NotEqual(odeck.Cards, Opendecks.Decks[i].Cards)

				} else {
					//Without Shuffle, default order
					assert.Equal(odeck.Cards, Opendecks.Decks[i].Cards)
				}
			}
		} else {
			assert.Fail(err.Error())
		}
	}

}

//***************** TEST DRAW DECK **************************
func TestDrawDeck(t *testing.T) {

	assert := assert.New(t)

	countdraw := 6
	for i := 0; i < len(dis.did); i++ {
		if resp, err := http.Get(
			"http://127.0.0.1:8000/api/deck/" + dis.did[i].deck_id + "/draw?count=" + strconv.Itoa(countdraw)); err == nil {

			assert.Equalf(http.StatusOK, resp.StatusCode, "Status is not Ok")

			bodyBytes, _ := io.ReadAll(resp.Body)

			bodyString := string(bodyBytes)
			log.Info(bodyString)
			var deck Deck
			json.Unmarshal(bodyBytes, &deck)

			if deck.Success {
				assert.Equal(true, deck.Success)
				assert.Equal("", deck.Txterror)
				assert.GreaterOrEqual(deck.Remaining, 0)
				assert.Equal((ds.newdeck[i].remaining - countdraw), len(Opendecks.Decks[i].Cards))
			} else {
				assert.Equal(false, deck.Success)
				assert.NotEmpty(deck.Txterror)
			}
			log.Info(Opendecks.Decks[i].Cards)

		} else {
			assert.Fail(err.Error())
		}
	}

}
