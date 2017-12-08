package bowling

import "strconv"

type SpareFrame struct {
	Throw int
}

func NewSpare(throw1 string) Frame {
	score, _ := strconv.Atoi(throw1)
	return SpareFrame{
		Throw: score,
	}
}

func (f SpareFrame) GetScore(g *Game) int {
	next := g.GetNextFrame(f)
	return 10 + next.GetSpareAddition()
}

func (f SpareFrame) GetSpareAddition() int {
	return f.Throw
}

func (f SpareFrame) GetStrikeAddition(g *Game) int {
	return 10
}
