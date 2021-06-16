package cardgames

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	//g "./src/globals"
)

/*
//------------------- Get Param ---------------------------
func GetParam(res http.ResponseWriter, req *http.Request, param string) string {

	result := ""
	r, ok := req.URL.Query()[param]
	if !ok || len(r[0]) < 1 {
		result = ""
	} else {
		result = strings.Join(r, " ")
	}
	if result == "" {
		ShowError(res, param+" parameter not found")
	}
	return result
}
*/

// -----------------------Readfile ----------------------------------

func ReadDataFile(res http.ResponseWriter, filename string) []byte {

	var pathfile string
	var err error
	//filename, err := filepath.Abs(fmt.Sprintf("%s%s", "data/"+game, ".json"))
	test := os.Getenv("TEST")
	if test == "true" {
		pathfile, err = filepath.Abs(fmt.Sprintf("%s%s", "../data/", filename))
	} else {
		pathfile, err = filepath.Abs(fmt.Sprintf("%s%s", "data/", filename))
	}

	//jsonFile, err := os.Open(filename)
	//jsonFile, err := ioutil.ReadFile(filename)
	//byteValue, err := ioutil.ReadFile(filename)
	byteValue, err := ioutil.ReadFile(pathfile)
	// if we os.Open returns an error then handle it
	if err != nil {
		ShowError(res, "File or game not found")
		return nil
		os.Exit(0)
	}

	return byteValue
}

// -------------------------- Generate Token --------------------------------------
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// ---------------------------- LOAD README --------------------------------------------
func Loadreadme() string {
	// SHOW README file
	/*
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

	*/

	Readme = `

CARDGAMES (Poker) App README
----------------------------

Requirements

Golang >=1.6x

Echo Framework

    Installation:
        cd <PROJECT FOLDER IN $GOPATH>
        go get -u github.com/labstack/echo

Unzip zip file on a Go directory:
cd $GOPATH
mkdir cardgames
cd cardgames

Run app in script mode:
go run cardgames.go

Test app:
go test .

Verbose test app:
go test . -v

Build app:
go build

Run (Execute binary):
./cardgames

===========
API RESTFUL
===========

API Help homepage
=================
http://domain:8000/api

Example:
http://127.0.0.1:8000/api

NEW DECK
========

http://domain:8000/api/deck/:game/new/:shuffle

Game values:
    poker
Shuffle values:
    shuffle OR 1
    withoutshuffle (default) OR 0

Parameters:
    ?cards=CODE_CARDS_SEPARATED_COMMAS

Examples:
http://127.0.0.1:8000/api/deck/poker/new/withoutshuffle
http://127.0.0.1:8000/api/deck/poker/new/shuffle?cards=AS,KD,AC,2C,KH

OPEN DECK
=========

http://domain:8000/api/deck/:deck_id

Example:
http://127.0.0.1:8000/api/deck/a31548c8d6baf5f3987233e38ae81d31


DRAW DECK
=========

http://domain:8000/api/deck/:deck_id/draw?count=number_of_cards_draw

Example:
http://127.0.0.1:8000/api/deck/a31548c8d6baf5f3987233e38ae81d31/draw?count=5
`
	return Readme
}
