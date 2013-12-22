package main

import (
	"container/list"
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

		stack := list.New()
		success := true

		for i := range line {
			if isOpenBracket(line[i:i+1]) {
				stack.PushBack(line[i:i+1])
				continue
			}
			if stack.Len() == 0 {
				success = false
				break
			}

			opener := stack.Remove(stack.Back()).(string)

			if !isMatchingBracket(opener, line[i:i+1]) {
				success = false
				break
			}
		}

		if success == false || stack.Len() != 0 {
			fmt.Println("False")
		} else {
			fmt.Println("True")
		}
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

func isOpenBracket(c string) bool {
	if c == "(" || c == "[" || c == "{" {
		return true
	}

	return false
}

func isMatchingBracket(opener, closer string) bool {
	if opener == "(" && closer == ")" {
		return true
	}
	if opener == "[" && closer == "]" {
		return true
	}
	if opener == "{" && closer == "}" {
		return true
	}

	return false
}
