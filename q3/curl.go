package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	flag.Parse()
	url := flag.Args()[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%q\n", body)
	}
}
