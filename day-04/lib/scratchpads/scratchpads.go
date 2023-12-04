package scratchpads

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Scratchpad struct {
	Cards []Card
}

func (s *Scratchpad) GetCard(id int) *Card {
	for _, card := range s.Cards {
		if card.ID == id {
			return &card
		}
	}

	return nil
}

func shit(cards []Card, cardID int) {
	card := &cards[cardID]
	matches := card.GetMatches()

	for i := cardID + 1; i < len(cards) && matches > 0; i++ {
		cards[i].Wins = append(cards[i].Wins, card.ID)
		matches--
		shit(cards, i) // Recursively distribute winnings to the next cards
	}
}

func (s *Scratchpad) CountWonScratchpads() int {
	log.Printf("before %+v", s.Cards)
	for i := 0; i < len(s.Cards); i++ {
		if s.Cards[i].GetMatches() > 0 {
			shit(s.Cards, i)
		}
	}
	sum := 0
	for _, v := range s.Cards {
		sum += len(v.Wins)
	}
	log.Printf("After %+v", s.Cards)

	log.Printf("out: %+v", sum)

	return sum + len(s.Cards)
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
		line = strings.Replace(line, "Card ", "", 1)
		parts := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Trim(parts[0], " "))
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
	ID             int
	Numbers        []int
	WinningNumbers []int
	Wins           []int
}

func (c *Card) GetPoints() int {
	count := c.GetMatches()
	if count == 1 {
		return 1
	}

	if count > 1 {
		return 1 << (count - 1)
	}

	return 0
}

func (c *Card) GetMatches() int {
	count := 0
	for _, num := range c.Numbers {
		if slices.Contains(c.WinningNumbers, num) {
			count++
		}
	}

	return count
}
