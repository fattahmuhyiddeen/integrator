package main

//https://flaviocopes.com/go-date-time-format/

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fattahmuhyiddeen/integrator/model"
	FMIUtil "github.com/fattahmuhyiddeen/integrator/util"

	// "github.com/fattahmuhyiddeen/integrator/model"
	"github.com/joho/godotenv"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	p := fmt.Println

	FMIUtil.Logg("Integrator started")
	model.TestConnection("DB1")
	model.TestConnection("DB2")

	// model.GetAllUsers()
	if true {
		return
	}
	for {
		err := godotenv.Load("config.txt")
		if err != nil {
			log.Fatal("Error cannot log config.txt")
		}
		timeToFetchFromDB1 := os.Getenv("TIME_TO_FETCH_FROM_DB1")
		timeToFetchFromDB2 := os.Getenv("TIME_TO_FETCH_FROM_DB2")

		t := time.Now()
		timeString := ""
		timeString = t.Format("3:04PM")

		p(timeToFetchFromDB2)

		p(timeString)
		p(timeToFetchFromDB1 + " time from env")
		if timeString == timeToFetchFromDB1 {
			p("triggered")
			time.Sleep(1 * time.Minute)
		}
	}
}
