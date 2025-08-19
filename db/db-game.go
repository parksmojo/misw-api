package db

import (
	"context"
	"encoding/json"
	"errors"
	"misw/model"

	"github.com/jackc/pgx/v5"
)

func GetGamesForUser(conn *pgx.Conn, userID int) ([]model.Game, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, created_at, updated_at, user_id, start_time, end_time, width, height, number_of_bombs, number_of_moves, bomb_locations, board FROM games WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []model.Game
	for rows.Next() {
		var game model.Game
		if err := rows.Scan(
			&game.ID,
			&game.CreatedAt,
			&game.UpdatedAt,
			&game.UserID,
			&game.StartTime,
			&game.EndTime,
			&game.Width,
			&game.Height,
			&game.NumberOfBombs,
			&game.NumberOfMoves,
			&game.BombLocations,
			&game.Board,
		); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func CreateGame(conn *pgx.Conn, userID, width, height, bombCount int, bombLocations []model.Coord, board [][]string) (int, error) {
	bombsJSON, err := json.Marshal(bombLocations)
	if err != nil {
		return -1, err
	}

	var gameID int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO games (user_id, width, height, number_of_bombs, bomb_locations, board)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		userID, width, height, bombCount, bombsJSON, board,
	).Scan(&gameID)
	if err != nil {
		return -1, err
	}
	return gameID, nil
}

func GetGame(conn *pgx.Conn, userID, gameID int) ([][]string, error) {
	var board [][]string
	err := conn.QueryRow(
		context.Background(), 
		`SELECT board FROM games WHERE user_id=$1 AND id=$2`, 
		userID, gameID,
	).Scan(&board)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return board, nil
}