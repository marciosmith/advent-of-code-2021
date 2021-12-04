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
	// 	"forward 5",
	// 	"down 5",
	// 	"forward 8",
	// 	"up 3",
	// 	"down 8",
	// 	"forward 2",
	// }

	directions := map[string]int{
		"forward": 0,
		"down":    0,
		"up":      0,
	}

	for _, input := range inputs {
		movement := strings.Fields(input)
		direction, value := movement[0], movement[1]
		valueNum, _ := strconv.Atoi(value)

		directions[direction] = directions[direction] + valueNum

	}

	fmt.Println("directions :: ", directions)
	finalDepth := 0
	if directions["down"] > directions["up"] {
		finalDepth = directions["down"] - directions["up"]
	} else {
		finalDepth = directions["up"] - directions["down"]
	}

	finalHorizontal := directions["forward"]
	fmt.Println("answer :: ", finalDepth*finalHorizontal)
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
