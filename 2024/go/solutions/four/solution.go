package four

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var dayInput string

func PartOne(input string) int {
	if input == "" {
		input = dayInput
	}

	parsedInput := parse(input)
	return getXmasCount(parsedInput)
}

func PartTwo(input string) int {
	if input == "" {
		input = dayInput
	}

	parsedInput := parse(input)
	return getCrossMassCount(parsedInput)
}

func parse(input string) [][]string {
	rows := strings.Split(input, "\n")

	var out [][]string
	for i := range len(rows) {
		r := strings.TrimSpace(rows[i])
		if r != "" {
			out = append(out, strings.Split(r, ""))
		}
	}

	return out
}

func getXmasCount(input [][]string) int {
	if len(input) < 1 {
		return 0
	}

	type dir struct {
		x int
		y int
	}
	dirs := []dir{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	total := 0
	for y := range len(input) {
		for x := range len(input[0]) {
			for _, d := range dirs {
				if isXmas(input, x, y, d.x, d.y) {
					total++
				}
			}
		}
	}

	return total
}

func isXmas(input [][]string, x, y, dx, dy int) bool {
	if len(input) < 1 {
		return false
	}

	if len(input[0]) < 1 {
		return false
	}

	if x < 0 || x >= len(input[0]) || (x+3*dx) < 0 || (x+3*dx) >= len(input[0]) {
		return false
	}

	if y < 0 || y >= len(input) || (y+3*dy) < 0 || (y+3*dy) >= len(input) {
		return false
	}

	val := ""
	for i := range 4 {
		val += input[y+i*dy][x+i*dx]
	}

	return val == "XMAS"
}

func getCrossMassCount(input [][]string) int {
	if len(input) < 1 {
		return 0
	}

	total := 0
	for y := range len(input) {
		for x := range len(input[0]) {
			if isCrossMass(input, x, y) {
				total++
			}
		}
	}

	return total
}

func isCrossMass(input [][]string, x, y int) bool {
	if len(input) < 1 {
		return false
	}

	if len(input[0]) < 1 {
		return false
	}

	if x < 1 || x >= len(input[0])-1 {
		return false
	}

	if y < 1 || y >= len(input)-1 {
		return false
	}

	a := ""
	b := ""

	for i := -1; i < 2; i++ {
		a += input[y+i][x+i]
		b += input[y+i][x-i]
	}

	return (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM")
}
