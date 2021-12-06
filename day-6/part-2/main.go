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
	// inputs := []int{
	// 	3, 4, 3, 1, 2,
	// }

	totalLaternFishes := calculateLaternFishByDays(inputs, 256)

	fmt.Println("totalLaternFishes :: ", totalLaternFishes)
}

func calculateLaternFishByDays(startingLanternFishes []int, days int) int {
	totalLanternFishEachDay := make([]int, 9)

	for _, startingFishIndex := range startingLanternFishes {
		totalLanternFishEachDay[startingFishIndex]++
	}

	for day := 0; day < days; day++ {
		// count of how many fishes are born today
		bornToday := totalLanternFishEachDay[0]

		// Shuffle everyone down one day, moving today's to the end
		totalLanternFishEachDay = append(totalLanternFishEachDay[1:], totalLanternFishEachDay[0])

		// It takes 6 days for today's fish to recharge
		totalLanternFishEachDay[6] += bornToday

		// and 8 days for the new fish to start producing more fish
		totalLanternFishEachDay[8] = bornToday

	}

	total := 0
	for _, n := range totalLanternFishEachDay {
		total += n
	}

	return total
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
