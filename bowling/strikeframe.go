package bowling

type StrikeFrame struct{}

func NewStrike() Frame {
	return StrikeFrame{}
}

func (f StrikeFrame) GetScore(g *Game) int {
	next := g.GetNextFrame(f)
	return 10 + next.GetStrikeAddition(g)
}

func (f StrikeFrame) GetSpareAddition() int {
	return 10
}

func (f StrikeFrame) GetStrikeAddition(g *Game) int {
	next := g.GetNextFrame(f)
	return 10 + next.GetSpareAddition()
}
