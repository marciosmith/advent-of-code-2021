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

	oxygenGeneratorRating, err := convertBinaryToDecimal(
		getLastInputRecursively(inputs, "oxygen", 0)[0])
	co2ScrubberRating, err := convertBinaryToDecimal(
		getLastInputRecursively(inputs, "co2", 0)[0])

	if err != nil {
		panic("unable to convert binary to decimal")
	}

	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating
	fmt.Println("oxygenGeneratorRating :: ", oxygenGeneratorRating)
	fmt.Println("co2ScrubberRating :: ", co2ScrubberRating)
	fmt.Println("lifeSupportRating :: ", lifeSupportRating)

}

func getLastInputRecursively(inputs []string, ratingType string, position int) []string {

	if len(inputs) == 2 && ratingType == "oxygen" {
		return filterInputByPostionAndValue(inputs, position, "1")
	}

	if len(inputs) == 2 && ratingType == "co2" {
		return filterInputByPostionAndValue(inputs, position, "0")
	}

	if len(inputs) > 1 {
		newInput, newPosition := filterInputs(inputs, ratingType, position)
		return getLastInputRecursively(newInput, ratingType, newPosition)
	}

	return inputs
}

func filterInputs(inputs []string, ratingType string, position int) ([]string, int) {
	inputIndexPositionTracker := map[string]int{}
	finalInput := []string{}

	for _, binary := range inputs {
		inputIndexPositionTracker[strings.Split(binary, "")[position]]++
	}

	for _, binary := range inputs {

		if ratingType == "oxygen" {
			if inputIndexPositionTracker["0"] > inputIndexPositionTracker["1"] && strings.Split(binary, "")[position] == "0" {
				finalInput = append(finalInput, binary)
			}

			if inputIndexPositionTracker["1"] > inputIndexPositionTracker["0"] && strings.Split(binary, "")[position] == "1" {
				finalInput = append(finalInput, binary)
			}
		}

		if ratingType == "co2" {
			if inputIndexPositionTracker["0"] < inputIndexPositionTracker["1"] && strings.Split(binary, "")[position] == "0" {
				finalInput = append(finalInput, binary)
			}

			if inputIndexPositionTracker["1"] < inputIndexPositionTracker["0"] && strings.Split(binary, "")[position] == "1" {
				finalInput = append(finalInput, binary)
			}
		}

	}
	return finalInput, position + 1
}

func filterInputByPostionAndValue(inputs []string, position int, value string) []string {
	lastInput := []string{}
	for _, binary := range inputs {
		if strings.Split(binary, "")[position] == value {
			lastInput = append(lastInput, binary)
		}
	}

	return lastInput
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
