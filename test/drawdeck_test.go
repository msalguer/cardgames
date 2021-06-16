package cardgames

import (
	"encoding/json"
	c "main/cardgames"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

//***************** TEST DRAW DECK **************************
func TestDrawDeck(t *testing.T) {

	assert := assert.New(t)

	countdraw := 6
	z := 0
	//for i := 0; i < len(Dis.did); i++ {
	for _, di := range Dis.did {

		/*
			if resp, err := http.Get(

				//"http://127.0.0.1:8000/api/deck/" + dis.did[i].deck_id + "/draw?count=" + strconv.Itoa(countdraw)); err == nil {
				"/api/deck/" + Dis.did[i].deck_id + "/draw?count=" + strconv.Itoa(countdraw)); err == nil {

		*/

		//req, err := http.NewRequest("GET", "/api/deck/"+Dis.did[i].deck_id+"/draw?count="+strconv.Itoa(countdraw), nil)
		req, err := http.NewRequest("GET", "/api/deck/"+di.deck_id+"/draw?count="+strconv.Itoa(countdraw), nil)
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(c.DrawDeck)
		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(res, req)

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		//assert.Equalf(http.StatusOK, resp.StatusCode, "Status is not Ok")
		assert.Equalf(http.StatusOK, res.Code, "Status is not Ok")

		//bodyBytes, _ := ioutil.ReadAll(resp.Body)
		//bodyString := string(bodyBytes)

		bodyBytes := res.Body.Bytes()
		bodyString := res.Body.String()
		log.Info(bodyString)
		var deck c.Deck
		json.Unmarshal(bodyBytes, &deck)
		//json.Unmarshal([]byte(bodyString), &deck)

		//fmt.Println("I:", strconv.Itoa(i))
		if deck.Success {
			assert.Equal(true, deck.Success)
			assert.Equal("", deck.Txterror)
			assert.GreaterOrEqual(deck.Remaining, 0)
			//assert.Equal((Ds.newdeck[i].remaining), len(c.Opendecks.Decks[i].Cards))
			//assert.Equal((Ds.newdeck[z].remaining), len(c.Opendecks.Decks[z].Cards))
		} else {
			assert.Equal(false, deck.Success)
			assert.NotEmpty(deck.Txterror)
		}

		z++
		//log.Info(c.Opendecks.Decks[i].Cards)

		/*		} else {
				assert.Fail(err.Error())
			} */
	}

}
