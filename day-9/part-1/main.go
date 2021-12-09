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
	// 	"2199943210",
	// 	"3987894921",
	// 	"9856789892",
	// 	"8767896789",
	// 	"9899965678",
	// }

	lowRisk := []int{}
	for y, row := range inputs {
		// if y != 0 {
		// 	continue
		// }
		for x := range strings.Split(row, "") {
			currentNumber := strToInt(string(inputs[y][x]))
			// fmt.Println("currentNumber :: ", currentNumber)

			left := false
			right := false
			top := false
			topLeft := false
			topRight := false
			bottom := false
			bottomLeft := false
			bottomRight := false
			// left, right, top, topLeft, topRight, bottom, bottomLeft, bottomRight := false

			// left
			if x-1 >= 0 && x-1 < len(row) {
				left = currentNumber < strToInt(string(row[x-1]))
			} else {
				left = true
			}

			// right
			if x+1 < len(row) {
				right = currentNumber < strToInt(string(row[x+1]))
			} else {
				right = true
			}

			// top
			if y-1 >= 0 {

				//left
				if x-1 >= 0 && x-1 < len(row) {
					topLeft = currentNumber < strToInt(string(string(inputs[y-1][x-1])))
				} else {
					topLeft = true
				}
				//right
				if x+1 < len(row) {
					topRight = currentNumber < strToInt(string(string(inputs[y-1][x+1])))
				} else {
					topRight = true
				}

				top = currentNumber < strToInt(string(string(inputs[y-1][x])))

			} else {
				top = true
				topRight = true
				topLeft = true
			}

			// bottom
			if y+1 < len(inputs) {

				//left
				if x-1 >= 0 && x-1 < len(row) {
					bottomLeft = currentNumber < strToInt(string(string(inputs[y+1][x-1])))
				} else {
					bottomLeft = true
				}
				//right
				if x+1 < len(row) {
					bottomRight = currentNumber < strToInt(string(string(inputs[y+1][x+1])))
				} else {
					bottomRight = true
				}

				bottom = currentNumber < strToInt(string(string(inputs[y+1][x])))

			} else {
				bottom = true
				bottomLeft = true
				bottomRight = true
			}

			// fmt.Println(left, right, top, topLeft, topRight, bottom, bottomLeft, bottomRight)
			if left && right && top && topLeft && topRight && bottom && bottomLeft && bottomRight {
				lowRisk = append(lowRisk, currentNumber)
			}

		}
	}

	fmt.Println("low risk", sum(lowRisk))

}

func sum(array []int) int {
	result := 0
	for _, v := range array {

		result += v + 1
	}
	return result
}

func strToInt(number string) int {
	numberInt, _ := strconv.Atoi(number)
	return numberInt
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
