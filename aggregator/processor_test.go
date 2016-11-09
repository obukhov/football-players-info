package aggregator

import (
	"github.com/obukhov/football-players-info/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"
)

type PlayersCollectionMock struct {
	mock.Mock
}

func (p *PlayersCollectionMock) Add(name string, age int, team string) {
	p.Called(name, age, team)
}

func (p *PlayersCollectionMock) Output(output io.Writer) {
	p.Called(output)
}

func TestAggregator(t *testing.T) {
	playersCollectionMock := new(PlayersCollectionMock)

	aggregator := &Aggregator{
		playersCollection: playersCollectionMock,
	}

	team := api.Team{
		Id:   1,
		Name: "Hogwards",
		Players: []api.Player{
			{
				Name: "Harry Potter",
				Age:  api.NewStringedInt(20),
			},
			{
				Name: "Hermione Granger",
				Age:  api.NewStringedInt(18),
			},
		},
	}

	playersCollectionMock.On("Add", "Harry Potter", 20, "Hogwards")
	playersCollectionMock.On("Add", "Hermione Granger", 18, "Hogwards")

	aggregator.Process(team)

	playersCollectionMock.AssertCalled(t, "Add", "Harry Potter", 20, "Hogwards")
	playersCollectionMock.AssertCalled(t, "Add", "Hermione Granger", 18, "Hogwards")
}

func TestAggregator_Collection(t *testing.T) {
	playersCollectionMock := new(PlayersCollectionMock)

	aggregator := &Aggregator{playersCollection: playersCollectionMock}

	assert.Equal(t, playersCollectionMock, aggregator.Collection())
}
