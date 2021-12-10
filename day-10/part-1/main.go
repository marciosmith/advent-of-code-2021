package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	scores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
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
	total := 0

	for _, row := range inputs {
		var stack []string
		for _, char := range strings.Split(row, "") {
			if opening[char] {
				stack = append(stack, char)
			}

			if closing[char] {
				lastChar := stack[len(stack)-1]
				if lastChar != openToClose[char] {
					total += scores[char]
					break
				}
				stack = stack[:len(stack)-1]
				fmt.Println("stack :: ", stack)
			}
		}
	}

	fmt.Println("total :: ", total)
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
