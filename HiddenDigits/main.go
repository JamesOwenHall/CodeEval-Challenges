package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}

	fileContents := MustOpenTextFile(flag.Args()[0])
	parser := regexp.MustCompile(`[a-j0-9]`)

	for _, line := range strings.Split(fileContents, "\n") {
		if len(line) == 0 {
			continue
		}

		output := ""

		for _, d := range parser.FindAllString(line, -1) {
			ascii := []byte(d)

			if ascii[0] >= 48 && ascii[0] <= 57 {
				output += d
			} else {
				ascii[0] -= 49
				output += string(ascii)
			}
		}

		if len(output) == 0 {
			output = "NONE"
		}
		fmt.Println(output)
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
