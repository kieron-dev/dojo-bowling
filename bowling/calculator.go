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
