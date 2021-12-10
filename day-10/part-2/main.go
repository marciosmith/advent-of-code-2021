package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	scores = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	opening = map[string]bool{
		"(": true,
		"<": true,
		"[": true,
		"{": true,
	}
	closing = map[string]bool{
		")": true,
		">": true,
		"]": true,
		"}": true,
	}
	openToClose = map[string]string{
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
		")": "(",
		">": "<",
		"]": "[",
		"}": "{",
	}
)

var corruptedLines = map[string]int{}

func main() {
	inputs := getInputs()
	// inputs := []string{
	// 	"[({(<(())[]>[[{[]{<()<>>",
	// 	"[(()[<>])]({[<{<<[]>>(",
	// 	"{([(<{}[<>[]}>{[]{[(<()>", // corrupted
	// 	"(((({<>}<{<{<>}{[]{[]{}",
	// 	"[[<[([]))<([[{}[[()]]]",  // corrupted
	// 	"[{[{({}]{}}([{[{{{}}([]", // corrupted
	// 	"{<[[]]>}<{[{[{[]{()[[[]",
	// 	"[<(<(<(<{}))><([]([]()", // corrupted
	// 	"<{([([[(<>()){}]>(<<{{", // corrupted
	// 	"<{([{{}}[<[[[<>{}]]]>[]]",
	// }

	var result []int
	for _, row := range inputs {
		var stack []string
		valid := true

		for _, char := range strings.Split(row, "") {
			if opening[char] {
				stack = append(stack, char)
			}

			if closing[char] {
				lastChar := stack[len(stack)-1]
				if lastChar != openToClose[char] {
					valid = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		if !valid {
			continue
		}

		score := 0
		for i := range stack {
			s := stack[len(stack)-1-i]
			score *= 5
			score += scores[openToClose[s]]
		}
		result = append(result, score)

	}
	sort.Ints(result)
	fmt.Println("middle :: ", result[len(result)/2])
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
