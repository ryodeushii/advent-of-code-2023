package game_test

import (
	"testing"

	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/game"
)

var dimes = map[string]int{
    "red":   12,
    "green": 13,
    "blue":  14,
}

func TestParse(t *testing.T) {
	gameInstance := game.NewGame("3 blue, 4 red")
	gameInstance.Parse()

    if gameInstance.Parsed["blue"] != 3 {
        t.Errorf("Game should have 3 blue, got %d", gameInstance.Parsed["blue"])
    }

    if gameInstance.Parsed["red"] != 4 {
        t.Errorf("Game should have 4 red, got %d", gameInstance.Parsed["red"])
    }

    if !gameInstance.IsPossible(dimes) {
        t.Errorf("Game should be possible")
    }

}
