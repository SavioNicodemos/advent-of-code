package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var digitsPerLine [][]int
	for _, line := range lines {
		var digits []int

		for _, digit := range line {
			if digit >= '0' && digit <= '9' {
				number, _ := strconv.Atoi(string(digit))
				digits = append(digits, number)
			}
		}

		digitsPerLine = append(digitsPerLine, digits)
	}

	var sum int = 0
	for _, digits := range digitsPerLine {
		if len(digits) == 0 {
			continue
		}
		if len(digits) == 1 {
			sum += concatenateNumber(digits[0], digits[0])
		}
		if len(digits) > 1 {
			sum += concatenateNumber(digits[0], digits[len(digits)-1])
		}
	}

	fmt.Println(sum)
	return
}

func concatenateNumber(num1 int, num2 int) int {
	return num1*10 + num2
}
