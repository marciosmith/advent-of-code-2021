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

	for i := range inputs {
		firstNumber := sumOfInput(inputs, i)
		secondNumber := sumOfInput(inputs, i+1)

		// if we are unable to determine the second number
		// we can exit this loop and print the results
		if secondNumber == 0 {
			break
		}

		if firstNumber < secondNumber {
			increased++
		}
	}

	fmt.Println("increased :: ", increased)
}

func sumOfInput(inputs []int, i int) int {
	lengthOfInputs := len(inputs)
	if i >= lengthOfInputs || i+1 >= lengthOfInputs || i+2 >= lengthOfInputs {
		return 0
	}

	return inputs[i] + inputs[i+1] + inputs[i+2]
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
