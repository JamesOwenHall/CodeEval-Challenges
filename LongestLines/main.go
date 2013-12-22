package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		return
	}

	fileContents := MustOpenTextFile(flag.Args()[0])
	lines := strings.Split(fileContents, "\n")
	
	n, _ := strconv.Atoi(lines[0])
	sort.Sort(ByLength(lines[1:]))

	for _, v := range lines[1:n+1] {
		fmt.Println(v)
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

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	// We want to sort in descending order, so we use > instead of <
	return len(s[i]) > len(s[j])
}
