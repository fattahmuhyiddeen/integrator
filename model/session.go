package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var MAIN_DB_USER string
var MAIN_DB_PASSWORD string
var MAIN_DB_NAME string

func init() {
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Fatal("Error cannot log config.txt")
	}
	MAIN_DB_USER = os.Getenv("MAIN_DB_USER")
	MAIN_DB_PASSWORD = os.Getenv("MAIN_DB_PASSWORD")
	MAIN_DB_NAME = os.Getenv("MAIN_DB_NAME")

}

var PGDB *sql.DB

func connectPG() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		MAIN_DB_USER, MAIN_DB_PASSWORD, MAIN_DB_NAME)
	var err error
	PGDB, err = sql.Open("postgres", dbinfo)
	checkErr(err)
}

func disconnectPG() {
	PGDB.Close()
}
