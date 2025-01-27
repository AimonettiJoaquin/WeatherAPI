package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to the database")
	return db, nil
}

func CreateUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			optout BOOLEAN DEFAULT FALSE,
			notification_time VARCHAR(5) DEFAULT '08:00'
		)
	`)
	if err != nil {
		return err
	}
	log.Println("Table Users created successfully")
	return nil
}
