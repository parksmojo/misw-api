package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"misw/auth"
	"misw/db"
	"misw/model"
	"net/http"
)

type newGameRequest struct {
	Width int `json:"width"`
	Height int `json:"height"`
	BombCount int `json:"bombCount"`
}

type newGameResponse struct {
	ID int `json:"id"`
	Board [][]string `json:"board"`
}

func NewGameHandler(w http.ResponseWriter, r *http.Request){
	conn := db.OpenConnection()
	defer db.CloseConnection(conn)

	user := auth.ValidateRequestingUser(w, r, conn)
	if user == nil {
		return
	}

	var body newGameRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.Width < 1 || body.Width > 100 || body.Height < 1 || body.Height > 100 || body.BombCount < 1 || body.BombCount > body.Width * body.Height {
		http.Error(w, "Cannot create specified game board", http.StatusBadRequest)
		return
	}

	board := make([][]string, body.Height)
	for i := range board {
		board[i] = make([]string, body.Width)
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	bombs := make([]model.Coord, 0, body.BombCount)
	used := make(map[string]bool)
	for len(bombs) < body.BombCount {
		x := rand.Intn(body.Width)
		y := rand.Intn(body.Height)
		key := fmt.Sprintf("%d,%d", x, y)
		if !used[key] {
			bombs = append(bombs, model.Coord{X: x, Y: y})
			used[key] = true
		}
	}

	gameId, err := db.CreateGame(conn, user.ID, body.Width, body.Height, body.BombCount, bombs, board)
	if err != nil {
		http.Error(w, "Could not create game", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(newGameResponse{ ID: gameId, Board: board })
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}