package playground

import (
	"github.com/ryodeushii/advent-of-code-2023/day-02/lib/record"
)

type Playground struct {
	Records []record.Record

	Dimes map[string]int
}

func NewPlayground(input []string, dimes map[string]int) *Playground {
	items := make([]record.Record, 0)

	for _, str := range input {
		recordItem := record.NewRecord(str)
		recordItem.Parse()

		items = append(items, *recordItem)
	}

	return &Playground{
		Records: items,
		Dimes:   dimes,
	}
}

func (p *Playground) SumIdsOfPossibleRecords() int {
	sum := 0

	for _, r := range p.Records {
		if r.IsPossible(p.Dimes) {
			sum += r.ID
		}
	}

	return sum
}

func (p *Playground) SumOfPowers() int {
    sum := 0

    for _, r := range p.Records {
        sum += r.GetPower()
    }

   return sum
}
