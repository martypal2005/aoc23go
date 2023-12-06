package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		line = strings.Replace(line, "one", "o1e", -1)
		line = strings.Replace(line, "two", "t2o", -1)
		line = strings.Replace(line, "three", "t3e", -1)
		line = strings.Replace(line, "four", "f4r", -1)
		line = strings.Replace(line, "five", "f5e", -1)
		line = strings.Replace(line, "six", "s6x", -1)
		line = strings.Replace(line, "seven", "s7n", -1)
		line = strings.Replace(line, "eight", "e8t", -1)
		line = strings.Replace(line, "nine", "n9e", -1)

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
