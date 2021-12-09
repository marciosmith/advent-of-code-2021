package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// understood that I needed recursion to solve for the flow tracking of the
// basin out the the highest point that isn't 9 but couldn't figure it out and relied heavily
// on https://github.com/alextanhongpin/advent-of-code-2021/blob/main/day09/main.go for part 2

type point struct {
	x, y int
}

func main() {
	inputs := getInputs()
	// inputs := []string{
	// 		"2199943210",
	// 		"3987894921",
	// 		"9856789892",
	// 	"8767896789",
	// 	"9899965678",
	// }

	rowMap := make(map[point]int)
	for y, line := range inputs {
		row := strings.Split(line, "")
		for x, col := range row {
			n := toInt(col)
			rowMap[point{x: x, y: y}] = n
		}
	}
	delta := []point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}

	var basins []int
	for coord, n := range rowMap {
		var adj, lower int
		for _, d := range delta {
			c := coord
			c.x += d.x
			c.y += d.y
			v, ok := rowMap[c]
			if ok {
				adj++
				if n < v {
					lower++
				}
			}
		}
		if adj == lower {
			basins = append(basins, search(rowMap, coord))
		}
	}
	sort.Ints(basins)
	top3 := basins[len(basins)-3:]
	total := 1
	for _, n := range top3 {
		total *= n
	}
	fmt.Println(total)
}

func search(m map[point]int, start point) int {
	seen := make(map[point]bool)
	stack := []point{start}
	delta := []point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}
	var count int
	for len(stack) > 0 {
		head := stack[0]
		stack = stack[1:len(stack)]
		for _, d := range delta {
			c := head
			c.x += d.x
			c.y += d.y

			if seen[c] {
				continue
			}
			seen[c] = true
			v, ok := m[c]
			if !ok {
				continue
			}
			if v < 9 {
				count++
				stack = append(stack, c)
			}
		}
	}
	return count
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
func getInputs() []string {
	var inputs []string
	file, _ := os.Open("../input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputs
}
