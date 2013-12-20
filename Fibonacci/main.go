package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

		n, _ := strconv.Atoi(line)

		if n <= 1 {
			fmt.Printf("%d\n", n)
			continue
		}

		a, b := 0, 1
		for i := 1; i < n; i++ {
			a, b = b, a + b
		}

		fmt.Printf("%d\n", b)
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
