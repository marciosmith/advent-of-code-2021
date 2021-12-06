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
	// inputs := getInputs()
	inputs := []int{
		3, 4, 3, 1, 2,
	}
	fmt.Println("len(inputs) :: ", len(inputs))
	fmt.Println("inputs :: ", inputs)
	totalLaternFishes := calculateLaternFishByDays(inputs, 80)

	fmt.Println("totalLaternFishes :: ", totalLaternFishes)
}

func calculateLaternFishByDays(startingLanternFishes []int, days int) int {
	totalLanternFish := []int{}
	totalLanternFish = append(totalLanternFish, startingLanternFishes...)

	for day := 0; day < days; day++ {
		startingNumberOfFishForCurrentDay := len(totalLanternFish) - 1
		newLanternFish := []int{}
		for fishIndex := 0; fishIndex < len(totalLanternFish); fishIndex++ {
			// whenever we add a new laternfish we want to skip decrementing internal timer
			if fishIndex > startingNumberOfFishForCurrentDay {
				continue
			}

			// if current fish is on 0, create a new fish and reset internal timer to 6
			if totalLanternFish[fishIndex] == 0 {
				newLanternFish = append(newLanternFish, 8)
				totalLanternFish[fishIndex] = 6
				continue
			}
			totalLanternFish[fishIndex] = totalLanternFish[fishIndex] - 1
		}

		totalLanternFish = append(totalLanternFish, newLanternFish...)
	}

	return len(totalLanternFish)
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
