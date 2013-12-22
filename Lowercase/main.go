package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}

	fileContents := MustOpenTextFile(flag.Args()[0])
	fmt.Print(strings.ToLower(fileContents))
}

// MustOpenTextFile reads a text file or panics
func MustOpenTextFile(filename string) string {
	inputData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(inputData)
}
