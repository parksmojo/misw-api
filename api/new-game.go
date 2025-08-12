package api

import (
	"fmt"
	"net/http"
)


func NewGameHandler(w http.ResponseWriter, r *http.Request){
	newGame(1,1,1,1, 1)
}

func newGame(dbConn, gameId, userId, x, y int) {
	// get game
	// err if no game

	// check x and y are in bounds
	// check (x,y) is a hidden space
	// check (x,y) is not a bomb

	// do opening algorithm

	// check if game end

	// update game board in db

	// return game board
	fmt.Println("your mom gay lol")
}