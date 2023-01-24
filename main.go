package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "CSV file in format of 'question,answer' to be importeed")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErrNil(err)
	os.Exit(1)

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	checkErrNil(err)
	fmt.Println(lines)

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
