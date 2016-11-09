package main

import (
	"github.com/obukhov/football-players-info/aggregator"
	"github.com/obukhov/football-players-info/api"
	"github.com/obukhov/football-players-info/multithread"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// init and run api getter
	// init and run data processor
	// format data
	teamNames := []string{
		//"ololo",
		"England",
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester Utd",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"FC Bayern Munich",
	}

	threads := 10

	reader := multithread.NewMultiThreadApiReader(api.NewApiClient(), threads, 500)
	processor := aggregator.NewAggregator()
	searcher := multithread.NewSearcher(teamNames, reader, processor)

	//you can optionally replace ioutil.Discard by os.Stderr to see logs
	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	logger.Printf("Start traversing api in %d threads", threads)
	searcher.Start()
	logger.Println("Waiting for result")
	searcher.Wait()

	if false == searcher.Found() {
		found, search := searcher.FoundStat()
		logger.Printf("Done! Result: not all found: %d of %d", found, search)
		os.Exit(1)
	}

	logger.Println("Done! Result: found")
	processor.Collection().Output(os.Stdout)
}
