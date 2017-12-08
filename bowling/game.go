package bowling

type Game struct {
	Frames []Frame
}

type Frame interface {
	GetScore(g *Game) int
	GetSpareAddition() int
	GetStrikeAddition(g *Game) int
}

func NewGame(throws string) *Game {
	frames := Throws2Frames(throws)
	return &Game{
		Frames: frames,
	}
}

func (g *Game) GetNextFrame(f Frame) Frame {
	for i, frame := range g.Frames {
		if f == frame {
			return g.Frames[i+1]
		}
	}
	return RegularFrame{}
}

func Throws2Frames(throws string) []Frame {
	if len(throws) == 0 {
		return []Frame{}
	}
	if string(throws[0]) == "X" {
		return append([]Frame{NewStrike()}, Throws2Frames(throws[1:])...)
	}
	if len(throws) > 1 {
		if throws[1] == '/' {
			return append([]Frame{NewSpare(throws[0:1])}, Throws2Frames(throws[2:])...)
		}
		return append([]Frame{NewFrame(throws[:2])}, Throws2Frames(throws[2:])...)
	}
	return []Frame{NewFrame(throws)}
}

func (g *Game) Score() int {
	var sum int
	for i := 0; i < 10 && i < len(g.Frames); i++ {
		sum += g.Frames[i].GetScore(g)
	}
	return sum
}
