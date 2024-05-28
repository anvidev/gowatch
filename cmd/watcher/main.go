package main

import (
	"log"
	"os"

	"github.com/andreasgylche/gowatch/internal/watcher"
)

func main() {
	dir := "C:/Users/AndreasVifert/Desktop/gowatch"
	ext := ".csv"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", dir)
	}

	w, err := watcher.New(dir, ext)
	if err != nil {
		log.Fatal(err)
	}

	if err = w.Start(); err != nil {
		log.Fatal(err)
	}
}
