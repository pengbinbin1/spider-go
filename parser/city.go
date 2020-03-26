package parser

import (
	"log"
	"spider-go/engine"
	"spider-go/regex"
)

const (
	citymatch = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
)

func ParseCity(content []byte) engine.ParseRes {
	res := reg.DoByteSubRegex(content, citymatch)
	var parRes engine.ParseRes

	for _, singleRecord := range res {
		//log.Println("url:", singleRecord[1], ",city:", singleRecord[2])
		parRes.Items = append(parRes.Items, singleRecord[2])
		log.Println("http:", singleRecord[1])
		log.Println("user:", singleRecord[2])
		parRes.Requests = append(parRes.Requests, engine.Request{URL: string(singleRecord[1]), ParseFunc: ParseProfile})

	}
	return parRes
}
