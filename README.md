
# CARDGAMES (Poker) App README
----------------------------

Author: Manuel Salguero

### IMPORTANT
This proyect in an test exercise on Golang with TDD tests.

It is not the complete game. It is not perfect. I am aware of its flaws. It has static errors, files bad organization and bad struct. In this case does not have clean code practices. But what is done, it works !!

Surely, by spending more time on it, you could get a higher quality source code. But time is a precious commodity. And it is better to show what you have to not show what you don't know when you will have it.

#### The worst project is the one that was not tried ;-)

### Requirements

Golang >=1.6x

Echo Framework

    Installation:
        cd <PROJECT FOLDER IN $GOPATH>
        go get -u github.com/labstack/echo

Unzip zip file on a Go directory:
cd $GOPATH
mkdir cardgames
cd cardgames

### Run app in script mode:
go run cardgames.go

### Test app:
go test .

### Verbose test app:
go test . -v

Build app:
go build

Run (Execute binary):
./cardgames

### ===========
### API RESTFUL
### ===========

### API Help homepage
http://domain:8000/api

Example:
http://127.0.0.1:8000/api

### NEW DECK

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

### OPEN DECK

http://domain:8000/api/deck/:deck_id

Example:
http://127.0.0.1:8000/api/deck/a31548c8d6baf5f3987233e38ae81d31


### DRAW DECK

http://domain:8000/api/deck/:deck_id/draw?count=number_of_cards_draw

Example:
http://127.0.0.1:8000/api/deck/a31548c8d6baf5f3987233e38ae81d31/draw?count=5
