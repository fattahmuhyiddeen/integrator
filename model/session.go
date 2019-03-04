package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	//blank import to init database here
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var DB1User string
var DB1Password string
var DB1Name string
var DB1Driver string
var DB1Server string
var DB1Port string

var DB2User string
var DB2Password string
var DB2Name string
var DB2Driver string
var DB2Server string
var DB2Port string

func init() {
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Fatal("Error cannot log config.txt")
	}
	DB1User = os.Getenv("DB1_USER")
	DB1Password = os.Getenv("DB1_PASSWORD")
	DB1Name = os.Getenv("DB1_NAME")
	DB1Driver = os.Getenv("DB1_DRIVER")
	DB1Server = os.Getenv("DB1_SERVER")
	DB1Port = os.Getenv("DB1_PORT")

	DB2User = os.Getenv("DB2_USER")
	DB2Password = os.Getenv("DB2_PASSWORD")
	DB2Name = os.Getenv("DB2_NAME")
	DB2Driver = os.Getenv("DB2_DRIVER")
	DB2Server = os.Getenv("DB2_SERVER")
	DB2Port = os.Getenv("DB2_PORT")
}

var DB1 *sql.DB
var DB2 *sql.DB

func getDriver(driver string) string {
	if strings.Contains(driver, "post") || strings.Contains(driver, "pg") || strings.Contains(driver, "gres") {
		return "postgres"
	}
	if strings.Contains(driver, "ms") || strings.Contains(driver, "sqlserver") || strings.Contains(driver, "microsoft") {
		return "sqlserver"
	}
	return ""
}

func getDBInfo(server, port, user, password, name, driver string) string {
	dbinfo := ""
	if driver == "postgres" {
		dbinfo = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
			server, user, password, port, name)
	} else if driver == "sqlserver" {
		dbinfo = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
			server, user, password, port, name)
	}

	return dbinfo
}

func connectDB1() {
	var err error
	driver := getDriver(DB1Driver)
	DB1, err = sql.Open(driver, getDBInfo(DB1Server, DB1Port, DB1User, DB1Password, DB1Name, driver))
	checkErr(err)
}

func connectDB2() {
	var err error
	driver := getDriver(DB2Driver)
	DB2, err = sql.Open(driver, getDBInfo(DB2Server, DB2Port, DB2User, DB2Password, DB2Name, driver))
	checkErr(err)
}

// func connectDB2() {
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
// 		DB2User, DB2Password, DB2Name)
// 	var err error
// 	DB2, err = sql.Open(getDriver(DB2Driver), dbinfo)
// 	checkErr(err)
// }

func disconnectDB1() {
	DB1.Close()
}

func disconnectDB2() {
	DB2.Close()
}

func SelectVersion(db *sql.DB) {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
