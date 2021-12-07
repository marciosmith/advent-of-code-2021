package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputs()
	// inputs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	max := 0
	for _, crab := range inputs {
		if crab > max {
			max = crab
		}
	}
	fmt.Println("max :: ", max)
	distance := make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, crab := range inputs {
			distance[i] += cos(int(math.Abs(float64(crab - i))))
		}
	}

	min := 999999999999
	for _, value := range distance {
		if value < min {
			min = value
		}
	}

	fmt.Println("fuelConsumption :: ", min)

	// wrong answers: [351077]

}

func cos(number int) int {
	return (number * (number + 1)) / 2
}

func getInputs() []int {
	var inputs []int
	file, _ := os.Open("../input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), ",")
		for _, numberStr := range numbers {
			number, _ := strconv.Atoi(numberStr)
			inputs = append(inputs, number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputs
}
