package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

const filePath = "../../../data/dec3/testinput.txt"

type numIndex struct {
	number int
	index  int
}

// Generator function: get the numbers and index
func getNumbers(s string) <-chan numIndex {
	channel := make(chan numIndex)
	go func() {
		defer close(channel)
		numStr := ""
		index := 0
		for i, r := range s {
			if unicode.IsDigit(r) {
				if index == 0 {
					index = i
				}
				numStr += string(r)
			} else {
				if numStr == "" {
					continue
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				channel <- numIndex{num, index}
				index = 0
				numStr = ""
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

	for i := 0; i < len(lines); i++ {
		for ni := range getNumbers(lines[i]) {

		}
	}

	// Check if error reading file
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
