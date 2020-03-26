package reg

import (
	"fmt"
	//	"log"
	"regexp"
)

const (
	str = `my email is test@gmail.1com 
	my penmg@qq.com `
)

func GetEmail() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	//match := re.FindString(str)
	match := re.FindAllString(str, -1)
	//match := re.FindAll([]byte(str), -1)
	fmt.Println("match:", match)
}

func DoStringRegex(source string, match string) []string {
	re := regexp.MustCompile(match)
	matchStrs := re.FindAllString(source, -1)
	return matchStrs
}

func DoStringSubRegex(source string, match string) [][]string {
	re := regexp.MustCompile(match)
	matchStr := re.FindAllStringSubmatch(source, -1)
	return matchStr

}

func DoByteSubRegex(source []byte, match string) [][]string {
	var res [][]string

	re := regexp.MustCompile(match)
	matchStr := re.FindAllSubmatch(source, -1)
	for _, strings := range matchStr {
		var tempRes []string
		for _, singlestr := range strings {
			tempRes = append(tempRes, string(singlestr))
		}
		res = append(res, tempRes)
	}
	return res
}

func DoByteRegex(source []byte, match string) []string {
	var res []string
	re := regexp.MustCompile(match)
	matchStrs := re.FindAll(source, -1)
	for _, temp := range matchStrs {
		res = append(res, string(temp))
	}
	return res
}

func DoProfile(source []byte, match string) string {
	var res []string

	re := regexp.MustCompile(match)
	matchStr := re.FindSubmatch(source)

	for _, strings := range matchStr {
		res = append(res, string(strings))
	}

	if len(res) >= 2 {
		return res[1]
	} else {
		//log.Println("res is invalid")
		return ""
	}

}

/*
func main() {
	match := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`
	res := DoStringSubRegex(str, match)
	res1 := DoByteSubRegex([]byte(str), match)
	fmt.Println("res:", res)
	fmt.Println("res1:", res1)
}*/
