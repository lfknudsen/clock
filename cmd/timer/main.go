package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

var rxSplit = regexp.MustCompile(`(\d+)\s*(\w+)`)

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help == true {
		printHelp()
		os.Exit(0)
	}
	if len(flag.Args()) < 1 {
		printHelp()
		os.Exit(0)
	}
	var hours int64 = 0
	var minutes int64 = 0
	var seconds int64 = 0

	joined := strings.Join(flag.Args(), " ")
	matches := rxSplit.FindAllStringSubmatch(joined, -1)
	for _, match := range matches {
		numberStr := match[1]
		wordStr := match[2]

		number, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			printHelp()
			os.Exit(1)
		}
		switch wordStr[0] {
		case 'h', 'H':
			hours += number
		case 'm', 'M':
			minutes += number
		case 's', 'S':
			seconds += number
		}
	}

	duration := time.Duration(hours*int64(time.Hour) + minutes*int64(time.Minute) + seconds*int64(time.Second))
	goal := time.Now().Add(duration)
	ch := time.After(duration)
	for {
		select {
		case <-ch:
			goto exit
		default:
			fmt.Printf("\r%s", time.Until(goal))
		}
	}
exit:
	fmt.Print("\rTimer completed.                                                                                  ")
}
