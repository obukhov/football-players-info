package main

import (
	"github.com/obukhov/football-players-info/aggregator"
	"github.com/obukhov/football-players-info/api"
	"github.com/obukhov/football-players-info/multithread"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	teamNames := getTeamNames()
	threads, maxIdLimit := getEnvConfig()

	reader := multithread.NewMultiThreadApiReader(api.NewApiClient(), threads, maxIdLimit)
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
func getEnvConfig() (int, int) {
	threads, err := strconv.Atoi(os.Getenv("API_THREADS"))
	if nil != err || threads < 1 {
		threads = 10
	}

	maxIdLimit, err := strconv.Atoi(os.Getenv("MAX_ID_LIMIT"))
	if nil != err || maxIdLimit < 1 {
		maxIdLimit = 1000
	}

	return threads, maxIdLimit
}

func getTeamNames() []string {

	if len(os.Args) > 1 {
		return os.Args[1:]
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		byteIn, err := ioutil.ReadAll(os.Stdin)
		teamNames := strings.Split(string(byteIn), "\n")
		if nil == err || len(teamNames) > 0 {
			return teamNames
		}
	}

	return []string{
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
}
