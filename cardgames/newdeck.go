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
//func NewDeck(c echo.Context) error {
func NewDeck(res http.ResponseWriter, req *http.Request) { //"/api/deck/"+d.game+"/new/"+d.shuffle+d.query
	//func NewDeck(c echo.Context) error {
	//func NewDeck(res http.ResponseWriter, req *http.Request) {

	urlPart := strings.Split(req.URL.Path, "/")

	fmt.Println("Endpoint Hit: createDeck")
	game := urlPart[3]
	shuffle := urlPart[5]
	cards := req.URL.Query().Get("cards")

	/*
		game := c.Param("game") //  c.Param("game")
		shuffle := c.Param("shuffle")
		cards := c.QueryParam("cards")
	*/

	//game := GetParam(c, "game")
	//shuffle := GetParam(c, "shuffle")
	//cards := GetParam(res, req, "cards")
	/*
		game := GetParam(res, req, "game")
		shuffle := GetParam(res, req, "shuffle")
		cards := GetParam(res, req, "cards")
	*/
	/*
		if cards == "" {
			ShowError(res, req, "cards parameter not found")
		}
	*/
	/*
		fmt.Println("Game parameter not found")
		var deck Deck
		deck.Success = false
		deck.Txterror = "Game parameter not found"
		//return c.JSON(http.StatusOK, deck)
		//req.Header().Set("Content-Type", "application/json")
		//req.Write(json.Marshal(deck))

		//res.Header.Set()  Set("Content-Type", "application/json")
		//res.Header.Write(http.StatusCreated)
		json.NewEncoder(res).Encode(deck)
	*/

	print(cards)

	shuffled := false
	if string(shuffle) == "shuffle" || shuffle == "1" {
		shuffled = true
	} else {
		shuffled = false
	}

	//filename := path.Dir(fmt.Sprintf("%s%s", "../data/"+game, ".json"))
	//testf := filepath.Abs(filename)
	//filename := fmt.Sprintf("%s%s", "./../data/"+game, ".json")
	//filename := fmt.Sprintf("%s%s%s", os.Getwd(), "./../data/"+game, ".json")
	//filename := os.Getwd + "/data/" + game + ".json"
	//filename, err := filepath.Abs(fmt.Sprintf("%s%s", "data/"+game, ".json"))
	filename := fmt.Sprintf("%s%s", game, ".json")
	//f := filepath.Join()
	/*
		//Actual file an directory
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("No caller information")
		}
		fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	*/
	/*
		jsonFile, err := os.Open(filename)
		//jsonFile, err := ioutil.ReadFile(filename)
		//byteValue, err := ioutil.ReadFile(filename)

		// if we os.Open returns an error then handle it
		if err != nil {
			ShowError(res, "Game not found")
			os.Exit(0)
		}

	*/
	byteValue := ReadDataFile(res, filename)
	fmt.Println("Successfully Opened pocker.json")
	// defer the closing of our jsonFile so that we can parse it later on
	//defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	//byteValue, _ := ioutil.ReadAll(jsonFile)

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

	//return c.JSON(http.StatusCreated, deck)
	//json.NewEncoder(c.Echo().AcquireContext().Response().Writer).Encode(deck)
	json.NewEncoder(res).Encode(deck)

}
