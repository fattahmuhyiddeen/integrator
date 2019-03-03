package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	//blank import to init database here
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var mainDBUser string
var mainDBPassword string
var mainDBName string

func init() {
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Fatal("Error cannot log config.txt")
	}
	mainDBUser = os.Getenv("MAIN_DB_USER")
	mainDBPassword = os.Getenv("MAIN_DB_PASSWORD")
	mainDBName = os.Getenv("MAIN_DB_NAME")
}

//PgDB is instance of postgreSQL
var PgDB *sql.DB

func connectPG() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		mainDBUser, mainDBPassword, mainDBName)
	var err error
	PgDB, err = sql.Open("postgres", dbinfo)
	checkErr(err)
}

func disconnectPG() {
	PgDB.Close()
}
