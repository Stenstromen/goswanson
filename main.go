package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Quote []string

const ENDPOINT = "https://ron-swanson-quotes.herokuapp.com"

func printUsage() {
	fmt.Print("A CLI tool to interact with " + ENDPOINT + "\n\n")
	fmt.Print("Usage:\n")
	fmt.Print(os.Args[0] + " [Flag]\n\n")
	fmt.Print("Flags:\n")
	fmt.Print("    " + "-h, --help" + "             " + "This help text" + "\n")
	fmt.Print("    " + "-s, --search string" + "    " + "Search Quote" + "\n")
	fmt.Print("    " + "-No flag-" + "              " + "Get Random Quote\n")
}

func getQuote() {
	resp, err := http.Get(ENDPOINT + "/v2/quotes")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	var res Quote
	if err := json.Unmarshal(body, &res); err != nil {
		os.Exit(1)
	}

	fmt.Print(res[0] + "\n")
}

func searchQuote(term string) {
	resp, err := http.Get(ENDPOINT + "/v2/quotes/search/" + term)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	var res Quote
	if err := json.Unmarshal(body, &res); err != nil {
		os.Exit(1)
	}

	for i := 0; i < len(res); i++ {
		fmt.Print("* " + res[i] + "\n")
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printUsage()
		}

		if os.Args[1] == "-s" || os.Args[1] == "--search" {
			searchQuote(os.Args[2])
		}
	} else {
		getQuote()
	}
	os.Exit(0)
}
