package game

import (
	"strconv"
	"strings"
)

type Game struct {
	Input  string
	Parsed map[string]int
}

func NewGame(input string) *Game {
	return &Game{
		Input:  input,
		Parsed: make(map[string]int),
	}
}

func (g *Game) Parse() {
	pairs := strings.Split(g.Input, ", ")
	for _, pair := range pairs {
		split := strings.Split(pair, " ")
		key := split[1]
		value, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		g.Parsed[key] = value
	}
}


func (g *Game) IsPossible(dimes map[string]int) bool {
    for key, value := range g.Parsed {
        if dimes[key] < value {
            return false
        }
    }
    return true
}
