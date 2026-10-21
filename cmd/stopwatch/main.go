package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var help = flag.Bool("h", false, "Print this help message and exit.")

func printHelp() {
	fmt.Println("Usage: stopwatch [OPTIONS]\nA stopwatch will start. Ctrl+C stops.\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help == true {
		printHelp()
		os.Exit(0)
	}
	start := time.Now()
	for {
		fmt.Printf("\r%v", time.Since(start))
	}
}
