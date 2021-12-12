package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputs := getInputs()
	// inputs := []string{
	// 	"start-A",
	// 	"start-b",
	// 	"A-c",
	// 	"A-b",
	// 	"b-d",
	// 	"A-end",
	// 	"b-end",
	// }

	fmt.Println("part1 (input):", part1(inputs))
	fmt.Println("part2 (input):", part2(inputs))

}

func part1(lines []string) int {
	m := make(map[string]map[string]bool)
	for _, line := range lines {
		paths := strings.Split(line, "-")
		from, to := paths[0], paths[1]
		if m[from] == nil {
			m[from] = make(map[string]bool)
		}
		if m[to] == nil {
			m[to] = make(map[string]bool)
		}
		m[from][to] = true
		m[to][from] = true
	}

	var count int
	stack := [][]string{{"start"}}
	for len(stack) > 0 {
		var last []string
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		tail := last[len(last)-1]
		for k := range m[tail] {
			visited := make(map[string]bool)
			for _, door := range last {
				visited[door] = true
			}
			if strings.ToLower(k) == k && visited[k] {
				continue
			}

			t := make([]string, len(last))
			copy(t, last)
			t = append(t, k)

			if k == "end" {
				count++
			}
			stack = append(stack, t)
		}
	}
	return count
}

func part2(lines []string) int {

	m := make(map[string]map[string]bool)
	for _, line := range lines {
		paths := strings.Split(line, "-")
		from, to := paths[0], paths[1]
		if m[from] == nil {
			m[from] = make(map[string]bool)
		}
		if m[to] == nil {
			m[to] = make(map[string]bool)
		}
		m[from][to] = true
		m[to][from] = true
	}

	stack := [][]string{{"start"}}
	var count int
	for len(stack) > 0 {
		var last []string
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		tail := last[len(last)-1]
		for k := range m[tail] {
			visited := make(map[string]int)
			valid := true
			twice := 0
			for _, door := range last {
				visited[door]++
				if strings.ToLower(door) == door && visited[door] >= 2 {
					twice++
					if twice > 1 {
						valid = false
						break
					}
				}
				if door == "start" && visited["start"] > 1 {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			t := make([]string, len(last))
			copy(t, last)
			t = append(t, k)

			if k == "end" {
				count++
				continue
			}
			stack = append(stack, t)
		}
	}
	return count
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
