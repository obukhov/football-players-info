package multithread

import (
	"github.com/obukhov/football-players-info/api"
	"log"
	"sync"
)

func NewMultiThreadApiReader(apiClient api.ApiClientInterface, threads int, maxIdLimit int) *multiThreadApiReader {
	return &multiThreadApiReader{
		apiClient:  apiClient,
		results:    make(chan api.Team),
		stopChan:   make(chan bool),
		doneChan:   make(chan bool),
		threads:    threads,
		maxIdLimit: maxIdLimit,
	}
}

type multiThreadApiReader struct {
	apiClient  api.ApiClientInterface
	results    chan api.Team
	stopChan   chan bool
	doneChan   chan bool
	threads    int
	maxIdLimit int
}

func (m *multiThreadApiReader) Start() {
	wg := &sync.WaitGroup{}
	wg.Add(m.threads)

	go m.wait(wg)

	for i := 1; i <= m.threads; i++ {
		go m.thread(NewIdGenerator(i, m.threads, m.maxIdLimit), i, wg)
	}
}

func (m *multiThreadApiReader) Stop() {
	close(m.stopChan)
}

func (m *multiThreadApiReader) wait(wg *sync.WaitGroup) {
	wg.Wait()
	close(m.doneChan)
}

func (m *multiThreadApiReader) thread(gen *idGenerator, threadNum int, wg *sync.WaitGroup) {
	complete := false
	for false == complete {
		select {
		case <-m.stopChan:
			complete = true
			return
		default:
			id := gen.Current()
			team, err := m.apiClient.GetTeam(id)
			if nil != err {
				log.Printf("Thread %d error %s fetching id %d", threadNum, err.Error(), id)
				continue
			}

			m.results <- *team
			if valid := gen.GenerateNext(); false == valid {
				complete = true
			}
		}
	}
	wg.Done()
}

func (m *multiThreadApiReader) Read() api.Team {
	return <-m.results
}
