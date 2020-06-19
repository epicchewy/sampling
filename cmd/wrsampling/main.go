package main

import (
	"fmt"
	"os"
)

var version string

func main() {
	fmt.Println("No binary available for weight reservoir sampling.")
	fmt.Println("Please refer to source code and README about usage.")

	// Uncomment the following section and modify to fit your needs
	/*
		samples := flag.Int("-s", 10, "number of items to sample")
		flag.Parse()

		br := bufio.NewReader(os.Stdin)
		itr := sampling.NewBufferedReaderIterator(br)

		reservoir := sampling.NewWeightedReservoir(*samples, itr)

		for _, s := range reservoir.Sample() {
			fmt.Println(s)
		}
	*/

	os.Exit(0)
}
