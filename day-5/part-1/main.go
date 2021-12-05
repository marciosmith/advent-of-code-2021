package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputs()
	// inputs := []string{
	// 	"0,9 -> 5,9",
	// 	"8,0 -> 0,8",
	// 	"9,4 -> 3,4",
	// 	"2,2 -> 2,1",
	// 	"7,0 -> 7,4",
	// 	"6,4 -> 2,0",
	// 	"0,9 -> 2,9",
	// 	"3,4 -> 1,4",
	// 	"0,0 -> 8,8",
	// 	"5,5 -> 8,2",
	// }

	allCoordinatesPoints := map[string]int{}
	overlappingPoints := 0

	for _, line := range inputs { // line = "0,9 -> 5,9"
		startingPoint := strings.Split(line, " -> ")[0] // 0,9
		endingPoint := strings.Split(line, " -> ")[1]   // 5,9

		startingPointX, _ := strconv.Atoi(strings.Split(startingPoint, ",")[0]) // 0
		endingPointX, _ := strconv.Atoi(strings.Split(endingPoint, ",")[0])     // 5

		startingPointY, _ := strconv.Atoi(strings.Split(startingPoint, ",")[1]) // 9
		endingPointY, _ := strconv.Atoi(strings.Split(endingPoint, ",")[1])     // 9

		currentCoordinatesPoints := determineCurrentCoordinatesPoints(startingPointX, endingPointX, startingPointY, endingPointY)

		for _, cp := range currentCoordinatesPoints {
			allCoordinatesPoints[cp] = allCoordinatesPoints[cp] + 1
		}

	}

	allCoordinatesPointsdata, _ := json.Marshal(allCoordinatesPoints)
	fmt.Println("allCoordinatesPoints :: ", string(allCoordinatesPointsdata))

	for coordinateKey := range allCoordinatesPoints {
		if allCoordinatesPoints[coordinateKey] >= 2 {
			overlappingPoints++
		}
	}

	fmt.Println("overlappingPoints :: ", overlappingPoints)

}

func determineCurrentCoordinatesPoints(startingPointX, endingPointX, startingPointY, endingPointY int) []string {
	/*						x		y
	starting 			0		9
								1		9
								2		9
								3		9
								4		9
	ending				5		9
	based off this example we know:
	> either the x or y axis will stay constant
	> once we know which axis is constant we need to figure out if the other axis starting - ending is positive or negative
	> if it's negative we need to increment the starting until we get to the ending
	> if its postive we need to decrement the ending until we get to the beginning
	> in each if statement we will append the new plotted point to the coordinates array
	*/
	coordinates := []string{}
	if startingPointX == endingPointX {
		// positive
		if startingPointY-endingPointY > 0 {
			for i := startingPointY; i != endingPointY-1; i-- {
				// startingPointX or endingPointX can be used as the second string value
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", startingPointX, i))
			}
		}

		// negative
		if startingPointY-endingPointY < 0 {
			for i := startingPointY; i != endingPointY+1; i++ {
				// startingPointX or endingPointX can be used as the second string value
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", startingPointX, i))
			}
		}
	}

	if startingPointY == endingPointY {
		// positive
		if startingPointX-endingPointX > 0 {
			for i := startingPointX; i != endingPointX-1; i-- {
				// startingPointY or endingPointY can be used as the second string value
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", i, startingPointY))
			}
		}

		// negative
		if startingPointX-endingPointX < 0 {
			for i := startingPointX; i != endingPointX+1; i++ {
				// startingPointY or endingPointY can be used as the second string value
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", i, startingPointY))
			}
		}
	}

	return coordinates
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
