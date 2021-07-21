package cardgames

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

const MinSecretKeySize = 16

//------- Load readme file ---------
var Readme string
