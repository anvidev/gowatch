package watcher

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/andreasgylche/gowatch/internal/mover"
	"github.com/andreasgylche/gowatch/internal/processor"
	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
)

const (
	ErrorFolderName     = "error"
	ProcessedFolderName = "processed"
	IncorrectFolderName = "incorrect"
)

type Watcher struct {
	watcher *fsnotify.Watcher
	dir     string
	ext     string
}

func New(dir, ext string) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		watcher: watcher,
		dir:     dir,
		ext:     ext,
	}, nil
}

func (w *Watcher) Start() error {
	log.Println("watching", w.dir, "for", w.ext, "files")

	defer w.watcher.Close()

	done := make(chan bool)

	go w.handleEvents()

	err := w.watcher.Add(w.dir)
	if err != nil {
		return err
	}

	<-done
	return nil
}

func (w *Watcher) handleEvents() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			w.handleEvent(event)
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Println("event", "error:", err)
		}
	}
}

func (w *Watcher) handleEvent(event fsnotify.Event) {
	if event.Op&fsnotify.Create == fsnotify.Create {
		w.handleCreateEvent(event)
	}
}

func (w *Watcher) handleCreateEvent(event fsnotify.Event) {
	if filepath.Ext(event.Name) == w.ext {
		log.Println("watcher", "file created:", event.Name)

		if err := processor.ProcessCSV(event.Name); err != nil {
			log.Println("watcher", "error processing csv:", err)

			if err := mover.MoveFileToFolder(event.Name, ErrorFolderName); err != nil {
				log.Println("watcher", "error moving file to error folder:", err)
				log.Println("watcher", "error details:", fmt.Sprintf("%+v", err))
			}
			if err := beeep.Notify("Error", fmt.Sprintf("Error processing file: %s", event.Name), ""); err != nil {
				log.Println("wathcer", "error sending notification:", err)
			}

		} else {

			if err := mover.MoveFileToFolder(event.Name, ProcessedFolderName); err != nil {
				log.Println("watcher", "error moving file to processed folder:", err)
			}
		}

	} else {

		if err := mover.MoveFileToFolder(event.Name, IncorrectFolderName); err != nil {
			log.Println("watcher", "error moving file to incorrect folder:", err)
		}
	}
}
