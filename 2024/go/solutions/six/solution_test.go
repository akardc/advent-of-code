package six_test

import (
	_ "embed"
	"testing"

	"github.com/akardc/advent-of-code/2024/go/solutions/six"
)

//go:embed test_input.txt
var testInput string

func TestPartOne(t *testing.T) {
	if six.PartOne(testInput) != 41 {
		t.Errorf("Part one failed")
	}
}

func TestPartTwo(t *testing.T) {
	if res := six.PartTwo(testInput); res != 6 {
		t.Errorf("Expected 6 but got %d", res)
	}

	in := `##..
...#
....
^.#.`
	if res := six.PartTwo(in); res != 0 {
		t.Errorf("Expected 0 but got %d", res)
	}
}
