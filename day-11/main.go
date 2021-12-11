package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputs()
	// inputs := []string{
	// 	"5483143223",
	// 	"2745854711",
	// 	"5264556173",
	// 	"6141336146",
	// 	"6357385478",
	// 	"4167524645",
	// 	"2176841721",
	// 	"6882881134",
	// 	"4846848554",
	// 	"5283751526",
	// }

	grid := make(map[xy]int)
	for y, line := range inputs {
		nums := numbers(line)
		for x, level := range nums {
			grid[xy{x, y}] = level
		}
	}

	fmt.Println("grid :: ", grid)

	var flashes int
	for step := 0; step < 100000; step++ {
		flashed := make(map[xy]bool)
		var flash func(pos xy)
		flash = func(pos xy) {
			if flashed[pos] {
				return
			}
			flashed[pos] = true
			flashes++
			// adj
			for _, adjX := range []int{1, 0, -1} {
				for _, adjY := range []int{1, 0, -1} {
					if adjX == 0 && adjY == 0 {
						continue
					}
					np := xy{pos.x + adjX, pos.y + adjY}
					if _, ok := grid[np]; !ok {
						continue
					}
					if flashed[np] {
						continue
					}
					grid[np] = grid[np] + 1
					if grid[np] > 9 {
						flash(np)
					}
				}
			}
		}
		for pos, val := range grid {
			grid[pos] = val + 1
			if grid[pos] > 9 {
				flash(pos)
			}
		}
		if len(flashed) == len(grid) {
			fmt.Println("part2", step+1)
			fmt.Println("len(flashed) :: ", len(flashed))
			return
		}
		for p := range flashed {
			grid[p] = 0
		}
		if step == 99 {
			fmt.Println("part1", flashes)
		}
	}

}

func numbers(in string) []int {
	var list []int
	for _, word := range strings.Split(in, "") {
		number, _ := strconv.Atoi(word)
		list = append(list, number)
	}
	return list
}

type xy struct {
	x int
	y int
}

func getInputs() []string {
	var inputs []string
	file, _ := os.Open("./input.txt")

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
