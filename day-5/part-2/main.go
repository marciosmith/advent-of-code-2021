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
	// horizontal
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

		return coordinates
	}
	// vertical
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

		return coordinates
	}

	// 45deg vertical
	/*						x		y						x		y				x	y					x		y
	starting 			9		7						7		9				1 1					8		0
								8		8						8		8				2	2
	ending				7		9						9		7				3	3					0		8
	*/
	// case	[9, 7 => 7, 9]
	if startingPointX == endingPointY && startingPointY == endingPointX {
		// positive
		if startingPointX < endingPointX {
			for y, x := startingPointY, startingPointX; y != endingPointY-1; {
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", x, y))
				x++
				y--
			}
		}

		// negative
		if startingPointX > endingPointX {
			for x, y := startingPointX, startingPointY; x != endingPointX-1; {
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", x, y))
				x--
				y++
			}
		}
		return coordinates
	}
	// case	[1, 1 => 3, 3]
	if startingPointX == startingPointY && endingPointX == endingPointY {
		if startingPointX-endingPointX < 0 {
			for a := startingPointX; a != endingPointX+1; a++ {
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", a, a))
			}

		}
		if startingPointX-endingPointX > 0 {
			for a := startingPointX; a != endingPointX-1; a-- {
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", a, a))
			}

		}
		return coordinates
	}

	// case	[5, 5 => 8, 2]
	if math.Abs(float64(startingPointX-endingPointX)) == math.Abs(float64(startingPointY-endingPointY)) {
		for x, y := startingPointX, startingPointY; true; {
			coordinates = append(coordinates, fmt.Sprintf("%d,%d", x, y))
			if x == endingPointX {
				break
			}

			if startingPointX > endingPointX {
				x--
			} else {
				x++
			}

			if startingPointY > endingPointY {
				y--
			} else {
				y++
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
