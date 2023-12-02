package main

import (
	"bufio"
	"log"
	"os"

	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/playground"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file", fileName)
		return
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	dimes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	playgroundInstance := playground.NewPlayground(input, dimes)
	sum := playgroundInstance.SumOfPowers()

	log.Printf("Sum of powers %d", sum)
}
