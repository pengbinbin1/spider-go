package engine

import (
	"log"
	"spider-go/fetcher"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		firstReq := requests[0]
		requests = requests[1:]

		//log.Println("itemes len=", len(parseRes.Items))
		parseRes, err := worker(firstReq)
		if err != nil {
			continue
		}
		requests = append(requests, parseRes.Requests...)
		for _, item := range parseRes.Items {
			log.Printf("item:%v", item)
		}
	}
}

func worker(r Request) (ParseRes, error) {
	log.Println("Fechting:", r.URL)
	body, err := fetcher.Fetch(r.URL)

	if err != nil {
		log.Println("fetcher failed:", err)
		return ParseRes{}, err
	}
	parseRes := r.ParseFunc(body)
	return parseRes, nil
}
