package main

import (
	"log"
	"spider-go/engine"
	"spider-go/parser"
	"spider-go/scheduler"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	url := "http://www.zhenai.com/zhenghun"
	//url := "http://album.zhenai.com/u/1177733413"
	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{URL: url, ParseFunc: parser.ParseCityList})
	//concurrentEngine := engine.ConcurrentEngine{Sch: &scheduler.SimpleScheduler{}, WorkerCnt: 100}
	//concurrentEngine.Run(engine.Request{URL: url, ParseFunc: parser.ParseCityList})
	concurrentEngine := engine.ConcurrentEngine{Sch: &scheduler.QueueScheduler{}, WorkerCnt: 100}
	concurrentEngine.Run(engine.Request{URL: url, ParseFunc: parser.ParseCityList})

}
