// Example: read a CSV into a Ticker and print it.
package main

import (
	"log"

	"github.com/rangertaha/gotal/io"
)

func main() {
	ts, err := io.Read("prices.csv")
	if err != nil {
		log.Fatal(err)
	}
	ts.Print()
}
