package main

import (
	"fmt"
	"log"
	"misw/api"
	"misw/db"
	"net/http"

	"github.com/joho/godotenv"
)


func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  err = db.Init()
  if err != nil {
    log.Fatal("Error initializing database" + err.Error())
  }

  http.Handle("POST /game", http.HandlerFunc(api.MakeMoveHandler))
  http.Handle("PUT /game", http.HandlerFunc(api.NewGameHandler))

  fmt.Println("Listening on port 8321")
  log.Fatal(http.ListenAndServe(":8321", nil))
}
