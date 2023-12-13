package main

import (
	"chatgpt4/tdameritrade"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// read .env file
	// read env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set the location to New York
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	//
	tdameritrade.StartStreamingData(loc)
}
