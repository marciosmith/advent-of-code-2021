package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := getInputs()
	increased := 0

	for i, number := range inputs {
		if i-1 == -1 {
			// increased++
			continue
		}

		if number > inputs[i-1] {
			increased++
		}
	}
	fmt.Println("increased :: ", increased)
}

func getInputs() []int {
	var inputs []int
	file, _ := os.Open("../input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputs
}
