package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/epicchewy/sampling"
)

var version string

func main() {
	// seed random
	rand.Seed(time.Now().UnixNano())

	var samples int
	flag.IntVar(&samples, "samples", 10, "number of items to sample")
	flag.Parse()

	br := bufio.NewReader(os.Stdin)
	itr := sampling.NewBufferedReaderIterator(br)

	reservoir := sampling.NewReservoir(samples, itr)

	for _, s := range reservoir.Sample() {
		fmt.Println(s)
	}

	os.Exit(0)
}
