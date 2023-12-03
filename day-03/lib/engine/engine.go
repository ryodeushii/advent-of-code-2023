package engine

import (
	"log"
	"slices"
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

func (e *Engine) FindPartNumbers() int {
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

type Gear struct {
	Left  int
	Right int
}

func NewGear(left, right int) *Gear {
	return &Gear{Left: left, Right: right}
}

func (g *Gear) Ratio() int {
	return g.Left * g.Right
}

func (e *Engine) FindGears() int {
	gears := make([]Gear, 0)
	for y := 0; y < len(*e); y++ {
		row := (*e)[y]
		for x := 0; x < len(row); x++ {
			cell := row[x]
			if isAsterisk(cell) {
				// if cell is asterisk:
				// 0. for any numeric neighbour cell (include corners) - find numeric sequences
				// 1. find all neighbouring numeric sequences around this cell (there may be from 0 to 2 sequences)
				// 2. if there are 2 numeric sequences - create gear and push to array
				// 3. if there is 1 or 0 numeric sequences - ignore this cell

				// 0. for any numeric neighbour cell (include corners) - find numeric sequences
				numeric_sequences := make([]string, 0)
				// check if any neighbour cell is numeric and then find whole number sequences
				leftx := x - 1
				lefty := y
				rightx := x + 1
				righty := y
				topx := x
				topy := y - 1
				bottomx := x
				bottomy := y + 1
				coords := [][]int{
					{leftx, lefty},
					{rightx, righty},
					{topx, topy},
					{bottomx, bottomy},
					{leftx, topy},
					{rightx, topy},
					{leftx, bottomy},
					{rightx, bottomy},
				}
				for _, pair := range coords {
					if value := e.FindSequenceFrom(pair[0], pair[1]); value != nil {
						numeric_sequences = append(numeric_sequences, *value)
					}
				}

				slices.Sort(numeric_sequences)
                slices.Reverse(numeric_sequences)
                numeric_sequences = slices.Compact(numeric_sequences)

				log.Println("Numeric sequences:", numeric_sequences)
				if len(numeric_sequences) == 2 {
					left, _ := strconv.Atoi(string(numeric_sequences[0]))
					right, _ := strconv.Atoi(string(numeric_sequences[1]))
					gear := NewGear(left, right)
					gears = append(gears, *gear)
				}

			}
		}
	}
	log.Println("Gears:", gears)

	sum := 0
	for _, g := range gears {
		sum += g.Ratio()
	}

	return sum
}

func (e *Engine) FindSequenceFrom(x, y int) *string {
	if x < 0 || y < 0 || x > len((*e)[0])-1 || y > len(*e)-1 {
		return nil
	}

	cell := (*e)[y][x]
	if isNumber(cell) {
		// if cell is numeric - find whole sequence of numbers up to any non-number character to both sides of cell
		numeric_sequence := make([]byte, 0)
		for x1 := x; x1 >= 0; x1-- {
			if isNumber((*e)[y][x1]) {
				numeric_sequence = append(numeric_sequence, (*e)[y][x1])
			} else {
				break
			}
		}
		slices.Reverse(numeric_sequence)
		for x2 := x + 1; x2 < len((*e)[0]); x2++ {
			if isNumber((*e)[y][x2]) {
				numeric_sequence = append(numeric_sequence, (*e)[y][x2])
			} else {
				break
			}
		}
		if len(numeric_sequence) > 0 {
			number := string(numeric_sequence)
			return &number
		}
	}

	return nil
}

func isAsterisk(c byte) bool {
	return c == '*'
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
