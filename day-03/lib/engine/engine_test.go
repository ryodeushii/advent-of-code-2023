package engine_test

import (
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-03/lib/engine"
)

func TestEngineLoop(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	machine := engine.NewEngine(input)

    expect := 4361
    sum := machine.Loop()
    if sum != expect {
        t.Errorf("expect %d, but got %d", expect, sum)
    }

}
