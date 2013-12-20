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

		output := ""
		for _, word := range strings.Split(line, ";") {
			switch word {
			case "zero":
				output += "0"
				break
			case "one":
				output += "1"
				break
			case "two":
				output += "2"
				break
			case "three":
				output += "3"
				break
			case "four":
				output += "4"
				break
			case "five":
				output += "5"
				break
			case "six":
				output += "6"
				break
			case "seven":
				output += "7"
				break
			case "eight":
				output += "8"
				break
			case "nine":
				output += "9"
				break
			}
		}

		fmt.Println(output)
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
