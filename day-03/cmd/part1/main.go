package main

import (
	"bufio"
	"log"
	"os"

	"github.com/ryodeushii/advent-of-code-2023/day-03/lib/engine"
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

	machine := engine.NewEngine(input)
	sum := machine.FindPartNumbers()
	log.Printf("Sum of part numbers is %d", sum)

}
