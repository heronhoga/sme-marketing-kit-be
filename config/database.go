package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	_ "github.com/go-sql-driver/mysql"
)

func DBInit() (*sql.DB, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbPort := os.Getenv("DATABASE_PORT")

	connString := fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?parseTime=true",
    dbUser,
    dbPassword,
    dbHost,
    dbPort,
    dbName,
	)
	
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("Failed To Connect To Database")
		return nil, err
	}

	log.Println("Successfully Connected To Database")
	return db, nil
}