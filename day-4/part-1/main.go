package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// inputs := getInputs()
	inputs := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		"8  2 23  4 24",
		"21  9 14 16  7",
		"6 10  3 18  5",
		"1 12 20 15 19",
		"",
		"3 15  0  2 22",
		"9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		"2  0 12  3  7",
	}

	type BingoResults struct {
		BingoNumbers []string
		BingoCard    []string
	}

	bingoCards := getBingoCards(inputs)
	winningBingoCards := map[string]BingoResults{}

	bingoNumbers := strings.Split(inputs[0], ",")
	spew.Dump(bingoNumbers)

	// spew.Dump(getCurrentBingoNumbers(bingoNumbers, 2))

	for bingoNumberIndex := 0; bingoNumberIndex < len(bingoNumbers); bingoNumberIndex++ {

		for _, bingoCard := range bingoCards {
			// bingoCardID = the first line concatted (may cause issues if there are duplicate, may make more unique by appending the sum of the bingo cards)
			bingoCardID := strings.ReplaceAll(bingoCard[0], " ", "")
			bingoNumbers := getCurrentBingoNumbers(bingoNumbers, bingoNumberIndex)
			// only add bingo card on it's first confirmed bingo
			if _, ok := winningBingoCards[bingoCardID]; bingo(bingoNumbers, bingoCard) && !ok {
				winningBingoCards[bingoCardID] = BingoResults{
					BingoNumbers: bingoNumbers,
					BingoCard:    bingoCard,
				}

			}

			// if len(winningBingoCard.BingoCard) == 0 && bingo(bingoNumbers, bingoCard) {
			// 	winningBingoCard.BingoCard = bingoCard
			// 	winningBingoCard.BingoNumbers = bingoNumbers
			// 	break
			// }
		}

	}

	fmt.Println("winningBingoCards :: ")
	spew.Dump(winningBingoCards)

	for _, winningBingoCard := range winningBingoCards {
		sumOfUnmatchedBingoNumbers := sumOfUnmatchedNumbers(winningBingoCard.BingoNumbers, winningBingoCard.BingoCard)
		bingoNumber, _ := strconv.Atoi(winningBingoCard.BingoNumbers[len(winningBingoCard.BingoNumbers)-1])

		fmt.Println("answer :: ", sumOfUnmatchedBingoNumbers*bingoNumber, "len of bingo numbers :: ", len(winningBingoCard.BingoNumbers), "last bingo number :: ", bingoNumber)
	}

}

func sumOfUnmatchedNumbers(bingoNumbers []string, bingoCard []string) int {
	sum := 0
	for _, bingoCardNumbers := range bingoCard {
		for _, cardNumber := range strings.Fields(bingoCardNumbers) {
			if !stringInSlice(cardNumber, bingoNumbers) {
				num, _ := strconv.Atoi(cardNumber)
				sum += num
			}
		}
	}

	return sum
}

func getCurrentBingoNumbers(inputs []string, index int) []string {
	currentBingoNumbers := []string{}
	for i := 0; i < len(inputs); i++ {
		currentBingoNumbers = append(currentBingoNumbers, inputs[i])
		if i == index {
			break
		}
	}

	return currentBingoNumbers
}

// loop over each bingo card row and determine if all the numbers in the row
// have been called
func bingo(bingoNumbers []string, bingoCard []string) bool {

	for _, bingoCardNumbers := range bingoCard {
		bingo := true
		for _, bingoCardNumber := range strings.Fields(bingoCardNumbers) {
			if !stringInSlice(bingoCardNumber, bingoNumbers) {
				bingo = false
			}
		}

		if bingo {
			return true
		}
	}

	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getBingoCards(inputs []string) [][]string {
	bingoCards := [][]string{}
	bingoCardBuilder := []string{}

	for line := 0; line < len(inputs); line++ {
		// Determine if the line is a bingo card line
		if len(strings.Split(inputs[line], " ")) < 4 {
			continue
		}

		bingoCardBuilder = append(bingoCardBuilder, inputs[line])

		// once we have a built out bingo card we can add it to the final array of bingo cards
		if len(bingoCardBuilder) == 5 {
			bingoCards = append(bingoCards, bingoCardBuilder)
			bingoCardBuilder = []string{}
		}

	}

	return bingoCards
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
