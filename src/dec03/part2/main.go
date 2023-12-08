package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const filePath = "../../../data/dec3/input.txt"

type numIndex struct {
	number int
	start  int
	end    int
}

// Check if symbol
func isSymbol(s string, index int) bool {
	r := []rune(s)
	if index < 0 || index == len(r) {
		return false
	}
	return !unicode.IsDigit(r[index]) && r[index] != '.'
}

// Generator function: get the numbers and index
func getNumbers(s string) <-chan numIndex {
	channel := make(chan numIndex)
	go func() {
		defer close(channel)
		numStr := ""
		index := 0
		count := 0
		first := true
		for i, r := range s {
			if unicode.IsDigit(r) {
				if first {
					index = i
					first = false
				}
				numStr += string(r)
				count++
				if i == len(s)-1 {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					channel <- numIndex{num, index - 1, index + count}
				}
			} else {
				if numStr == "" {
					continue
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				channel <- numIndex{num, index - 1, index + count}
				index = 0
				numStr = ""
				count = 0
				first = true
			}
		}

	}()
	return channel
}

func main() {

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var lines []string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var lineNumbers = map[int][]numIndex{}

	// Add to lineNumbers map
	for i := 0; i < len(lines); i++ {
		empty := true
		for ni := range getNumbers(lines[i]) {
			lineNumbers[i] = append(lineNumbers[i], ni)
			empty = false
		}

		if empty {
			lineNumbers[i] = []numIndex{}
		}
	}

	total := 0

	// Go through all the symbols
	for i := 0; i < len(lines); i++ {
		for j, r := range lines[i] {
			if unicode.IsDigit(r) || r == '.' {
				continue
			}

			var numbers []int

			fmt.Println("Found symbol", string(r), "at line", i, "at position", j)

			// Is there a number above
			if i > 0 && len(lineNumbers[i-1]) > 0 {
				for _, ni := range lineNumbers[i-1] {
					if j >= ni.start && j <= ni.end {
						fmt.Printf("  Found gear above: %d\n", ni.number)
						numbers = append(numbers, ni.number)
					}
				}
			}

			// Is there a number below
			if i < len(lineNumbers)-1 && len(lineNumbers[i+1]) > 0 {
				for _, ni := range lineNumbers[i+1] {
					if j >= ni.start && j <= ni.end {
						fmt.Printf("  Found gear below: %d\n", ni.number)
						numbers = append(numbers, ni.number)
					}
				}
			}

			// Is there a line to the left
			if len(lineNumbers[i]) > 0 {
				for _, ni := range lineNumbers[i] {
					if j == ni.end {
						fmt.Printf("  Found gear to left: %d\n", ni.number)
						numbers = append(numbers, ni.number)
					}
				}
			}

			// Is there a line to the left
			if len(lineNumbers[i]) > 0 {
				for _, ni := range lineNumbers[i] {
					if j == ni.start {
						fmt.Printf("  Found gear to right: %d\n", ni.number)
						numbers = append(numbers, ni.number)
					}
				}
			}

			if len(numbers) != 2 {
				continue
			}

			total += numbers[0] * numbers[1]
		}

	}

	fmt.Printf("--\nTotal: %d\n", total)

	// Check if error reading file
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
