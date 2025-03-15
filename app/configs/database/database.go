package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DBInfo struct {
	host     string
	port     string
	user     string
	password string
	schema   string
	dbName   string
}

func (info *DBInfo) createConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		info.user,
		info.password,
		info.host,
		info.port,
		info.dbName,
	)
}

func (info *DBInfo) DBSessionCreate() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", info.createConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (info *DBInfo) DBSessionPing() *sqlx.DB {
	db, err := sqlx.Open("postgres", info.createConnectionString())

	if err != nil {
		log.Fatal("Error in creating settings for connection", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error in network:", err.Error())
	}

	err = db.Close()
	if err != nil {
		log.Fatal("Error in closing session:", err.Error())
	}

	return db
}

func LoadDatabaseSettings(path string) *DBInfo {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	return &DBInfo{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SCHEMA"),
		os.Getenv("DB_NAME"),
	}
}
