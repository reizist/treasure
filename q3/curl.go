package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Options struct {
	method    string
	headers   string
	data      string
	isVerbose bool
	isHelp    bool
	url       string
}

func ParseOptions() *Options {
	options := &Options{}
	method := flag.String("X", "GET", "Specify request command to use")
	headers := flag.String("H", "", "Pass custom header LINE to server (H)")
	data := flag.String("d", "", "HTTP POST data (H)")
	isVerbose := flag.Bool("v", false, "Make the operation more talkative")
	isHelp := flag.Bool("h", false, "This help text")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 1 && !(*isHelp) {
		fmt.Println("try 'curl -h' for more information")
		os.Exit(2)
	}

	if len(flag.Args()) > 0 {
		url := flag.Args()[0]
		options.url = url
	}

	options.method = *method
	options.headers = *headers
	options.data = *data
	options.isVerbose = *isVerbose
	options.isHelp = *isHelp

	return options
}

func createRequest(options Options) *http.Request {

	var jsonBytes []byte
	if options.method == "POST" || options.method == "PATCH" || options.method == "PUT" {
		jsonBytes = ([]byte)(options.data)
	}

	req, _ := http.NewRequest(options.method, options.url, bytes.NewBuffer(jsonBytes))

	if options.headers != "" {
		headers := strings.Split(options.headers, ";")
		for _, v := range headers {
			headerKey := strings.Split(v, ":")[0]
			headerValue := strings.Split(v, ":")[1]
			req.Header.Add(headerKey, headerValue)
		}
	}

	return req
}

func main() {
	options := ParseOptions()

	if options.isHelp {
		usage()
		os.Exit(1)
	}

	request := createRequest(*options)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
	} else {
		defer resp.Body.Close()

		if options.isVerbose {
			for k, v := range request.Header {
				fmt.Printf("> ")
				fmt.Println(k, ":", strings.Join(v, " "))
			}
			fmt.Println("> ")

			for k, v := range resp.Header {
				fmt.Printf("< ")
				fmt.Println(k, ":", strings.Join(v, " "))
			}
			fmt.Println("< ")
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%q\n", body)
	}
}

var usageinfo string = `curl is a Go implemented CLI cURL-like tool for humans.
Usage:
	curl [flags] URL
flags:
  -v: Make the operation more talkative
METHOD:
  defaults to either GET (if there is no specified method)
URL:
  The only information needed to perform a request is a URL.
Example:
	curl https://example.com
more help information please refer to https://github.com/reizist/treasure/q3
`

func usage() {
	fmt.Println(usageinfo)
	os.Exit(2)
}
