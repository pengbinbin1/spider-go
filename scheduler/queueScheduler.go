package scheduler

import (
	"spider-go/engine"
)

type QueueScheduler struct {
	reqChan    chan engine.Request
	workerChan chan chan engine.Request
}

func (sch *QueueScheduler) Submit(r engine.Request) {
	sch.reqChan <- r

}
func (sch *QueueScheduler) WorkerReady(w chan engine.Request) {
	sch.workerChan <- w
}
func (sch *QueueScheduler) ConfigWorkerChan(chanRequest chan engine.Request) {

}
func (sch *QueueScheduler) Run() {
	sch.reqChan = make(chan engine.Request)
	sch.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request

		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-sch.reqChan:
				{
					requestQ = append(requestQ, r)
				}
			case w := <-sch.workerChan:
				{
					workerQ = append(workerQ, w)
				}
			case activeWorker <- activeRequest:
				{
					requestQ = requestQ[1:]
					workerQ = workerQ[1:]
				}
			}

		}
	}()
}
