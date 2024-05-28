package database

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func InitDB(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS csv_data (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "data" TEXT
    );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func InsertData(db *sql.DB, data string) error {
	insertSQL := `INSERT INTO csv_data(data) VALUES (?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(data)
	if err != nil {
		return err
	}

	return nil
}
