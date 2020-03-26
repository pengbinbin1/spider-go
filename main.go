package main

import (
	"log"
	"spider-go/engine"
	"spider-go/parser"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	url := "http://www.zhenai.com/zhenghun"
	//url := "http://album.zhenai.com/u/1177733413"
	engine.Run(engine.Request{URL: url, ParseFunc: parser.ParseCityList})

}
