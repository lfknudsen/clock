package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("\r%s", time.Now().Format(time.DateTime))
	}
}
