package bowling

import (
	"strconv"
)

type FrameType int

const (
	Regular FrameType = iota
	Spare
	Strike
)

type Frame struct {
	Throw1    int
	Throw2    int
	FrameType FrameType
}

func NewStrike() Frame {
	return Frame{Throw1: 10, FrameType: Strike}
}

func NewSpare(throw1 string) Frame {
	score, _ := strconv.Atoi(throw1)
	return Frame{
		Throw1:    score,
		Throw2:    10 - score,
		FrameType: Spare,
	}
}

func NewFrame(throws string) Frame {
	t1, _ := strconv.Atoi(string(throws[0]))
	frame := Frame{Throw1: t1}
	if len(throws) < 2 {
		return frame
	}
	t2, _ := strconv.Atoi(string(throws[1]))
	frame.Throw2 = t2
	return frame
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

func Score(throws string) int {
	frames := Throws2Frames(throws)
	var sum int
	for i := 0; i < 10 && i < len(frames); i++ {
		sum += frames[i].Throw1 + frames[i].Throw2
		if frames[i].FrameType == Spare {
			sum += frames[i+1].Throw1
		} else if frames[i].FrameType == Strike {
			sum += frames[i+1].Throw1
			if frames[i+1].FrameType == Strike {
				sum += frames[i+2].Throw1
			} else {
				sum += frames[i+1].Throw2
			}
		}
	}
	return sum
}
