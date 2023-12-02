package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
	digits := regexp.MustCompile("[0-9]").FindAllString(s, -1)
	firstDigit, _ := strconv.Atoi(digits[0])
	lastDigit, _ := strconv.Atoi(digits[len(digits)-1])
	return firstDigit, lastDigit
}

func concatenateNumber(num1 int, num2 int) int {
	return num1*10 + num2
}
