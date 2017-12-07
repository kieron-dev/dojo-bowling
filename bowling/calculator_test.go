package bowling_test

import (
	"github.com/kieron-pivotal/dojo-bowling/bowling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	It("can calculate scores if you never frame more than 9", func() {
		Expect(bowling.Score("--")).To(Equal(0))
		Expect(bowling.Score("1-")).To(Equal(1))
		Expect(bowling.Score("2-")).To(Equal(2))
		Expect(bowling.Score("12")).To(Equal(3))
		Expect(bowling.Score("123-")).To(Equal(6))
		Expect(bowling.Score("X12")).To(Equal(16))
		Expect(bowling.Score("XXXXXXXXXXXX")).To(Equal(300))
		Expect(bowling.Score("9-9-9-9-9-9-9-9-9-9-")).To(Equal(90))
		Expect(bowling.Score("5/5/5/5/5/5/5/5/5/5/5")).To(Equal(150))
	})

	It("can calculate scores with non-final spares", func() {
		Expect(bowling.Score("3/4-")).To(Equal(18))
	})

	It("can parse inputs into framelists", func() {
		Expect(bowling.Throws2Frames("34")).To(Equal([]bowling.Frame{
			{Throw1: 3, Throw2: 4, FrameType: bowling.Regular},
		}))
		Expect(bowling.Throws2Frames("-4")).To(Equal([]bowling.Frame{
			{Throw1: 0, Throw2: 4, FrameType: bowling.Regular},
		}))
		Expect(bowling.Throws2Frames("3/")).To(Equal([]bowling.Frame{
			{Throw1: 3, Throw2: 7, FrameType: bowling.Spare},
		}))

		Expect(bowling.Throws2Frames("225/")).To(Equal([]bowling.Frame{
			{Throw1: 2, Throw2: 2, FrameType: bowling.Regular},
			{Throw1: 5, Throw2: 5, FrameType: bowling.Spare},
		}))

		Expect(bowling.Throws2Frames("X")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
		Expect(bowling.Throws2Frames("XX")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
		Expect(bowling.Throws2Frames("X72X")).To(Equal([]bowling.Frame{
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
			{Throw1: 7, Throw2: 2, FrameType: bowling.Regular},
			{Throw1: 10, Throw2: 0, FrameType: bowling.Strike},
		}))
	})

	It("Deals with a single number in a frame", func() {
		Expect(bowling.NewFrame("5")).To(Equal(bowling.Frame{
			Throw1: 5,
		}))
	})
})
