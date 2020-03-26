package engine

import (
	"log"
	"spider-go/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		firstReq := requests[0]
		requests = requests[1:]
		log.Println("Fechting:", firstReq.URL)
		body, err := fetcher.Fetch(firstReq.URL)

		if err != nil {
			log.Println("fetcher failed:", err)
			continue
		}
		parseRes := firstReq.ParseFunc(body)
		requests = append(requests, parseRes.Requests...)

		for _, item := range parseRes.Items {
			log.Printf("item:%v", item)
		}
		//log.Println("itemes len=", len(parseRes.Items))
	}
}
