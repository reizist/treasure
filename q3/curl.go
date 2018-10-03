package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Options struct {
	method    string
	headers   string
	isVerbose bool
	url       string
}

func ParseOptions() *Options {
	options := &Options{}
	method := flag.String("X", "GET", "Specify request command to use")
	headers := flag.String("H", "", "Pass custom header LINE to server (H)")
	isVerbose := flag.Bool("v", false, "Make the operation more talkative")
	flag.Parse()

	url := flag.Args()[0]

	options.url = url
	options.method = *method
	options.headers = *headers
	options.isVerbose = *isVerbose

	return options
}

func createRequest(options Options) *http.Request {

	req, _ := http.NewRequest(options.method, options.url, nil)

	if options.headers != "" {
		headerKey := strings.Split(options.headers, ":")[0]
		headerValue := strings.Split(options.headers, ":")[1]
		req.Header.Add(headerKey, headerValue)
	}

	return req
}

func main() {
	options := ParseOptions()

	request := createRequest(*options)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
	} else {
		defer resp.Body.Close()

		if options.isVerbose {
			fmt.Printf("%q\n", resp.Header)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%q\n", body)
	}
}
