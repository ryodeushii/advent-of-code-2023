package playground_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/playground"
)

func TestNewPlayground(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	dimes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	playgroundInstance := playground.NewPlayground(input, dimes)

	if len(playgroundInstance.Records) != 5 {
		t.Errorf("Playground should have 5 records, got %d", len(playgroundInstance.Records))
	}

	expectedSum := 8
	if sum := playgroundInstance.SumIdsOfPossibleRecords(); sum != expectedSum {
		t.Errorf("Sum of ids of possible records should be %d, got %d", expectedSum, sum)
	}
}

func TestOnRealData(t *testing.T) {
	fileName := "../../input.txt"
	file, err := os.Open(fileName)
    expectedSum := 2256
	if err != nil {
		t.Errorf("Failed to open file %s", fileName)
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
    sum := playgroundInstance.SumIdsOfPossibleRecords()
    if sum != expectedSum {
        t.Errorf("Sum of ids of possible records should be %d, got %d", expectedSum, sum)
    }

}
