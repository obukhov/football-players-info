package main

import (
	"github.com/obukhov/football-players-info/api"
	"github.com/obukhov/football-players-info/multithread"
	"log"
)

func main() {
	// init and run api getter
	// init and run data processor
	// format data
	teamNames := []string{
		"England",
	}

	threads := 3

	reader := multithread.NewMultiThreadApiReader(api.NewApiClient(), threads)
	searcher := multithread.NewSearcher(teamNames, reader)

	log.Printf("Start traversing api in %d threads", threads)
	searcher.Start()
	log.Println("Waiting for result")
	searcher.Wait()

	if searcher.Found() {
		log.Println("Done! Result: found")
	} else {
		log.Println("Done! Result: not found")
	}
}
