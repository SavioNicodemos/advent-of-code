package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var spelledNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		firstDigit, lastDigit := extractDigits(scanner.Text())
		sum += concatenateNumber(firstDigit, lastDigit)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sum)
}

func extractDigits(s string) (int, int) {
	firstDigit := extractFirstDigit(s)
	lastDigit := extractLastDigit(s)

	return firstDigit, lastDigit
}

func extractLastDigit(s string) int {
	for i := len(s); i >= 0; i-- {
		substr := s[i:]

		if substr == "" {
			continue
		}

		digit, err := strconv.Atoi(string(s[i]))
		if err == nil {
			return digit
		}

		for word, number := range spelledNumbers {
			if strings.Contains(substr, word) {
				return number
			}
		}
	}
	return 0
}

func extractFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		substr := s[:i+1]
		if substr == "" {
			continue
		}

		digit, err := strconv.Atoi(string(s[i]))
		if err == nil {
			return digit
		}

		for word, number := range spelledNumbers {
			if strings.Contains(substr, word) {
				return number
			}
		}
	}
	return 0
}

func concatenateNumber(num1 int, num2 int) int {
	return num1*10 + num2
}
