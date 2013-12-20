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

		fmt.Println(intToText(n) + "Dollars")
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

// intToText converts an int (< 1 billion) to its textual representation.
func intToText(n int) string {
	output := ""

	// Millioin
	if n >= 1000000 {
		output += intToText(n/1000000) + "Million"
		n = n % 1000000
	}

	// Thousand
	if n >= 1000 {
		output += intToText(n/1000) + "Thousand"
		n = n % 1000
	}

	// Hundred
	if n >= 100 {
		output += intToText(n/100) + "Hundred"
		n = n % 100
	}

	// Twenty to Ninety
	if n >= 20 {
		switch n / 10 {
		case 2:
			output += "Twenty"
			break
		case 3:
			output += "Thirty"
			break
		case 4:
			output += "Forty"
			break
		case 5:
			output += "Fifty"
			break
		case 6:
			output += "Sixty"
			break
		case 7:
			output += "Seventy"
			break
		case 8:
			output += "Eighty"
			break
		case 9:
			output += "Ninety"
			break
		}

		n = n % 10
	}

	// Teens
	if n > 0 {
		switch n {
		case 1:
			output += "One"
			break
		case 2:
			output += "Two"
			break
		case 3:
			output += "Three"
			break
		case 4:
			output += "Four"
			break
		case 5:
			output += "Five"
			break
		case 6:
			output += "Six"
			break
		case 7:
			output += "Seven"
			break
		case 8:
			output += "Eight"
			break
		case 9:
			output += "Nine"
			break
		case 10:
			output += "Ten"
			break
		case 11:
			output += "Eleven"
			break
		case 12:
			output += "Twelve"
			break
		case 13:
			output += "Thirteen"
			break
		case 14:
			output += "Forteen"
			break
		case 15:
			output += "Fifteen"
			break
		case 16:
			output += "Sixteen"
			break
		case 17:
			output += "Seventeen"
			break
		case 18:
			output += "Eighteen"
			break
		case 19:
			output += "Nineteen"
			break
		}
	}

	return output
}
