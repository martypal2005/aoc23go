package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const inputFile = "../../../data/dec5/input.txt"

type destStruct struct {
	destination int
	source      int
	count       int
}

type seedLocationStruct struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type destinations []destStruct

func (a destinations) Len() int           { return len(a) }
func (a destinations) Less(i, j int) bool { return a[i].source < a[j].source }
func (a destinations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getSourceSlice(d destinations) []int {
	var numbers []int
	for i := 0; i < len(d); i++ {
		numbers = append(numbers, d[i].source)
	}
	return numbers
}

// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
func printSeedLocation(seedLocation seedLocationStruct) {
	fmt.Printf(
		"Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d.\n",
		seedLocation.seed,
		seedLocation.soil,
		seedLocation.fertilizer,
		seedLocation.water,
		seedLocation.light,
		seedLocation.temperature,
		seedLocation.humidity,
		seedLocation.location,
	)
}

func findDestination(source int, d destinations) int {

	var match destStruct

	for _, dest := range d {
		if source >= dest.source {
			match = dest
		}
	}

	if source < match.source+match.count {
		diff := source - match.source
		return match.destination + diff
	}

	return source
}

func getDestinations(lines []string) destinations {
	var toReturn destinations

	for _, line := range lines {
		numbers := getNumbers(line)
		if len(numbers) != 3 {
			fmt.Println("Invalid numbers:", numbers)
		}

		var sourceToDestination destStruct
		sourceToDestination.destination = numbers[0]
		sourceToDestination.source = numbers[1]
		sourceToDestination.count = numbers[2]

		toReturn = append(toReturn, sourceToDestination)
	}

	sort.Sort(toReturn)

	return toReturn
}

func getLines(input []string, index int) []string {
	var lines []string
	for i := index; i < len(input); i++ {
		if input[i] == "" {
			break
		}
		lines = append(lines, input[i])
	}
	return lines
}

func getNumbers(line string) []int {
	var numbers []int
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)
	if matches == nil {
		return numbers
	}
	for _, match := range matches {
		number, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var seeds []int
	var soil []string
	var fertilizer []string
	var water []string
	var light []string
	var temperature []string
	var humidity []string
	var location []string

	var seedToSoil destinations
	var soilToFertilizer destinations
	var fertilizerToWater destinations
	var waterToLight destinations
	var lightToTemperature destinations
	var temperatureToHumidity destinations
	var humidityToLocation destinations

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "seeds") {
			seeds = getNumbers(lines[i])
		} else if strings.Contains(lines[i], "seed-to-soil") {
			soil = getLines(lines, i+1)
			seedToSoil = getDestinations(soil)
		} else if strings.Contains(lines[i], "soil-to-fertilizer") {
			fertilizer = getLines(lines, i+1)
			soilToFertilizer = getDestinations(fertilizer)
		} else if strings.Contains(lines[i], "fertilizer-to-water") {
			water = getLines(lines, i+1)
			fertilizerToWater = getDestinations(water)
		} else if strings.Contains(lines[i], "water-to-light") {
			light = getLines(lines, i+1)
			waterToLight = getDestinations(light)
		} else if strings.Contains(lines[i], "light-to-temperature") {
			temperature = getLines(lines, i+1)
			lightToTemperature = getDestinations(temperature)
		} else if strings.Contains(lines[i], "temperature-to-humidity") {
			humidity = getLines(lines, i+1)
			temperatureToHumidity = getDestinations(humidity)
		} else if strings.Contains(lines[i], "humidity-to-location") {
			location = getLines(lines, i+1)
			humidityToLocation = getDestinations(location)
		}
	}

	// fmt.Println(seeds)
	// fmt.Println(getSourceSlice(seedToSoil))
	// fmt.Println(getSourceSlice(soilToFertilizer))
	// fmt.Println(getSourceSlice(fertilizerToWater))
	// fmt.Println(getSourceSlice(waterToLight))
	// fmt.Println(getSourceSlice(lightToTemperature))
	// fmt.Println(getSourceSlice(temperatureToHumidity))
	// fmt.Println(getSourceSlice(humidityToLocation))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var seedLocations []seedLocationStruct
	smallestLocation := -1

	for _, seed := range seeds {

		var seedLocation seedLocationStruct

		seedLocation.seed = seed
		seedLocation.soil = findDestination(seed, seedToSoil)
		seedLocation.fertilizer = findDestination(seedLocation.soil, soilToFertilizer)
		seedLocation.water = findDestination(seedLocation.fertilizer, fertilizerToWater)
		seedLocation.light = findDestination(seedLocation.water, waterToLight)
		seedLocation.temperature = findDestination(seedLocation.light, lightToTemperature)
		seedLocation.humidity = findDestination(seedLocation.temperature, temperatureToHumidity)
		seedLocation.location = findDestination(seedLocation.humidity, humidityToLocation)

		seedLocations = append(seedLocations, seedLocation)

		printSeedLocation(seedLocation)

		if smallestLocation == -1 || seedLocation.location < smallestLocation {
			smallestLocation = seedLocation.location
		}
	}

	fmt.Printf("Smallest location: %d\n", smallestLocation)

}
