package trebuchet

import (
	"log"
	"strings"
)

type Trebuchet struct {
	input       []string
	validDigits map[string]int
}

func NewTrebuchet(input []string, addDigits *map[string]int) Trebuchet {
	var validDigits map[string]int
	defaultDigits := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	if addDigits == nil || len(*addDigits) == 0 {
		validDigits = defaultDigits
	} else {
		validDigits = *addDigits
		for key, value := range defaultDigits {
			validDigits[key] = value
		}
	}
	return Trebuchet{input, validDigits}
}

func (t Trebuchet) GetCalibrationNumber(input string) int {
	first := t.getFirstDigit(input)
	last := t.getLastDigit(input)
	value := first*10 + last
	log.Printf("[%s]:%d", input, value)
	return value
}

func (t Trebuchet) GetCalibrationSum() int {
	sum := 0
	for _, value := range t.input {
		sum += t.GetCalibrationNumber(value)
	}
	return sum

}

func (t Trebuchet) getFirstDigit(input string) int {
	index := len(input)
	outValue := 0
	for key, value := range t.validDigits {
		idx := strings.Index(input, key)
		if idx == -1 {
			continue
		}
		if idx <= index {
			index = idx
			outValue = value
		}
	}
	return outValue
}

func (t Trebuchet) getLastDigit(input string) int {
	index := 0
	outValue := 0
	for key, value := range t.validDigits {
		idx := strings.LastIndex(input, key)
		if idx == -1 {
			continue
		}
		if idx >= index {
			index = idx
			outValue = value
		}
	}
	return outValue
}
