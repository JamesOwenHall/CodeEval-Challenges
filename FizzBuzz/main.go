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
	lineParser := regexp.MustCompile(`(\d+)\s(\d+)\s(\d+)`)

	for _, v := range strings.Split(fileContents, "\n") {
		if len(v) == 0 {
			continue
		}

		components := lineParser.FindStringSubmatch(v)

		a, _ := strconv.Atoi(components[1])
		b, _ := strconv.Atoi(components[2])
		n, _ := strconv.Atoi(components[3])

		output := make([]string, n)

		for i := 1; i <= n; i++ {
			if i%a == 0 {
				output[i-1] += "F"
			}
			if i%b == 0 {
				output[i-1] += "B"
			}
			if len(output[i-1]) == 0 {
				output[i-1] = strconv.Itoa(i)
			}
		}

		fmt.Println(strings.Join(output, " "))
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
