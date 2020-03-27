package fetcher

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
)

var errNum int

//var rateLimt = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<-rateLimt
	resp, err := http.Get(url)
	if err != nil {
		log.Println("htt get failed:", err)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("errcode :", resp.StatusCode)
		errNum++
		log.Println("errNum:", errNum)
		return nil, errors.New("http statuscod is not ok:")
	}

	respstr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil read all failed:", err)
		return nil, err
	}

	return respstr, nil
}
