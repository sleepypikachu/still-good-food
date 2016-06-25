package main

import (
	"encoding/json"
	"fmt"
	"github.com/sleepypikachu/food"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>\n", os.Args[0])
		return
	}

	r, err := food.Scrape(os.Args[1])
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}
