package solutions

import (
	"github.com/akardc/advent-of-code/2024/go/solutions/five"
	"github.com/akardc/advent-of-code/2024/go/solutions/four"
	"github.com/akardc/advent-of-code/2024/go/solutions/six"
)

type Parts struct {
	One func(input string) int
	Two func(input string) int
}

var All = map[int]Parts{
	4: {
		One: four.PartOne,
		Two: four.PartTwo,
	},
	5: {
		One: five.PartOne,
		Two: five.PartTwo,
	},
	6: {
		One: six.PartOne,
		Two: six.PartTwo,
	},
}
