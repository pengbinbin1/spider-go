package parser

import (
	"log"
	"spider-go/engine"
	"spider-go/regex"
)

const (
	match = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)"[^>]*>([^<]+)</a>`
)

func ParseCityList(content []byte) engine.ParseRes {
	res := reg.DoByteSubRegex(content, match)
	var parRes engine.ParseRes
	//limit := 10
	for _, singleRecord := range res {
		//log.Println("url:", singleRecord[1], ",city:", singleRecord[2])
		parRes.Items = append(parRes.Items, singleRecord[2])
		log.Println("City", singleRecord[2])
		parRes.Requests = append(parRes.Requests, engine.Request{URL: string(singleRecord[1]), ParseFunc: ParseCity})
		/*limit--
		if limit <= 0 {
			break
		}*/
	}
	return parRes
}
