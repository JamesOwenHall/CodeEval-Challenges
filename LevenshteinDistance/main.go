package main

import (
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	flag.Parse()

	testWords, dictionary := parseInputFile(flag.Args()[0])

	for _, word := range testWords {
		fmt.Println(sizeOfSocialNetwork(word, dictionary))
	}
}

// file parsing

func parseInputFile(filename string) ([]string, []string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	lines = removeBlankLines(lines)
	indexEOI := getIndexOfEOI(lines)

	return lines[:indexEOI], lines[indexEOI+1:]
}

func removeBlankLines(in []string) []string {
	l := list.New()

	for _, line := range in {
		if len(line) != 0 {
			l.PushBack(line)
		}
	}

	return listToStringSlice(l)
}

func listToStringSlice(l *list.List) []string {
	result := make([]string, l.Len())
	el := l.Front()

	for i := range result {
		result[i] = el.Value.(string)
		el = el.Next()
	}

	return result
}

func getIndexOfEOI(lines []string) int {
	for i := 0; i < len(lines); i++ {
		if lines[i] == "END OF INPUT" {
			return i
		}
	}

	return -1
}

// levenshtein distance calculation

func hasLevenshteinEq1(word1, word2 string) bool {
	return has1Addition(word1, word2) || has1Subtraction(word1, word2) || has1Substitution(word1, word2)
}

func has1Addition(word1, word2 string) bool {
	if len(word2) != len(word1)+1 {
		return false
	}

	for i := range word2 {
		prefix := word2[:i]

		var suffix string
		if i+1 < len(word2) {
			suffix = word2[i+1:]
		}

		if word1 == prefix+suffix {
			return true
		}
	}

	return false
}

func has1Subtraction(word1, word2 string) bool {
	return has1Addition(word2, word1)
}

func has1Substitution(word1, word2 string) bool {
	if len(word2) != len(word1) {
		return false
	}

	for i := range word1 {
		prefix1 := word1[:i]
		prefix2 := word2[:i]

		var suffix1, suffix2 string
		if i+1 < len(word1) {
			suffix1 = word1[i+1:]
			suffix2 = word2[i+1:]
		}

		if prefix1 == prefix2 &&
			suffix1 == suffix2 &&
			word1[i] != word2[i] {
			return true
		}
	}

	return false
}

// social network

func sizeOfSocialNetwork(word string, dict []string) int {
	l := list.New()
	el := l.PushFront(word)

	for el != nil {
		for _, w := range dict {
			if listContains(l, w) {
				continue
			}

			if hasLevenshteinEq1(el.Value.(string), w) {
				l.PushBack(w)
			}
		}

		el = el.Next()
	}

	return l.Len()
}

func listContains(l *list.List, word string) bool {
	el := l.Front()

	for el != nil {
		if el.Value.(string) == word {
			return true
		}

		el = el.Next()
	}

	return false
}
