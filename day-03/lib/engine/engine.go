package engine

import (
	"log"
	"strconv"
)

type Engine [][]byte

func NewEngine(input []string) *Engine {
	e := Engine{}
	for _, line := range input {
		row := []byte(line)
		e = append(e, row)
	}
	return &e
}

var numbers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var ignored = []byte{'.', '\n', '\t', ' '}

func (e *Engine) Loop() int {
	part_numbers := make([]int, 0)
	for y := 0; y < len(*e); y++ {
		row := (*e)[y]
		for x := 0; x < len(row); x++ {
			cell := row[x]
			// if cell is number - find whole sequence of numbers up to any non-number character
			if isNumber(cell) {
				numeric_sequence := make([]byte, 0)
				lastIndex := x
				for i := x; i < len(row); i++ {
					if isNumber(row[i]) {
						numeric_sequence = append(numeric_sequence, row[i])
						lastIndex = i
					} else {
						break
					}
				}
				symbolsCount := 0
				for i := x; i < lastIndex+1; i++ {
					// current position is (i, y)
					// if any symbol is around this cell( include corners) - increase symbolsCount
					// if symbolsCount > 0 - break
					var (
						left, right, top, bottom, topLeft, topRight, bottomLeft, bottomRight byte
					)
					if i > 0 {
						left = row[i-1]
					} else {
						left = '.'
					}
					if i < len(row)-1 {
						right = row[i+1]
					} else {
						right = '.'
					}
					if y > 0 {
						top = (*e)[y-1][i]
					} else {
						top = '.'
					}
					if y < len(*e)-1 {
						bottom = (*e)[y+1][i]
					} else {
						bottom = '.'
					}
					if y > 0 && i > 0 {
						topLeft = (*e)[y-1][i-1]
					} else {
						topLeft = '.'
					}
					if y > 0 && i < len(row)-1 {
						topRight = (*e)[y-1][i+1]
					} else {
						topRight = '.'
					}
					if y < len(*e)-1 && i > 0 {
						bottomLeft = (*e)[y+1][i-1]
					} else {
						bottomLeft = '.'
					}
					if y < len(*e)-1 && i < len(row)-1 {
						bottomRight = (*e)[y+1][i+1]
					} else {
						bottomRight = '.'
					}
					//log.Printf("%c %c %c\n%c %c %c\n%c %c %c\n", topLeft, top, topRight, left, cell, right, bottomLeft, bottom, bottomRight)
					for _, symbol := range []byte{topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight} {
						if isSymbol(symbol) {
							symbolsCount++
							break
						}
					}

					if symbolsCount > 0 {
						break
					}

				}
				x += len(numeric_sequence)
				if len(numeric_sequence) > 0 && symbolsCount > 0 {
					number := string(numeric_sequence)
					value, _ := strconv.Atoi(number)
					part_numbers = append(part_numbers, value)
				}
			}

		}
	}
    log.Println("Part numbers:", part_numbers)
    sum := 0
    for _, n := range part_numbers {
        sum += n
    }
    log.Println("Sum:", sum)
    return sum
}

func isNumber(c byte) bool {
	for _, n := range numbers {
		if c == n {
			return true
		}
	}
	return false
}

func isLetter(c byte) bool {
	for _, n := range letters {
		if c == n {
			return true
		}
	}
	return false
}

func isIgnored(c byte) bool {
	for _, n := range ignored {
		if c == n {
			return true
		}
	}
	return false
}

func isSymbol(c byte) bool {
	if !isNumber(c) && !isLetter(c) && !isIgnored(c) {
		return true
	}
	return false
}
