package bowling_test

import (
	"github.com/kieron-pivotal/dojo-bowling/bowling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	It("can calculate scores if you never frame more than 9", func() {
		Expect(bowling.Score("-")).To(Equal(0))
		Expect(bowling.Score("1")).To(Equal(1))
		Expect(bowling.Score("2")).To(Equal(2))
		Expect(bowling.Score("12")).To(Equal(3))
		Expect(bowling.Score("123")).To(Equal(6))
	})

	It("can calculate scores with non-final spares", func() {
		Expect(bowling.Score("3/4")).To(Equal(18))
	})

	It("can parse inputs into framelists", func() {
		Expect(bowling.ParseFrames("34")).To(Equal([]bowling.Frame{
			{Throw1: 3, Throw2: 4, FrameType: bowling.Regular},
		}))
		Expect(bowling.ParseFrames("-4")).To(Equal([]bowling.Frame{
			{Throw1: 0, Throw2: 4, FrameType: bowling.Regular},
		}))
		Expect(bowling.ParseFrames("3/")).To(Equal([]bowling.Frame{
			{Throw1: 3, Throw2: 7, FrameType: bowling.Spare},
		}))

		Expect(bowling.ParseFrames("225/")).To(Equal([]bowling.Frame{
			{Throw1: 2, Throw2: 2, FrameType: bowling.Regular},
			{Throw1: 5, Throw2: 5, FrameType: bowling.Spare},
		}))

		Expect(bowling.ParseFrames("X")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
		Expect(bowling.ParseFrames("XX")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
		Expect(bowling.ParseFrames("X72X")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
			{Throw1: 7, Throw2: 2, FrameType: bowling.Regular},
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
	})
})
