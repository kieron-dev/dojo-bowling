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
	var sum int
	for i := 0; i < len(throws); i++ {
		if i < len(throws)-1 && throws[i+1] == '/' {
			i++
			n, _ := strconv.Atoi(string(throws[i+1]))
			sum += 10 + n
		} else {
			n, _ := strconv.Atoi(string(throws[i]))
			sum += n
		}
	}
	return sum
}

func ParseFrames(throws string) []Frame {
	frames := []Frame{}
	for i := 0; i < len(throws); i++ {
		if throws[i] == 'X' {
			frames = append(frames, Frame{Throw1: 10, FrameType: Strike})
			continue
		}

		first, _ := strconv.Atoi(string(throws[i]))
		i++
		var next int

		frameType := Regular
		var newFrame Frame
		if throws[i] == '/' {
			next = 10 - first
			frameType = Spare
		} else {
			next, _ = strconv.Atoi(string(throws[i]))
		}
		newFrame = Frame{Throw1: first, Throw2: next, FrameType: frameType}
		frames = append(frames, newFrame)
	}
	return frames
}
