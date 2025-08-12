package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Unable to connect to database: " + err.Error() + "\n")
	}
	return conn
}

func CloseConnection(conn *pgx.Conn) {
	conn.Close(context.Background())
}

func Init() error {
	conn := OpenConnection()
	defer CloseConnection(conn)

	_, err := conn.Exec(context.Background(), `
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      created_at TIMESTAMPTZ DEFAULT now(),
      updated_at TIMESTAMPTZ DEFAULT now(),
      username TEXT NOT NULL UNIQUE,
      password TEXT NOT NULL,
      last_login TIMESTAMPTZ
    );

    CREATE TABLE IF NOT EXISTS games (
      id SERIAL PRIMARY KEY,
      created_at TIMESTAMPTZ DEFAULT now(),
      updated_at TIMESTAMPTZ DEFAULT now(),
      user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
      start_time TIMESTAMPTZ DEFAULT now(),
      end_time TIMESTAMPTZ,
      width INTEGER NOT NULL,
      height INTEGER NOT NULL,
      number_of_bombs INTEGER NOT NULL,
      number_of_moves INTEGER DEFAULT 0,
      bomb_locations TEXT[][],
      board TEXT[][]
    );

    CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
    BEGIN
      NEW.updated_at = now();
      RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;

    DO $$
    BEGIN
      IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_on_users'
      ) THEN
        CREATE TRIGGER set_updated_at_on_users
        BEFORE UPDATE ON users
        FOR EACH ROW
        EXECUTE FUNCTION update_updated_at_column();
      END IF;
    END$$;

    DO $$
    BEGIN
      IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_updated_at_on_games'
      ) THEN
        CREATE TRIGGER set_updated_at_on_games
        BEFORE UPDATE ON games
        FOR EACH ROW
        EXECUTE FUNCTION update_updated_at_column();
      END IF;
    END$$;
  `)
	return err
}