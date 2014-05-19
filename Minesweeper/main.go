package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	MINE    = '*'
	MINESTR = "*"
)

func main() {
	flag.Parse()

	minefields := parseFile(flag.Args()[0])

	for _, minefield := range minefields {
		for y, row := range minefield {
			for x := range row {
				adjacent, isMine := adjacentMines(minefield, x, y)

				if isMine {
					fmt.Print(MINESTR)
				} else {
					fmt.Print(adjacent)
				}
			}
		}

		fmt.Println()
	}
}

func parseFile(filename string) [][][]rune {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	lines = removeBlankLines(lines)

	result := make([][][]rune, 0, len(lines))

	for _, line := range lines {
		result = append(result, getMinefield(line))
	}

	return result
}

func removeBlankLines(lines []string) []string {
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) != 0 {
			result = append(result, line)
		}
	}

	return result
}

func getMinefield(text string) [][]rune {
	sections := strings.Split(text, ";")

	var height, width int
	fmt.Sscanf(sections[0], "%d,%d", &height, &width)

	data := []rune(sections[1])

	rows := make([][]rune, height)

	for y := range rows {
		rows[y] = make([]rune, width)

		for x := range rows[y] {
			index := y*width + x
			rows[y][x] = data[index]
		}
	}

	return rows
}

func adjacentMines(minefield [][]rune, posx, posy int) (int, bool) {
	if minefield[posy][posx] == MINE {
		return 0, true
	}

	count := 0

	for y := posy - 1; y <= posy+1; y++ {
		// bounds check
		if y < 0 || y >= len(minefield) {
			continue
		}

		for x := posx - 1; x <= posx+1; x++ {
			// bounds check
			if x < 0 || x >= len(minefield[y]) {
				continue
			}

			if minefield[y][x] == MINE {
				count++
			}
		}
	}

	return count, false
}
