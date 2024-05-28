package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andreasgylche/gowatch/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Start() {
	addr := os.Getenv("ADDR")
	log.Printf("server listening on %s", addr)
	http.HandleFunc("/post", s.PostDataHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) PostDataHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = database.InsertData(s.db, data.Content)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data inserted successfully"))
}
