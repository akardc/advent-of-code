package six

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var dayInput string

func PartOne(input string) int {
	if input == "" {
		input = dayInput
	}

	guardMap, pos := parse(input)
	visitedPositions := map[location]bool{}

	// fmt.Println("Map:")
	// for _, row := range guardMap {
	// 	fmt.Printf("\trow: %v\n", row)
	// }
	// fmt.Printf("Guard position: %d, %d\n", pos.x, pos.y)

	visitedPositions[pos.location] = true

	// fmt.Printf("GuardMap: %d, %v\n", len(guardMap), guardMap[0])
	for range 100000 {
		pos, _ = move(guardMap, pos)
		if pos.exited {
			break
		}
		visitedPositions[pos.location] = true
	}

	// sb := strings.Builder{}
	// sb.Grow(len(input) - sb.Cap())
	// for y := range guardMap {
	// 	for x := range guardMap[y] {
	// 		if guardMap[y][x] {
	// 			sb.WriteString("#")
	// 		} else if visitedPositions[location{x, y}] {
	// 			sb.WriteString("X")
	// 		} else {
	// 			sb.WriteString(".")
	// 		}
	// 	}
	// 	sb.WriteString("\n")
	// }
	//
	// fmt.Printf("Positions:\n%v\n", visitedPositions)
	// fmt.Printf("Travel:\n%s\n", sb.String())
	return len(visitedPositions)
}

type direction struct {
	dx int
	dy int
}

type location struct {
	x int
	y int
}

type position struct {
	location
	direction
	exited bool
}

func move(guardMap [][]bool, pos position, obstacles ...location) (position, bool) {
	newX := pos.x + pos.dx
	newY := pos.y + pos.dy

	turned := false

	if newX < 0 || newX >= len(guardMap[0]) || newY < 0 || newY >= len(guardMap) {
		pos.exited = true
	} else if guardMap[newY][newX] {
		pos.direction = turnRight(pos.direction)
		turned = true
	} else if slices.Contains(obstacles, location{x: newX, y: newY}) {
		pos.direction = turnRight(pos.direction)
		turned = true
	} else {
		pos.x = newX
		pos.y = newY
	}

	return pos, turned
}

func turnRight(dir direction) direction {
	if dir.dy == -1 {
		dir.dx = 1
		dir.dy = 0
	} else if dir.dx == 1 {
		dir.dx = 0
		dir.dy = 1
	} else if dir.dy == 1 {
		dir.dx = -1
		dir.dy = 0
	} else {
		dir.dx = 0
		dir.dy = -1
	}

	return dir
}

func parse(input string) ([][]bool, position) {
	guardMap := [][]bool{}
	pos := position{}
	for row := range strings.SplitSeq(input, "\n") {
		if strings.TrimSpace(row) == "" {
			continue
		}
		mapRow := make([]bool, 0, len(row))
		for i, r := range row {
			switch r {
			case '^':
				pos.x = i
				pos.y = len(guardMap)
				pos.dx = 0
				pos.dy = -1
				mapRow = append(mapRow, false)
			case '#':
				mapRow = append(mapRow, true)
			default:
				mapRow = append(mapRow, false)
			}
		}
		guardMap = append(guardMap, mapRow)
	}

	return guardMap, pos
}

func isHittingWall(pos position, guardMap [][]bool) bool {
	x := pos.x + pos.dx
	y := pos.y + pos.dy
	if y < 0 || y >= len(guardMap) || x < 0 || x >= len(guardMap[y]) {
		return false
	}
	return guardMap[y][x]
}

var count int = 0

func leadsToLoop(pos position, obstacle location, guardMap [][]bool) bool {
	pathChecked := map[location][]direction{}
	// fmt.Printf("Checking for loop: %+v\n", pos)
	for !pos.exited {
		// fmt.Printf("\t%+v\n", pos)
		// visitedPos, ok := visitedPositions[pos.location]

		// if ok && slices.Contains(visitedPos, pos.direction) {
		// 	// We're on a visited space heading in the same direction so this is a loop
		// 	fmt.Println("Found loop via visited positions")
		// 	return true
		// }

		if checkedSpace, ok := pathChecked[pos.location]; ok && slices.Contains(checkedSpace, pos.direction) {
			// we already checked this space in this direction, so we're in a loop
			// fmt.Printf("\tFound loop via checks: %v - %v\n\n", checkedSpace, pos)
			if count < 10 {
				fmt.Println("Found a loop")
				sb := strings.Builder{}
				for y := range guardMap {
					for x := range guardMap[y] {
						if guardMap[y][x] {
							sb.WriteString("#")
						} else if _, ok := pathChecked[location{x, y}]; ok {
							sb.WriteString("X")
						} else {
							sb.WriteString(".")
						}
					}
					sb.WriteString("\n")
				}
				fmt.Println(sb.String())
			}
			return true
		}

		pathChecked[pos.location] = append(pathChecked[pos.location], pos.direction)
		pos, _ = move(guardMap, pos, obstacle)
	}

	// fmt.Printf("\tNo loop found\n\n")
	return false
}

func PartTwo(input string) int {
	if input == "" {
		input = dayInput
	}

	guardMap, pos := parse(input)
	visitedPositions := map[location][]direction{}

	visitedPositions[pos.location] = append(visitedPositions[pos.location], pos.direction)
	loopMakers := map[location]bool{}

	for range 100000 {
		if !isHittingWall(pos, guardMap) {
			turningRight := position{
				location:  pos.location,
				direction: turnRight(pos.direction),
			}
			obstacle := location{x: pos.x + pos.dx, y: pos.y + pos.dy}
			if len(visitedPositions[obstacle]) < 1 && leadsToLoop(turningRight, obstacle, guardMap) {
				loopMakers[obstacle] = true
			}
		}

		pos, _ = move(guardMap, pos)
		if pos.exited {
			break
		}

		visitedPositions[pos.location] = append(visitedPositions[pos.location], pos.direction)
	}

	// fmt.Printf("LoopMakers: %v\n", loopMakers)
	return len(loopMakers)
}
