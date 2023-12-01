package trebuchet_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-01/lib/trebuchet"
)

func TestGetCalibrationNumberPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expectedSingleValues := []int{
		12, 38, 15, 77}
	expectedSum := 142
	cannon := trebuchet.NewTrebuchet(input, nil)
	for i, value := range input {
		if result := cannon.GetCalibrationNumber(value); result != expectedSingleValues[i] {
			t.Errorf("Expected %d, got %d", expectedSingleValues[i], result)
		}
	}

	sum := cannon.GetCalibrationSum()
	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestGetCalibrationNumberPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expectedSingleValues := []int{29, 83, 13, 24, 42, 14, 76}
	extension := map[string]int{
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
	expectedSum := 281

	cannon := trebuchet.NewTrebuchet(input, &extension)
	for i, value := range input {
		if result := cannon.GetCalibrationNumber(value); result != expectedSingleValues[i] {
			t.Errorf("Expected %d, got %d", expectedSingleValues[i], result)
		}
	}

	sum := cannon.GetCalibrationSum()
	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}

}

func TestSumPart1(t *testing.T) {
	fileName := "../../input.txt"
	expectedSum := 54390
	file, err := os.Open(fileName)
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
	cannon := trebuchet.NewTrebuchet(input, nil)
	sum := cannon.GetCalibrationSum()
	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestStupid(t *testing.T) {
	input := "ftwo91"
	expect := 91
	cannon := trebuchet.NewTrebuchet([]string{input}, nil)
	result := cannon.GetCalibrationNumber(input)
	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}
