package main

import (
	"bufio"
	"fmt"

	"log"
	"os"

	"github.com/ryodeushii/advent-of-code-2023/day-01/lib/trebuchet"
)

func main() {
	addDigits := map[string]int{
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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	trebuchet := trebuchet.NewTrebuchet(lines, &addDigits)
	fmt.Println(trebuchet.GetCalibrationSum())
}
