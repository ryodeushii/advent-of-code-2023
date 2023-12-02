package record_test

import (
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/record"
)

var dimes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func TestParse(t *testing.T) {
	record := record.NewRecord("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	record.Parse()

	if record.ID != 1 {
		t.Errorf("ID should be 1, got %d", record.ID)
	}

	if len(record.Parsed) != 3 {
		t.Errorf("Parsed should have 3 games, got %d", len(record.Parsed))
	}

	if record.Parsed[0].Parsed["blue"] != 3 {
		t.Errorf("Game 1 should have 3 blue, got %d", record.Parsed[0].Parsed["blue"])
	}

	if record.Parsed[0].Parsed["red"] != 4 {
		t.Errorf("Game 1 should have 4 red, got %d", record.Parsed[0].Parsed["red"])
	}

	if record.Parsed[1].Parsed["red"] != 1 {
		t.Errorf("Game 2 should have 1 red, got %d", record.Parsed[1].Parsed["red"])
	}

	if !record.IsPossible(dimes) {
		t.Errorf("Record should be possible")
	}

    setOfDimes := record.FindMinimumSetOfDimes()
    if setOfDimes["red"] != 4 {
        t.Errorf("Set of dimes should have 4 red, got %d", setOfDimes["red"])
    }
    if setOfDimes["blue"] != 6 {
        t.Errorf("Set of dimes should have 6 blue, got %d", setOfDimes["blue"])
    }
    if setOfDimes["green"] != 2 {
        t.Errorf("Set of dimes should have 2 green, got %d", setOfDimes["green"])
    }
}

func TestImpossible(t *testing.T) {
	record := record.NewRecord("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	record.Parse()

	if record.ID != 3 {
		t.Errorf("ID should be 3, got %d", record.ID)
	}

	if record.IsPossible(dimes) {
		t.Errorf("Record should not be possible")
	}
}
