package main

import (
	"github.com/obukhov/football-players-info/api"
	"github.com/obukhov/football-players-info/multithread"
	"time"
	"log"
)

func main() {
	// init and run api getter
	// init and run data processor
	// format data

	reader := multithread.NewMultiThreadApiReader(api.NewApiClient())
	reader.Start(3)

	go func() {
		for {
			team := reader.Read()
			log.Println(team)
		}
	}()

	time.Sleep(time.Second)
}
