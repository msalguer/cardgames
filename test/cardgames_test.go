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

	//go c.SetupServer()                 //Start the HTTP server, concurrence mode
	//time.Sleep(200 * time.Millisecond) //wait to start server

}

/*
//********* GET URL SERVER *********
func GetHost() string {
	var o os
	s := o.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", u.Host)
	return u.Host
}
*/

/*
func Test_test(t *testing.T) {

	for _, d := range ds.newdeck {

		req, err := http.NewRequest("GET", "/api/deck/"+d.game+"/new/"+d.shuffle+d.query, nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		//c.DrawDeck(E, res, req)
		c.NewDeck(res, req)

		exp := "Hello World"
		act := res.Body.String()
		if exp != act {
			t.Fatalf("Expected %s gog %s", exp, act)
		}

	}

}
*/
