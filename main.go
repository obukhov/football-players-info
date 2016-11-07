package main

import (
	"github.com/obukhov/football-players-info/api"
	"log"
)

func main() {
	// init and run api getter
	// init and run data processor
	// format data

	client := api.NewApiClient()
	team, err := client.GetTeam(1)

	log.Println(team, err)
}
