package scratchpads_test

import (
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-04/lib/scratchpads"
)

/*

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

*/

var input = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestScratchPads(t *testing.T) {
	expected := 13
	caller := scratchpads.NewScratchpad(input)

	if caller.GetPoints() != expected {
		t.Errorf("Expected %v, got %v", expected, caller.GetPoints())
	}


	card1Expected := 8
	card1Actual := caller.Cards[0].GetPoints()
	if card1Actual != card1Expected {
		t.Errorf("Expected %v, got %v", card1Expected, card1Actual)
	}
}

func TestScratchPadsPart2(t *testing.T) {
    expected := 30
    caller := scratchpads.NewScratchpad(input)
    if actual := caller.CountWonScratchpads(); actual != expected {
        t.Errorf("Expected %v, got %v", expected, actual)
    }

}
