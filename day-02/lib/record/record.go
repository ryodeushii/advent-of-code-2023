package record

import (
	"strconv"
	"strings"

	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/game"
)

type Record struct {
	ID     int
	Input  string
	Parsed []game.Game
}

func NewRecord(input string) *Record {
	return &Record{
		Input:  input,
		Parsed: make([]game.Game, 0),
		ID:     -1,
	}
}

func (r *Record) Parse() {
	initial := strings.Split(r.Input, ": ")

	idpair := strings.Split(initial[0], " ")
	id, err := strconv.Atoi(idpair[1])
	if err != nil {
		panic(err)
	}
	r.ID = id

	gamesStrings := strings.Split(initial[1], "; ")
	for _, gameString := range gamesStrings {
		gameItem := game.NewGame(gameString)
		gameItem.Parse()
		r.Parsed = append(r.Parsed, *gameItem)
	}
}

func (r *Record) IsPossible(dimes map[string]int) bool {
    for _, game := range r.Parsed {
        if !game.IsPossible(dimes) {
            return false
        }
    }
    return true
}
