package main

import (
	"flag"
	"fmt"
)

func main() {

	mode := flag.String("mode", "", "what mode to operate in, server or csv")

	flag.Parse()

	if *mode == "server" {
		ServerMain()
	} else if *mode == "csv" {
		CsvMain()
	} else {
		fmt.Println("mode flag missing or invalid")
	}

}
