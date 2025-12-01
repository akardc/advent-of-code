package seven_test

import (
	_ "embed"
	"testing"

	"github.com/akardc/advent-of-code/2024/go/solutions/seven"
)

//go:embed test_input.txt
var testInput string

func TestPartOne(t *testing.T) {
	if res := seven.PartOne(testInput); res != 3749 {
		t.Errorf("Expected 3749 but got %d", res)
	}
}

func TestPartTwo(t *testing.T) {
	if out := seven.PartTwo(testInput); out != 11387 {
		t.Errorf("Expected 11387 but got %d", out)
	}
}
