package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filePath = "../../../data/dec2/input.txt"

func getNumColor(s string) (int, string) {
	parts := strings.Split(s, " ")

	if len(parts) != 2 {
		return 0, ""
	}

	count, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, ""
	}

	return count, parts[1]
}

func main() {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, ":")

		if len(tokens) != 2 {
			fmt.Printf("Imperly formatted line: %s\n", line)
		}

		newString := strings.ReplaceAll(tokens[1], ";", ",")

		sections := strings.Split(newString, ",")

		colorMax := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, section := range sections {
			trimmed := strings.TrimSpace(section)
			count, color := getNumColor(trimmed)
			if count > colorMax[color] {
				colorMax[color] = count
			}
		}

		total += colorMax["red"] * colorMax["green"] * colorMax["blue"]
	}

	fmt.Printf("--\nSum of powers: %d\n", total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
