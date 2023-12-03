package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var possibleConfigs = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := readInput(("input.txt"))

	partOneResult := partOne(lines)

	fmt.Println("Part one result:", partOneResult)
}

func readInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(rawFile), "\n")
}

/*
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

func getValues(line string) map[string]int {
	values := strings.Split(line, ":")
	colorValues := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	colorString := strings.Split(values[1], ";")
	for _, set := range colorString {
		items := strings.Split(set, ",")
		for _, item := range items {
			value := strings.Split(item, " ")

			colorValue := value[1]
			colorName := value[2]

			number, _ := strconv.Atoi(colorValue)

			colorValues[colorName] = getMaxValue(colorValues[colorName], number)
		}
	}
	return colorValues
}

func getMaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partOne(lines []string) int {
	var idSum int

	for id, line := range lines {
		lineValues := getValues(line)

		shouldSum := true
		for color, value := range lineValues {
			if value > possibleConfigs[color] {
				shouldSum = false
			}
		}

		if shouldSum {
			idSum += id + 1
		}
	}
	return idSum
}
