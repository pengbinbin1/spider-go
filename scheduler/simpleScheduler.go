package scheduler

import (
	"spider-go/engine"
)

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (this *SimpleScheduler) ConfigWorkerChan(workin chan engine.Request) {
	this.WorkChan = workin
}

func (this *SimpleScheduler) Submit(req engine.Request) {
	go func() { this.WorkChan <- req }()
}

func (this *SimpleScheduler) WorkerReady(ch chan engine.Request) {

}
func (this *SimpleScheduler) Run() {

}
