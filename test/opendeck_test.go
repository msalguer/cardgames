package cardgames

import (
	"encoding/json"
	"fmt"
	c "main/cardgames"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

//***************** TEST OPEN DECK ***********************
func TestOpenDeck(t *testing.T) {

	assert := assert.New(t)

	for i := 0; i < len(Dis.did); i++ {
		/*
			if resp, err := http.Get(
				//"http://127.0.0.1:8000/api/deck/" + dis.did[i].deck_id); err == nil {
				"/api/deck/" + dis.did[i].deck_id); err == nil {
		*/

		req, err := http.NewRequest("GET", "/api/deck/"+Dis.did[i].deck_id, nil)
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(c.OpenDeck)
		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(res, req)
		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		//assert.Equal(http.StatusOK, resp.StatusCode)
		assert.Equal(http.StatusOK, res.Code) //Duplicate
		//bodyBytes, _ := ioutil.ReadAll(resp.Body)
		//bodyString := string(bodyBytes)
		bodyBytes := res.Body.Bytes()
		bodyString := res.Body.String()
		log.Info(bodyString)
		var deck c.Deck
		json.Unmarshal(bodyBytes, &deck)

		assert.Equal(true, deck.Success)
		assert.Equal("", deck.Txterror)
		assert.Equal(Ds.newdeck[i].remaining, deck.Remaining)
		assert.Equal(Ds.newdeck[i].shuffled, deck.Shuffled)
		assert.Equal(Dis.did[i].deck_id, deck.Deck_id)
		assert.Equal(Ds.newdeck[i].remaining, len(c.Opendecks.Decks[i].Cards))

		log.Print(c.Opendecks.Decks[i])

		// -------- Check all deck --------
		var odeck c.Opendeck

		filename := fmt.Sprintf("%s%s", Ds.newdeck[i].game, ".json")
		byteValue := c.ReadDataFile(res, filename)
		json.Unmarshal(byteValue, &odeck)
		/*
			//filename := fmt.Sprintf("%s%s", Ds.newdeck[i].game, ".json")
			filename, err := filepath.Abs(fmt.Sprintf("%s%s", "./../data/"+Ds.newdeck[i].game, ".json"))
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
		*/
		if Ds.newdeck[i].maxcards == c.Opendecks.Decks[i].Remaining {
			if deck.Shuffled {
				//Shuffle
				assert.NotEqual(odeck.Cards, c.Opendecks.Decks[i].Cards)

			} else {
				//Without Shuffle, default order
				assert.Equal(odeck.Cards, c.Opendecks.Decks[i].Cards)
			}
		}
		/*
				} else {
				assert.Fail(err.Error())
			}
		*/
	}

}
