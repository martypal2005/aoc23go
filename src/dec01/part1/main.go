package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const filePath = "../../../data/dec1/input.txt"

func findFirstAndLastDigit(s string) (number int, found bool) {
	var first, last rune
	for _, r := range s {
		if unicode.IsDigit(r) {
			if !found {
				first = r
				found = true
			}
			last = r
		}
	}

	if found {
		stringNumber := string(first) + string(last)
		tmpNum, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		number = tmpNum
	}

	return number, found
}

func main() {

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create scanner object to read file
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		number, found := findFirstAndLastDigit(line)

		if found {
			total += number
		}
	}

	// Check for errors during Scan.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)

}
