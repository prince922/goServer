package scheduler

import (
	"crawler/engine"
	"golang.org/x/net/context"
	"time"
)

type QueuedScheduler struct {
	requstChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requstChan <- r
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// 把一个任务放到任务管道中分配（可以在这里加上一些过滤逻辑）
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requstChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {

			// 超时处理
			cxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)

			defer cancel()

			var activeReqeust engine.Request
			var activeWoker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeReqeust = requestQ[0]
				activeWoker = workerQ[0]
			}

			select {
			case <-cxt.Done():
				goto Loop
			case r := <-s.requstChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWoker <- activeReqeust:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}

	Loop:
	}()
}
