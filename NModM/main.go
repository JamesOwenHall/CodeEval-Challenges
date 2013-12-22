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
	for _, line := range strings.Split(fileContents, "\n") {
		if len(line) == 0 {
			continue
		}

		var n, m int
		fmt.Sscanf(line, "%d,%d", &n, &m)
		fmt.Printf("%d\n", modulus(n, m))
	}
}

// MustOpenTextFile reads a file or panics
func MustOpenTextFile(filename string) string {
	inputData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(inputData)
}

func modulus(n, m int) int {
	div := n / m
	return n - (m * div)
}
