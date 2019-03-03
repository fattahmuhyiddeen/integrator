package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	// fmt.Println(time.Now().Truncate(time.Minute))
	// t2, e := time.Parse(form, "8 41 PM")
	p := fmt.Println

	// model.GetAllUsers()
	// if true {
	// 	return
	// }
	for {
		err := godotenv.Load("config.txt")
		if err != nil {
			log.Fatal("Error cannot log config.txt")
		}
		timeToFetchFromERP := os.Getenv("TIME_TO_FETCH_FROM_ERP")
		timeToFetchFromMainDatabase := os.Getenv("TIME_TO_FETCH_FROM_MAIN_DATABASE")

		t := time.Now()
		timeString := ""
		timeString = t.Format("3:04PM")

		p(timeToFetchFromMainDatabase)

		p(timeString)
		p(timeToFetchFromERP + " time from env")
		if timeString == timeToFetchFromERP {
			p("triggered")
			time.Sleep(1 * time.Minute)
		}
	}
}
