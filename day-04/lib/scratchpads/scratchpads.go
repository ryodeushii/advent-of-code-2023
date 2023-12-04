package scratchpads

import (
	"slices"
	"strconv"
	"strings"
)

type Scratchpad struct {
	Cards []Card
}

func (s *Scratchpad) GetPoints() int {
	sum := 0
	for _, card := range s.Cards {
		sum += card.GetPoints()
	}

	return sum
}

func NewScratchpad(input []string) *Scratchpad {
	cards := make([]Card, len(input))
	for i, line := range input {
		//		trimmed := strings.Replace(line, "  ", "", -1)
		parts := strings.Split(line, ":")
		id := parts[0]
		parts = strings.Split(strings.Trim(parts[1], "   "), " | ")
		numbers_str := strings.Split(parts[0], " ")
		winningNumbers_str := strings.Split(parts[1], " ")
		numbers := make([]int, 0)
		winningNumbers := make([]int, 0)
		for _, number_str := range numbers_str {
			if number_str != "" {
				number, _ := strconv.Atoi(number_str)
				numbers = append(numbers, number)
			}
		}
		for _, number_str := range winningNumbers_str {
			if number_str != "" {
				number, _ := strconv.Atoi(number_str)
				winningNumbers = append(winningNumbers, number)
			}
		}
		cards[i] = Card{
			ID:             id,
			Numbers:        numbers,
			WinningNumbers: winningNumbers,
		}
	}

	return &Scratchpad{Cards: cards}
}

type Card struct {
	ID             string
	Numbers        []int
	WinningNumbers []int
}

func (c *Card) GetPoints() int {
	count := 0
	for _, num := range c.Numbers {
		if slices.Contains(c.WinningNumbers, num) {
			count++
		}
	}

	if count == 1 {
		return 1
	}

	if count > 1 {
		return 1 << (count - 1)
	}

	return 0
}
