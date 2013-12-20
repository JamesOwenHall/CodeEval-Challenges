package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}

	fileContents := MustOpenTextFile(flag.Args()[0])
	digitParser := regexp.MustCompile(`[1-9]`)

	for _, line := range strings.Split(fileContents, "\n") {
		if len(line) == 0 {
			continue
		}

		counter := make([]int, 9)
		nums := digitParser.FindAllString(line, -1)

		for _, n := range nums {
			i, _ := strconv.Atoi(n)
			counter[i-1]++
		}

		min := 0
		for i, x := range counter {
			if x == 1 {
				min = i
				break
			}
		}

		fmt.Printf("%d\n", min)
	}
}

// MustOpenTextFile reads a text file or panics
func MustOpenTextFile(filename string) string {
	inputData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(inputData)
}
