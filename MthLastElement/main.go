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
	problems := parseInput(fileContents)

	for _, p := range problems {
		if p.digit == 0 {
			continue
		}

		index := len(p.chars) - p.digit
		if index < 0 {
			continue
		}

		fmt.Printf("%s\n", p.chars[index])
	}
}

type Problem struct {
	chars []string
	digit int
}

// MustOpenTextFile reads a text file or panics
func MustOpenTextFile(filename string) string {
	inputData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(inputData)
}

func parseInput(input string) []Problem {
	digitParser := regexp.MustCompile(`\d+`)
	charParser := regexp.MustCompile(`[a-zA-Z]`)

	lines := strings.Split(input, "\n")
	output := make([]Problem, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		digit := digitParser.FindString(line)
		output[i].digit, _ = strconv.Atoi(digit)

		output[i].chars = charParser.FindAllString(line, -1)
	}

	return output
}
