package multithread

import (
	"github.com/obukhov/football-players-info/api"
	"log"
)

type Searcher struct {
	teamSearchFound map[string]bool
	apiReader       *multiThreadApiReader
	searchCount     int
	foundCont       int
	doneChan        chan bool
}

func NewSearcher(teamNames []string, apiReader *multiThreadApiReader) *Searcher {
	searcher := &Searcher{
		teamSearchFound:make(map[string]bool),
		apiReader: apiReader,
		doneChan: make(chan bool),
	}

	for _, teamName := range teamNames {
		if _, alreadyInSearch := searcher.teamSearchFound[teamName]; false == alreadyInSearch {
			searcher.teamSearchFound[teamName] = false
			searcher.searchCount++
		}
	}

	return searcher
}

func (s *Searcher) Start() {
	go s.receive(s.apiReader.results)
	s.apiReader.Start()
}

func (s *Searcher) receive(input chan *api.Team) {
	done := false
	for false == done {
		team := <-input
		found, lookingFor := s.teamSearchFound[team.Name]
		if false == lookingFor {
			continue
		}

		if false == found {
			s.foundCont++
			s.teamSearchFound[team.Name] = true
			if s.foundCont == s.searchCount {
				done = true
			}
		}
	}

	log.Println("Receiving results done")
	s.apiReader.Stop()
	close(s.doneChan)
}

// Blocks until all elements are found or limits reached
func (s *Searcher) Wait() {
	<- s.doneChan
}

func (s *Searcher) Found() bool {
	return s.foundCont == s.searchCount
}


