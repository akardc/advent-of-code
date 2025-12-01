package seven

import (
	_ "embed"
	"fmt"
	"iter"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var dayInput string

func PartOne(input string) int {
	if input == "" {
		input = dayInput
	}

	// largest := uint64(0)
	// for eq := range parse(input) {
	// 	if eq.answer > largest {
	// 		largest = eq.answer
	// 	}
	// }
	// fmt.Println("largest", largest)
	// return 0

	out := uint64(0)
	for eq := range parse(input) {
		// fmt.Printf("%d: %v\n", eq.answer, eq.values)
		// if !producesAnswer(eq) {
		// 	return 0
		// }
		if producesAnswer(eq, sum, mul) {
			out += eq.answer
		}
	}
	fmt.Println(out)

	return int(out)
}

func producesAnswer(eq equation, operators ...func(a, b uint64) uint64) bool {
	if len(operators) < 1 {
		return false
	}

	opCount := len(operators)
	for i := range pow(opCount, (len(eq.values) - 1)) {
		total := eq.values[0]
		// goodEq := fmt.Sprintf("%d == %d", eq.answer, eq.values[0])

		temp := i
		for j := range len(eq.values) - 1 {
			op := operators[temp%opCount]
			total = op(total, eq.values[j+1])
			temp /= opCount

		}
		// fmt.Println(goodEq)
		if total == eq.answer {
			// fmt.Println(goodEq)
			return true
		}
	}

	return false
}

func sum(a, b uint64) uint64 {
	return a + b
}

func mul(a, b uint64) uint64 {
	return a * b
}

func concat(a, b uint64) uint64 {
	val := strconv.FormatUint(a, 10) + strconv.FormatUint(b, 10)
	out, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		log.Fatalf("failed to parse concat (a=%d, b=%d): %v", a, b, err)
	}
	return out
}

func pow(x int, exp int) int {
	out := 1
	for range exp {
		out *= x
	}
	return out
}

func PartTwo(input string) int {
	if input == "" {
		input = dayInput
	}

	out := uint64(0)
	for eq := range parse(input) {
		if producesAnswer(eq, sum, mul, concat) {
			out += eq.answer
		}
	}
	fmt.Println(out)

	return int(out)
}

type equation struct {
	answer uint64
	values []uint64
}

func parse(input string) iter.Seq[equation] {
	return func(yield func(equation) bool) {
		for row := range strings.SplitSeq(input, "\n") {
			if strings.TrimSpace(row) == "" {
				continue
			}

			parts := strings.Split(row, ":")
			if len(parts) != 2 {
				log.Fatalf("Unable to parse equation - expected 2 parts but got %d: %s", len(parts), row)
			}

			answer, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 64)
			// answer, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("failed to parse equation - error converting '%s' to int from '%s': %v", parts[0], row, err)
			}

			valueParts := strings.Split(strings.TrimSpace(parts[1]), " ")
			values := make([]uint64, 0, len(valueParts))
			for i := range valueParts {
				if val := strings.TrimSpace(valueParts[i]); val != "" {
					v, err := strconv.ParseUint(strings.TrimSpace(valueParts[i]), 10, 64)
					// v, err := strconv.Atoi(strings.TrimSpace(valueParts[i]))
					if err != nil {
						log.Fatalf("failed to parse equation - error converting '%s' to int from '%s': %v", parts[1], row, err)
					}
					values = append(values, v)
				}
			}

			if !yield(equation{answer, values}) {
				return
			}
		}
	}
}
