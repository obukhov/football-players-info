package aggregator

import "github.com/obukhov/football-players-info/api"

func NewAggregator() *Aggregator {
	return &Aggregator{
		playersCollection: NewPlayersCollection(),
	}
}

type Aggregator struct {
	playersCollection *PlayersCollection
}

func (a *Aggregator) Process(team api.Team) {
	for _, player := range team.Players {
		a.playersCollection.Add(player.Name, player.Age.Int(), team.Name)
	}
}

func (a *Aggregator) Collection() *PlayersCollection {
	return a.playersCollection
}
