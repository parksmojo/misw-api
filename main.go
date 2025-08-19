package main

import (
	"fmt"
	"log"
	"misw/api"
	"misw/db"
	"misw/middleware"
	"net/http"

	"github.com/joho/godotenv"
)

const VERSION = "0.0.1"
const PORT = "8321"

func handle(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.ApplyTo(handler))
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  err = db.Init()
  if err != nil {
    log.Fatal("Error initializing database: " + err.Error())
  }

  handle("GET /", api.IndexHandlerFactory(VERSION))

  handle("PUT /auth/user", api.CreateUserHandler)
  handle("GET /auth/user", api.GetUserHandler)

  handle("GET /game", api.GetGameHandler)
  handle("PUT /game", api.NewGameHandler)
  handle("POST /game", api.MakeMoveHandler)

  fmt.Printf("Listening on port %s\n", PORT)
  log.Fatal(http.ListenAndServe(":" + PORT, nil))
}
