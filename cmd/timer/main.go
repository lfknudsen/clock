package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var help = flag.Bool("h", false, "Print this help message and exit.")

func printHelp() {
	fmt.Println("Usage: timer [OPTIONS] <duration>\nStart a timer.\n" +
		"Duration formatting: <num>h <num>m <num>s\n" +
		"Whitespace optional. Can use the full words (e.g. minute or minutes) too.\n" +
		"Case insensitive.")
	flag.PrintDefaults()
}

//var rxSplit = regexp.MustCompile(`(\d+)\s*h(?:ours?)?\s*(\d+)\s*m(?:inutes?)?\s*(\d+)\s*s(?:econds?)?`)

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help == true {
		printHelp()
		os.Exit(0)
	}
	args := flag.Args()
	if len(args) < 2 || len(args)%2 != 0 {
		printHelp()
		os.Exit(0)
	}
	hours := 0
	minutes := 0
	seconds := 0
	for i := 0; i < len(args); i += 2 {
		number, err := strconv.Atoi(args[i])
		if err != nil {
			printHelp()
			os.Exit(1)
		}
		switch args[i+1][0] {
		case 'h', 'H':
			hours += number
		case 'm', 'M':
			minutes += number
		case 's', 'S':
			seconds += number
		}
	}
}
