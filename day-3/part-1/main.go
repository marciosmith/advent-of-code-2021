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
	// 	"00100",
	// 	"11110",
	// 	"10110",
	// 	"10111",
	// 	"10101",
	// 	"01111",
	// 	"00111",
	// 	"11100",
	// 	"10000",
	// 	"11001",
	// 	"00010",
	// 	"01010",
	// }

	indexTracker := map[int]map[string]int{}
	gammeRate := ""
	epsilonRate := ""

	for i := range inputs {
		for r, bitValue := range strings.Split(inputs[i], "") {
			if indexTracker[r] == nil {
				indexTracker[r] = map[string]int{}
			}

			indexTracker[r][bitValue]++
		}
	}

	// why does this not work the same as the variable for loop?
	// for s := range indexTracker {
	for s := 0; s <= len(indexTracker)-1; s++ {
		binaryTracker := indexTracker[s]

		if binaryTracker["0"] > binaryTracker["1"] {
			gammeRate += "0"
			epsilonRate += "1"
		} else {
			gammeRate += "1"
			epsilonRate += "0"
		}
	}

	fmt.Println(indexTracker)
	fmt.Println("gammeRate :: ", gammeRate)
	fmt.Println("epsilonRate :: ", epsilonRate)

	gamma, err := convertBinaryToDecimal(gammeRate)
	epsilon, err := convertBinaryToDecimal(epsilonRate)
	if err != nil {
		panic("unable to convert binary to number")
	}

	fmt.Println("answer :: ", gamma*epsilon)
}

func convertBinaryToDecimal(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
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
