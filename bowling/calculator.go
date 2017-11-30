package bowling

import (
	"strconv"
)

func Score(throws string) int {
	var sum int
	for _, r := range throws {
		i, _ := strconv.Atoi(string(r))
		sum += i
	}
	return sum
}
