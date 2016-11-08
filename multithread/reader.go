package multithread

import (
	"github.com/obukhov/football-players-info/api"
	"log"
)

func NewMultiThreadApiReader(apiClient api.ApiClientInterface, threads int) *multiThreadApiReader {
	return &multiThreadApiReader{
		apiClient:apiClient,
		results: make(chan *api.Team),
		done: make(chan bool),
		threads:threads,
	}
}

type multiThreadApiReader struct {
	apiClient api.ApiClientInterface
	results   chan *api.Team
	done      chan bool
	threads   int
}

func (m *multiThreadApiReader) Start() {
	for i := 1; i <= m.threads; i++ {
		log.Printf("Thread %d starting", i)
		go m.thread(NewIdGenerator(i, m.threads), i)
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
			id := gen.Current()
			team, err := m.apiClient.GetTeam(id)
			if nil != err {
				log.Printf("Thread %d error %s", threadNum, err.Error())
			} else {
				log.Printf("Thread %d fetches [%d] %s", threadNum, team.Id, team.Name)
				m.results <- team
				gen.GenerateNext()
			}
		}
	}
}

func (m *multiThreadApiReader) Read() *api.Team {
	return <-m.results
}
