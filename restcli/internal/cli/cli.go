package cli

import (
	"flag"
	"fmt"
	"os"

	rq "restcli/internal/httprq"
)

//type httpRq interface {
//	PostRq(url, body *string)
//	GetRq(url *string)
//}

// Commands ...
func Commands() {
	fmt.Println("Hello")
	postCmd := flag.NewFlagSet("POST", flag.ExitOnError)
	getCmd := flag.NewFlagSet("GET", flag.ExitOnError)

	postURL := postCmd.String("url", "", "URL for POST request (Required)")
	postText := postCmd.String("body", "", "Body of POST request")
	//postHeaders := postCmd.String("headers", "", "Headers of POST request")

	getURL := getCmd.String("url", "", "URL for GET request (Required)")
	//getHeaders := getCmd.String("headers", "", "Headers of GET request")

	if len(os.Args) < 2 {
		fmt.Println("POST or GET subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "POST":
		postCmd.Parse(os.Args[2:])
	case "GET":
		getCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if postCmd.Parsed() {
		// Requred flags
		if *postURL == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		rq.NewHTTPRq().PostRq(*postURL, *postText)
	}

	if getCmd.Parsed() {
		// Requred flags
		if *getURL == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		rq.NewHTTPRq().GetRq(*postURL)
	}
}
