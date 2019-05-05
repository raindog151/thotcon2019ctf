package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var host = "134.209.163.195"
var port = "80"
var cnt  = 0

func main() {
	recursiveGet(host, port)
}

func recursiveGet(host, port string) {

	cnt += 1

	url := fmt.Sprintf("http://%s:%s/", host, port)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile("[0-9]+")
	find := string(re.Find(body))

	// Secondary directory

	url = fmt.Sprintf("http://%s:%s/%s", host, port, find)

	response, err = http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	re = regexp.MustCompile("[0-9]+")
	newport := string(re.Find(body))

	re = regexp.MustCompile("[a-f0-9]+-[a-f0-9]+-[a-f0-9]+-[a-f0-9]+-[a-f0-9]+")
	key := string(re.Find(body))
	if key != "" {
		fmt.Println(key)
		os.Exit(0)
	}

	fmt.Printf("Count: %d Newport: %s\n", cnt, newport)

	recursiveGet(host, newport)

}
