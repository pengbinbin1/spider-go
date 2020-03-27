package engine

import (
	"log"
	//"spider-go/fetcher"
)

type ConcurrentEngine struct {
	Sch       Scheduler
	WorkerCnt int
}

type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (this *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseRes)

	for i := 0; i < this.WorkerCnt; i++ {
		createWorker(out, this.Sch)
	}

	this.Sch.Run()

	for _, request := range seeds {
		this.Sch.Submit(request)
	}

	for {
		res := <-out
		for _, item := range res.Items {
			log.Println("Got item:", item)
		}

		for _, req := range res.Requests {
			this.Sch.Submit(req)
		}

	}
}

func createWorker(out chan ParseRes, sch Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			sch.WorkerReady(in)
			req := <-in

			parseRes, err := worker(req)
			if err != nil {
				continue
			}
			out <- parseRes
		}
	}()
}
