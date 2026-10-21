package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var help = flag.Bool("h", false, "Print this help message and exit.")

var repeat = flag.Bool("r", false, "Should the clock be continuously updated?")

func printHelp() {
	fmt.Println("Usage: clock [OPTIONS]\n")
	flag.PrintDefaults()
}

func printTime() {
	fmt.Printf("\r%s", time.Now().Format(time.DateTime))
}

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help == true {
		printHelp()
		os.Exit(0)
	}

	if *repeat == true {
		for {
			printTime()
		}
	} else {
		printTime()
	}
}
