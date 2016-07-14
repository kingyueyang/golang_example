package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	content := HttpGet("http://www.google.com")
	fmt.Println(content)
}

func HttpGet(urls string) (x string) {
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println("Http Get has some error!")
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read io error!")
		return ""
	}

	return string(body)
}

