package multithread

import (
	"github.com/obukhov/football-players-info/api"
	"log"
)

func NewMultiThreadApiReader(apiClient api.ApiClientInterface) *multiThreadApiReader {
	return &multiThreadApiReader{
		apiClient:apiClient,
		results: make(chan *api.Team),
		done: make(chan bool),
	}
}

type multiThreadApiReader struct {
	apiClient api.ApiClientInterface
	results   chan *api.Team
	done      chan bool
}

func (m *multiThreadApiReader) Start(threads int) {
	for i := 1; i <= threads; i++ {
		log.Printf("Thread %d starting", i)
		go m.thread(NewIdGenerator(i, threads), i)
	}
}

func (m *multiThreadApiReader) Stop() {
	close(m.done)
}

func (m *multiThreadApiReader) thread(gen *idGenerator, threadNum int) {
	log.Printf("Thread %d started", threadNum)
	complete := false
	for false == complete {
		select {
		case <-m.done:
			complete = true
			log.Printf("Thread %d stops", threadNum)
			return
		default:
			id := gen.Next()
			team, err := m.apiClient.GetTeam(id)
			if nil != err {
				log.Printf("Thread %d error %s", threadNum, err.Error())
			}
			log.Printf("Thread %d fetches [%d] %s", threadNum, team.Id, team.Name)
			m.results <- team
		}
	}
}

func (m *multiThreadApiReader) Read() *api.Team {
	return <-m.results
}
