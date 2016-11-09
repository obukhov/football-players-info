package aggregator

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Player struct {
	Name      string
	Age       int
	TeamNames []string
}

func NewPlayersCollection() *PlayersCollection {
	return &PlayersCollection{
		players:     make(map[string]*Player),
		playerNames: make([]string, 0),
	}
}

type PlayersCollection struct {
	players     map[string]*Player
	playerNames []string
}

type PlayersCollectionInterface interface {
	Add(name string, age int, teamName string)
	Output(output io.Writer)
}

func (pc *PlayersCollection) Add(name string, age int, teamName string) {
	if _, found := pc.players[name]; false == found {
		pc.players[name] = &Player{
			Name:      name,
			Age:       age,
			TeamNames: make([]string, 0),
		}
		pc.playerNames = append(pc.playerNames, name)
	}
	pc.players[name].TeamNames = append(pc.players[name].TeamNames, teamName)
}

func (pc *PlayersCollection) Output(output io.Writer) {
	sort.Strings(pc.playerNames)

	i := 1
	for _, playerName := range pc.playerNames {
		player := pc.players[playerName]
		sort.Strings(player.TeamNames)

		fmt.Fprintf(
			output,
			"%d. %s; %d; %s\n",
			i,
			player.Name,
			player.Age,
			strings.Join(player.TeamNames, ", "),
		)
		i++
	}
}
