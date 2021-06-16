package cardgames

import (
	"encoding/json"
	c "main/cardgames"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

//********************* TEST NEW DECK ***********************
func TestNewDeck(t *testing.T) {

	assert := assert.New(t)

	for _, d := range Ds.newdeck {

		//if resp, err := http.Get(
		//if resp, err := http.NewRequest(
		//"http://127.0.0.1:8000/api/deck/" + d.game + "/new/" + d.shuffle + d.query); err == nil {
		//	"/api/deck/"+d.game+"/new/"+d.shuffle+d.query, nil); err == nil {

		req, err := http.NewRequest("GET", "/api/deck/"+d.game+"/new/"+d.shuffle+d.query, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(c.NewDeck)
		//handler := http.HandlerFunc(NewDeck)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(res, req)
		//E.Logger.Fatal(E.Start(":8001"))

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		//res := httptest.NewRecorder()
		//c.NewDeck(res, req)

		/*
			assert.Equal(http.StatusCreated, resp.StatusCode)
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
		*/

		//assert.Equal(http.StatusCreated, res.Code)
		assert.Equal(http.StatusOK, res.Code)
		//bodyBytes, _ := ioutil.ReadAll(handler.ServeHTTP().Body) //   req.Body)
		bodyBytes := res.Body.Bytes()
		//bodyString := string(bodyBytes)
		bodyString := res.Body.String()
		log.Info(bodyString)

		var deck c.Deck
		json.Unmarshal(bodyBytes, &deck)

		assert.Equal(true, deck.Success)
		assert.Equal("", deck.Txterror)
		assert.Equal(d.remaining, deck.Remaining)
		assert.Equal(d.shuffled, deck.Shuffled)
		assert.NotEmptyf(deck.Deck_id, "", "Deck_id is Empty.")

		var newid Deckid
		newid.deck_id = deck.Deck_id
		Dis.did = append(Dis.did, newid)

		/*
			} else {
				assert.Fail(err.Error())
			}
		*/
	}

}
