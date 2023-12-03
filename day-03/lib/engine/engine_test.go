package engine_test

import (
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-03/lib/engine"
)

var input = []string{
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

var machine = engine.NewEngine(input)

func TestEngineLoop(t *testing.T) {
	expect := 4361
	sum := machine.FindPartNumbers()
	if sum != expect {
		t.Errorf("expect %d, but got %d", expect, sum)
	}

}

func TestGears(t *testing.T) {
    expected:= 467835
    sum :=	machine.FindGears()

    if sum != expected {
        t.Errorf("expect %d, but got %d", expected, sum)
    }
}
