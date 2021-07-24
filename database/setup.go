package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func SetupDB() *sql.DB {
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psql)

	if err != nil {
		log.Fatal("Error connect to DB")
	}

	return db
}
