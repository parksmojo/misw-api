package main

import (
	"fmt"
	"log"
	"misw/api"
	"misw/db"
	"net/http"

	"github.com/joho/godotenv"
)

const VERSION = "0.0.1"

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  err = db.Init()
  if err != nil {
    log.Fatal("Error initializing database: " + err.Error())
  }

  Handle("GET /", api.IndexHandlerFactory(VERSION))

  Handle("PUT /auth/user", api.CreateUserHandler)
  Handle("GET /auth/user", api.GetUserHandler)

  Handle("PUT /game", api.NewGameHandler)
  Handle("POST /game", api.MakeMoveHandler)

  fmt.Println("Listening on port 8321")
  log.Fatal(http.ListenAndServe(":8321", nil))
}
