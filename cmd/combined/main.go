package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/andreasgylche/gowatch/internal/database"
	"github.com/andreasgylche/gowatch/internal/server"
	"github.com/andreasgylche/gowatch/internal/watcher"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		tursoUrl string = os.Getenv("DATABASE_URL")
		token    string = os.Getenv("DATABASE_AUTH_TOKEN")
		url      string = fmt.Sprintf("%s?authToken=%s", tursoUrl, token)

		dir string = "C:/Users/AndreasVifert/Desktop/gowatch"
		ext string = ".csv"
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

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", dir)
	}

	w, err := watcher.New(dir, ext)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(db)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err = w.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		srv.Start()
	}()

	wg.Wait()
	log.Println("All services stopped")
}
