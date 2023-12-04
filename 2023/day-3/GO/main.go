package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	sumPartOne := partOne(f)

	fmt.Println("The result of first part:", sumPartOne)
}

func partOne(f *os.File) int {
	scanner := bufio.NewScanner(f)
	scheme := [][]rune{}

	sum := 0
	for scanner.Scan() {
		scheme = append(scheme, []rune(scanner.Text()))
	}

	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			if scheme[i][j] != '.' && !unicode.IsDigit(scheme[i][j]) {
				sum += getNums(scheme, i, j)
			}
		}

	}
	return sum
}

func getNums(sh [][]rune, i, j int) int {
	sum := 0
	if i != 0 {
		for _, v := range getSideNums(sh, i-1, j) {
			sum += v
		}
	}
	for _, v := range getSideNums(sh, i, j) {
		sum += v
	}
	if i+1 < len(sh) {
		for _, v := range getSideNums(sh, i+1, j) {
			sum += v
		}
	}
	return sum
}

func getSideNums(sh [][]rune, i int, j int) []int {
	var ln, rn string

	for k := j + 1; k < len(sh[i]); k++ {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		rn += string(sh[i][k])
	}
	for k := j - 1; k >= 0; k-- {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		ln = string(sh[i][k]) + ln
	}

	if unicode.IsDigit(sh[i][j]) {
		n, _ := strconv.Atoi(ln + string(sh[i][j]) + rn)
		return []int{n}
	}

	num := []int{}
	lnn, _ := strconv.Atoi(ln)
	rnn, _ := strconv.Atoi(rn)

	if lnn != 0 {
		num = append(num, lnn)
	}

	if rnn != 0 {
		num = append(num, rnn)
	}

	return num
}
