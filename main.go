package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Quote []string

func main() {
	resp, err := http.Get("https://ron-swanson-quotes.herokuapp.com/v2/quotes")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	var res Quote
	if err := json.Unmarshal(body, &res); err != nil {
		os.Exit(1)
	}

	fmt.Print(res[0] + "\n")
}
