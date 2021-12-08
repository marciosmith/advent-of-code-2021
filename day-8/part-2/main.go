package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := getInputs()
	// inputs := []string{
	// 	"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}
	var part1 int
	for _, line := range data {
		p := strings.Split(line, " | ")
		outputs := strings.Split(p[1], " ")
		for _, v := range outputs {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				part1++
			}
		}
	}

	// they gave the mapping of values
	digits := [10]string{
		"cagedb",  // 0
		"ab",      // 1
		"gcdfa",   // 2
		"fbcad",   // 3
		"eafb",    // 4
		"cdfbe",   // 5
		"cdfgeb",  // 6
		"dab",     // 7
		"acedgfb", // 8
		"cefabd",  // 9
	}
	var part2 int
	for _, line := range data {
		p := strings.Split(line, " | ")
		inputs := strings.Split(p[0], " ")
		outputs := strings.Split(p[1], " ")
		wires := [10]string{}
		for _, v := range inputs {
			l := len(v)
			if l == 2 {
				wires[1] = v
			} else if l == 3 {
				wires[7] = v
			} else if l == 4 {
				wires[4] = v
			} else if l == 7 {
				wires[8] = v
			}
		}

		// find 2
		for i := 0; i < 10; i++ {
			for _, v := range inputs {
				if wires[2] == "" && len(v) == 5 && overlaps(v, wires[7]) == 2 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[3]) == 4 && overlaps(v, wires[4]) == 2 {
					wires[2] = v
				}
				if wires[3] == "" && len(v) == 5 && overlaps(v, wires[1]) == 2 {
					wires[3] = v
				}
				if wires[9] == "" && len(v) == 6 && overlaps(v, wires[7]) == 3 && overlaps(v, wires[0]) == 5 {
					wires[9] = v
				}
				if wires[5] == "" && len(v) == 5 && overlaps(v, wires[2]) == 3 && overlaps(v, wires[4]) == 3 {
					wires[5] = v
				}
				if wires[6] == "" && len(v) == 6 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[2]) == 4 {
					wires[6] = v
				}
				if wires[0] == "" && len(v) == 6 && overlaps(v, wires[1]) == 2 && overlaps(v, wires[2]) == 4 && overlaps(v, wires[3]) == 4 {
					wires[0] = v
				}
			}
		}
		var o string
		for k, v := range wires {
			wires[k] = sortString(v)
		}
		for k, v := range digits {
			digits[k] = sortString(v)
		}
		for _, v := range outputs {
			v = sortString(v)
			for k, w := range wires {
				if w == v {
					o += fmt.Sprintf("%d", k)
				}
			}
		}
		part2 += atoi(o)
	}
	fmt.Println(part1, part2)
}
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func overlaps(a, b string) (count int) {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			count++
		}
	}
	return
}
func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
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
