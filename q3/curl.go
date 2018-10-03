package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createRequest() *http.Request {
	method := flag.String("X", "GET", "Specify request command to use")
	flag.Parse()
	url := flag.Args()[0]
	req, _ := http.NewRequest(*method, url, nil)
	return req
}

func main() {
	request := createRequest()
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%q\n", body)
	}
}
