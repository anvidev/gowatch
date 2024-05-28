package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreasgylche/gowatch/internal/database"
	"github.com/andreasgylche/gowatch/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		tursoUrl string = os.Getenv("DATABASE_URL")
		token    string = os.Getenv("DATABASE_AUTH_TOKEN")
		url      string = fmt.Sprintf("%s?authToken=%s", tursoUrl, token)
	)

	db, err := database.InitDB(url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(db)
	srv.Start()
}
