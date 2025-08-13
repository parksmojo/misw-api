package db

import (
	"context"
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