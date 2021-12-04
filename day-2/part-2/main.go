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
		"aim":     0,
		"depth":   0,
	}

	for _, input := range inputs {
		movement := strings.Fields(input)
		direction, value := movement[0], movement[1]
		valueNum, _ := strconv.Atoi(value)

		if direction == "up" {
			directions["aim"] = directions["aim"] - valueNum
		}

		if direction == "down" {
			directions["aim"] = directions["aim"] + valueNum
		}

		if direction == "forward" {
			directions["forward"] = directions["forward"] + valueNum

			if directions["aim"] > 0 {
				directions["depth"] += directions["aim"] * valueNum
			}

		}

	}

	fmt.Println("directions :: ", directions)
	fmt.Println("answer :: ", directions["depth"]*directions["forward"])
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
