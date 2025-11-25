package main

import (
	"flag"
	"log"
	"os"

	"github.com/akardc/advent-of-code/2024/go/solutions"
)

func main() {
	flags := flag.NewFlagSet("flags", flag.ExitOnError)
	dayFlag := flags.Int("day", 0, "Integer for the day to execute (1 - 25)")
	partFlag := flags.Int("part", 0, "Integer for the part to execute (1 or 2)")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Failed to parse args: %v", err)
	}

	if dayFlag == nil || partFlag == nil {
		log.Fatalf("Day and part flags must be provided")
	}

	day := *dayFlag
	part := *partFlag

	if day == 0 || part == 0 {
		log.Fatal("Must specify day and part")
	}

	solution, ok := solutions.All[day]
	if !ok {
		log.Fatalf("No solution found for day %d", day)
	}

	var solutionFunc func(string) int
	switch part {
	case 1:
		solutionFunc = solution.One
	case 2:
		solutionFunc = solution.Two
	default:
		log.Fatalf("Part must be 1 or 2")
	}

	if solutionFunc == nil {
		log.Fatalf("Part %d is not implemented for day %d", part, day)
	}

	answer := solutionFunc("")
	log.Printf("Day %d part %d answer: %d", day, part, answer)
}
