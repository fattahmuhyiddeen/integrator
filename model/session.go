package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	FMIUtil "github.com/fattahmuhyiddeen/integrator/util"

	"github.com/joho/godotenv"

	//blank import to init database here
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

func checkErr(err error, db string) {
	if err != nil {
		FMIUtil.Logg("Failed to connect to " + db)
		panic(err)
	} else {
		FMIUtil.Logg("Successfully connected to " + db)
	}
}

const PostgresDriver = "postgres"
const MssqlDriver = "sqlserver"

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

func TestConnection(db string) {
	query := ""
	if db == "DB1" {
		connectDB1()
		defer disconnectDB1()
		if getDriver(DB1Driver) == MssqlDriver {
			SelectVersion(DB1, "DB1 ("+DB1Server+")")
		} else if getDriver(DB1Driver) == PostgresDriver {
			query = `SELECT * FROM pg_catalog.pg_tables`
			_, err := DB1.Query(query)
			checkErr(err, "DB1 ("+DB1Server+")")
		}
	} else if db == "DB2" {
		connectDB2()
		defer disconnectDB2()
		if getDriver(DB2Driver) == MssqlDriver {
			SelectVersion(DB2, "DB2 ("+DB2Server+")")
		} else if getDriver(DB2Driver) == PostgresDriver {
			query = `SELECT * FROM pg_catalog.pg_tables`
			_, err := DB2.Query(query)
			checkErr(err, "DB2 ("+DB2Server+")")
		}
	}

}

var DB1 *sql.DB
var DB2 *sql.DB

func getDriver(driver string) string {
	if strings.Contains(driver, "post") || strings.Contains(driver, "pg") || strings.Contains(driver, "gres") {
		return PostgresDriver
	}
	if strings.Contains(driver, "ms") || strings.Contains(driver, "sqlserver") || strings.Contains(driver, "microsoft") {
		return MssqlDriver
	}
	return ""
}

func getDBInfo(server, port, user, password, name, driver string) string {
	dbinfo := ""
	if driver == PostgresDriver {
		dbinfo = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
			server, user, password, port, name)
	} else if driver == MssqlDriver {
		dbinfo = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
			server, user, password, port, name)
	}
	return dbinfo
}

func connectDB1() {
	var err error
	driver := getDriver(DB1Driver)
	DB1, err = sql.Open(driver, getDBInfo(DB1Server, DB1Port, DB1User, DB1Password, DB1Name, driver))
	checkErr(err, "DB1 ("+DB1Server+")")
}

func connectDB2() {
	var err error
	driver := getDriver(DB2Driver)
	DB2, err = sql.Open(driver, getDBInfo(DB2Server, DB2Port, DB2User, DB2Password, DB2Name, driver))
	checkErr(err, "DB2 ("+DB2Server+")")
}

func disconnectDB1() {
	DB1.Close()
}

func disconnectDB2() {
	DB2.Close()
}

// SelectVersion is to detect MS SQL Database version
func SelectVersion(db *sql.DB, dbname string) {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		checkErr(err, dbname)
		// log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
