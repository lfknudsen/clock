package main

import (
	"flag"
	"fmt"
	"time"
)

var help = flag.Bool("h", false, "Print this help message and exit.")

var repeat = flag.Bool("r", false, "Should the clock be continuously updated?")

func printTime() {
	fmt.Printf("\r%s", time.Now().Format(time.DateTime))
}

func printHelp() {
	fmt.Println("Usage: clock [OPTIONS]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help == true {
		printHelp()
	}

	if *repeat == true {
		for {
			printTime()
		}
	} else {
		printTime()
	}
}
