package helper

import (
	"log"
	"time"
)

//Start execution
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

//Duration execution
func Duration(msg string, start time.Time) {
	log.Printf("%v: %vmcs\n", msg, time.Since(start).Microseconds())
}
