package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const filePath = "../../../data/dec3/testinput.txt"

type numIndex struct {
	number int
	start  int
	end    int
}

// Check if symbol
func isSymbol(s string, index int) bool {
	r := []rune(s)
	if index < 0 || index == len(r) {
		return false // or handle error
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

	total := 0

	for i := 0; i < len(lines); i++ {
		for ni := range getNumbers(lines[i]) {

			// Is there a symbol to the left?
			if isSymbol(lines[i], ni.start) {
				total += ni.number
				continue
			}

			// Is there a symbol to the right?
			if isSymbol(lines[i], ni.end) {
				total += ni.number
				continue
			}
			/*

				// Check all indexes in line above
				if i > 0 {
					isPart := false
					for j := ni.start; j < ni.end; j++ {
						if isSymbol(lines[i-1], i) {
							total += ni.number
							isPart = true
							break
						}
					}
					if isPart {
						continue
					}
				}

					// Check all indexes in line below
					if i < len(lines) {
						isPart := false
						for j := ni.start; j < ni.end; j++ {
							if isSymbol(lines[i+1], i) {
								total += ni.number
								isPart = true
								break
							}
						}
						if isPart {
							continue
						}
					}
			*/

		}
	}

	fmt.Println(total)

	// Check if error reading file
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
