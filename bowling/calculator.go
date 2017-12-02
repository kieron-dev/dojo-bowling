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

func Score(throws string) int {
	frames := ParseFrames(throws)
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

func ParseFrames(throws string) []Frame {
	frames := []Frame{}
	for len(throws) > 0 {
		if throws[0] == 'X' {
			frames = append(frames, Frame{Throw1: 10, FrameType: Strike})
			throws = throws[1:]
			continue
		}

		first, _ := strconv.Atoi(string(throws[0]))
		var next int

		frameType := Regular
		var newFrame Frame
		if len(throws) > 1 {
			if throws[1] == '/' {
				next = 10 - first
				frameType = Spare
			} else {
				next, _ = strconv.Atoi(string(throws[1]))
			}
			throws = throws[2:]
		} else {
			throws = ""
		}
		newFrame = Frame{Throw1: first, Throw2: next, FrameType: frameType}
		frames = append(frames, newFrame)
	}
	return frames
}
