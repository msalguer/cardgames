package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// **************************** STRUCTS *************************************

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Deck struct {
	Deck_id   string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Success   bool   `json:"success"`
	Txterror  string `json:"txterror"`
}

type Opendeck struct {
	Deck_id   string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
	Success   bool   `json:"success"`
	Txterror  string `json:"txterror"`
}

type OpenDecks struct {
	Decks []Opendeck `json:"decks"`
}

//********************* GLOBAL VARIABLES AND CONSTANTS *******************************
//Global opendecks
var Opendecks OpenDecks

const minSecretKeySize = 16

// -------------------------- Generate Token --------------------------------------
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

//------- Load readme file ---------
var readme string

func loadreadme() string {
	// SHOW README file

	fptr := flag.String("fpath", "README", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)

	for s.Scan() {
		readme = readme + s.Text() + "\n"
		//fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return readme
}

//************************** API HOME PAGE *******************************************
func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, readme)
}

//************************** CREATE DECK *******************************************
func NewDeck(c echo.Context) error {

	fmt.Println("Endpoint Hit: createDeck")
	game := c.Param("game")
	shuffle := c.Param("shuffle")
	cards := c.QueryParam("cards")
	print(cards)

	shuffled := false
	if shuffle == "shuffle" || shuffle == "1" {
		shuffled = true
	} else {
		shuffled = false
	}

	filename := fmt.Sprintf("%s%s", game, ".json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Game not found")
		var deck Deck
		deck.Success = false
		deck.Txterror = "Game not found"
		return c.JSON(http.StatusOK, deck)
	}

	fmt.Println("Successfully Opened pocker.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

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

	opendeck.Deck_id = GenerateSecureToken(minSecretKeySize)
	opendeck.Shuffled = shuffled
	opendeck.Remaining = len(opendeck.Cards)
	opendeck.Success = true

	deck.Deck_id = opendeck.Deck_id
	deck.Shuffled = opendeck.Shuffled
	deck.Remaining = opendeck.Remaining
	deck.Success = opendeck.Success

	Opendecks.Decks = append(Opendecks.Decks, opendeck)

	return c.JSON(http.StatusCreated, deck)

}

//************************** OPEN DECK *******************************************
func OpenDeck(c echo.Context) error {
	fmt.Println("Endpoint Hit: openDeck")
	deck_id := c.Param("deck_id")
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
	return c.JSON(http.StatusOK, retopendeck)
}

//************************** DRAW DECK *******************************************
func DrawDeck(c echo.Context) error {
	fmt.Println("Endpoint Hit: drawDeck")
	deck_id := c.Param("deck_id")
	count, err := strconv.Atoi(c.QueryParam("count"))
	print(count)

	var retopendeck Opendeck

	if err != nil || deck_id == "" {
		retopendeck.Success = false
		retopendeck.Txterror = "Deck not found, or the card counter parameter does not have a value greater than 0 (or does not exist)"
		return c.JSON(http.StatusOK, retopendeck)
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
				return c.JSON(http.StatusOK, retopendeck)
			} else {
				retopendeck.Remaining = opendeck.Remaining - count
				retopendeck.Success = true

			}

			for i := 0; i < count; i++ {
				retopendeck.Cards = append(retopendeck.Cards, opendeck.Cards[i])
			}

			//Add new opendeck (replace/modify opendecks)
			newopendeck.Deck_id = opendeck.Deck_id
			newopendeck.Remaining = retopendeck.Remaining
			newopendeck.Success = true
			for i := count; i < len(opendeck.Cards); i++ {
				newopendeck.Cards = append(newopendeck.Cards, opendeck.Cards[i])
			}
			newopendecks.Decks = append(newopendecks.Decks, newopendeck)

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
	return c.JSON(http.StatusOK, retopendeck)
}

//********************* SETUP AND RUN SERVER *******************
//var Serv *echo.Echo
var E echo.Echo

func SetupServer() {

	//e := echo.New()
	E := echo.New()
	//defer e.Close()

	E.Use(middleware.Logger())
	E.Use(middleware.Recover())

	E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//Routes
	E.GET("/api", HomePage)

	E.GET("/api/deck/:game/new/:shuffle", NewDeck) //With cards parameter
	E.GET("/api/deck/:deck_id", OpenDeck)
	E.GET("/api/deck/:deck_id/draw", DrawDeck) //With count parameter

	//Listen web server
	E.Logger.Fatal(E.Start(":8000"))
	//Serv = e

}

//*******************************************

// Main
func main() {
	readme = loadreadme()
	SetupServer()
}
