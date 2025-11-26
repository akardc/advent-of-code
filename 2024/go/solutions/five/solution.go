package five

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var dayInput string

func PartOne(input string) int {
	if input == "" {
		input = dayInput
	}

	rules, updates := parse(input)

	fmt.Println("Rules:")
	for v, rule := range rules {
		fmt.Printf("\t%d: %v\n", v, rule)
	}
	fmt.Println("Updates:")
	sum := 0
	for _, u := range updates {
		isValid := isUpdateValid(u, rules)
		mid := getMidpoint(u)
		if isValid {
			sum += mid
		}
		fmt.Printf("\t%v - %t/%d\n", u, isValid, mid)
	}

	return sum
}

func parse(input string) (map[int]map[int]struct{}, [][]int) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		log.Fatalf("Unable to parse input - expected 2 parts, got %d", len(parts))
	}

	rules := map[int]map[int]struct{}{}
	updates := [][]int{}

	for ruleStr := range strings.SplitSeq(parts[0], "\n") {
		ruleParts := strings.Split(strings.TrimSpace(ruleStr), "|")
		if len(ruleParts) != 2 {
			log.Fatalf("Unable to parse rule - expected 2 parts but got %d: %s", len(ruleParts), ruleStr)
		}

		a, err := strconv.Atoi(ruleParts[0])
		if err != nil {
			log.Fatalf("failed to parse rule - error converting '%s' to int from '%s': %v", ruleParts[0], ruleStr, err)
		}

		b, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			log.Fatalf("failed to parse rule - error converting '%s' to int from '%s': %v", ruleParts[1], ruleStr, err)
		}

		rule, ok := rules[a]
		if !ok {
			rule = map[int]struct{}{}
			rules[a] = rule
		}
		rule[b] = struct{}{}
		rules[a] = rule
	}

	for update := range strings.SplitSeq(strings.TrimSpace(parts[1]), "\n") {
		valStrs := strings.Split(strings.TrimSpace(update), ",")
		vals := make([]int, 0, len(valStrs))
		for _, valStr := range valStrs {
			v, err := strconv.Atoi(valStr)
			if err != nil {
				log.Fatalf("Failed to parse update - error converting '%s' to int from: %v", valStr, err)
			}
			vals = append(vals, v)
		}
		updates = append(updates, vals)
	}

	return rules, updates
}

func isUpdateValid(update []int, rules map[int]map[int]struct{}) bool {
	for i := len(update) - 1; i > 0; i-- {
		rule, ok := rules[update[i]]
		if !ok {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if _, ok := rule[update[j]]; ok {
				return false
			}
		}
	}
	return true
}

func getMidpoint(update []int) int {
	return update[len(update)/2]
}

func PartTwo(input string) int {
	if input == "" {
		input = dayInput
	}

	rules, updates := parse(input)

	sum := 0
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			continue
		}

		fmt.Printf("before:\t%v\n", update)
		slices.SortFunc(update, func(a, b int) int {
			if rule, ok := rules[b]; ok {
				if _, ok := rule[a]; ok {
					return 1
				}
			}
			return -1
		})
		fmt.Printf("after:\t%v\n", update)
		sum += getMidpoint(update)
	}

	return sum
}
