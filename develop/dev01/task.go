package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func getCurrentTime() {
	exactTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("error receiving the current time: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Current time: %s", exactTime.Format(time.RFC1123))
}

func main() {
	getCurrentTime()
}
