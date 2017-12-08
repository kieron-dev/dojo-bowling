package bowling

import "strconv"

type RegularFrame struct {
	Throw1 int
	Throw2 int
}

func NewFrame(throws string) Frame {
	t1, _ := strconv.Atoi(string(throws[0]))
	frame := RegularFrame{Throw1: t1}
	if len(throws) < 2 {
		return frame
	}
	t2, _ := strconv.Atoi(string(throws[1]))
	frame.Throw2 = t2
	return frame
}

func (f RegularFrame) GetScore(g *Game) int {
	return f.Throw1 + f.Throw2
}

func (f RegularFrame) GetSpareAddition() int {
	return f.Throw1
}

func (f RegularFrame) GetStrikeAddition(g *Game) int {
	return f.Throw1 + f.Throw2
}
