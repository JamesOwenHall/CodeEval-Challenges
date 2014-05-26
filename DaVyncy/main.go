package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	flag.Parse()

	fileBytes, _ := ioutil.ReadFile(flag.Args()[0])
	fileStr := string(fileBytes)

	testCases := parseProblem(fileStr)

	for _, testCase := range testCases {
		fmt.Println(solve(testCase))
	}
}

func parseProblem(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = removeEmptyLines(lines)

	result := make([][]string, 0, len(lines))

	for _, line := range lines {
		result = append(result, strings.Split(line, ";"))
	}

	return result
}

func removeEmptyLines(lines []string) []string {
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}

func solve(segments []string) string {
	for len(segments) > 1 {
		maxCombo, maxLength := "", 0
		iMax, jMax := 0, 0

		for i := 0; i < len(segments)-1; i++ {
			for j := i + 1; j < len(segments); j++ {
				combo, length := largestOverlap(segments[i], segments[j])

				if length > maxLength {
					maxCombo, maxLength = combo, length
					iMax, jMax = i, j
				}
			}
		}

		segments = remove(segments, iMax, jMax)
		segments = append(segments, maxCombo)
	}

	return segments[0]
}

func largestOverlap(s1, s2 string) (combination string, length int) {
	for i := -len(s2) + 1; i < len(s1)-1; i++ {
		lowerBound := intMax(i, 0)
		upperBound := intMin(len(s1), len(s2)+i)

		if s1[lowerBound:upperBound] == s2[lowerBound-i:upperBound-i] &&
			upperBound-lowerBound > length {
			length = upperBound - lowerBound
			combination = ""

			// prefix
			if i < 0 {
				combination += s2[:-i]
			} else if i > 0 {
				combination += s1[:i]
			}

			combination += s1[lowerBound:upperBound]

			// suffix
			if i+len(s2) > len(s1) {
				combination += s2[len(s1)-i:]
			} else if i+len(s2) < len(s1) {
				combination += s1[len(s2)+i:]
			}
		}
	}

	return
}

func remove(input []string, index ...int) []string {
	result := make([]string, 0, len(input)-len(index))

	for i := 0; i < len(input); i++ {
		if !contains(index, i) {
			result = append(result, input[i])
		}
	}

	return result
}

func contains(input []int, element int) bool {
	for _, s := range input {
		if s == element {
			return true
		}
	}

	return false
}

func intMax(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}

	return i2
}

func intMin(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}

	return i2
}
