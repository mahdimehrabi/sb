package main

import (
	"flag"
	"m1-article-service/application/command"
	"m1-article-service/application/http"
)

func main() {
	// Define command line flags
	filePath := flag.String("file", "", "Path to the JSON file containing user data")
	flag.Parse()

	// read from file or run gin
	if *filePath != "" {
		command.Boot(*filePath)
	} else {
		http.Boot()
	}
}
