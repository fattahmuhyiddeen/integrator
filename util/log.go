package util

import (
	"io/ioutil"
	"os"
	"time"
)

//Logg is to write log into log file
func Logg(text string) error {
	logFile := "log.txt"
	t := time.Now()
	datetime := t.Format("2006-01-02 15:04:05")
	text = datetime + " - " + text

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		d1 := []byte(text)
		ioutil.WriteFile(logFile, d1, 0755)
		// return err
	}
	defer f.Close()

	_, err = f.WriteString("\n" + text)
	if err != nil {
		return err
	}
	return nil
}
