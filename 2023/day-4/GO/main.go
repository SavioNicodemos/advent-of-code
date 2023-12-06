package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winnerNumbers []int
	myNumbers     []int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	cards := extractNumsFromFile(file)

	totalPoints := 0

	for _, card := range cards {
		totalPoints += calculatePoints(card.winnerNumbers, card.myNumbers)
	}

	fmt.Printf("The final result is: %d points\n", totalPoints)
}

func extractNumsFromFile(file *os.File) []Card {
	var cardNums = []Card{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numsPlusBar := strings.Split(scanner.Text(), ":")[1]
		nums := strings.Split(numsPlusBar, "|")
		trimmedWinners := strings.TrimSpace(strings.ReplaceAll(nums[0], "  ", " "))
		trimmedMine := strings.TrimSpace(strings.ReplaceAll(nums[1], "  ", " "))
		winnerNums := []int{}
		myNums := []int{}
		for _, num := range strings.Split(trimmedWinners, " ") {
			n, _ := strconv.Atoi(num)
			winnerNums = append(winnerNums, n)
		}
		for _, num := range strings.Split(trimmedMine, " ") {
			n, _ := strconv.Atoi(num)
			myNums = append(myNums, n)
		}
		cardNums = append(cardNums, Card{winnerNums, myNums})
	}
	return cardNums
}

func calculatePoints(winnerNums, myNums []int) int {
	winnerNumsMap := make(map[int]bool)
	for _, num := range winnerNums {
		winnerNumsMap[num] = true
	}

	cardPoints := 0
	for _, num := range myNums {
		if winnerNumsMap[num] {
			cardPoints++
		}
	}

	if cardPoints > 0 {
		return 1 << (cardPoints - 1) // 2^(n-1)
	}
	return 0
}
