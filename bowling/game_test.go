package bowling_test

import (
	"github.com/kieron-pivotal/dojo-bowling/bowling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	It("can calculate scores if you never frame more than 9", func() {
		Expect(bowling.NewGame("--").Score()).To(Equal(0))
		Expect(bowling.NewGame("1-").Score()).To(Equal(1))
		Expect(bowling.NewGame("2-").Score()).To(Equal(2))
		Expect(bowling.NewGame("12").Score()).To(Equal(3))
		Expect(bowling.NewGame("123-").Score()).To(Equal(6))
		Expect(bowling.NewGame("X12").Score()).To(Equal(16))
		Expect(bowling.NewGame("XXXXXXXXXXXX").Score()).To(Equal(300))
		Expect(bowling.NewGame("9-9-9-9-9-9-9-9-9-9-").Score()).To(Equal(90))
		Expect(bowling.NewGame("5/5/5/5/5/5/5/5/5/5/5").Score()).To(Equal(150))
	})

	It("can calculate scores with non-final spares", func() {
		Expect(bowling.NewGame("3/4-").Score()).To(Equal(18))
	})

	It("can parse inputs into framelists", func() {
		Expect(bowling.Throws2Frames("34")).To(Equal([]bowling.Frame{
			bowling.RegularFrame{Throw1: 3, Throw2: 4},
		}))
		Expect(bowling.Throws2Frames("-4")).To(Equal([]bowling.Frame{
			bowling.RegularFrame{Throw1: 0, Throw2: 4},
		}))
		Expect(bowling.Throws2Frames("3/")).To(Equal([]bowling.Frame{
			bowling.SpareFrame{Throw: 3},
		}))

		Expect(bowling.Throws2Frames("225/")).To(Equal([]bowling.Frame{
			bowling.RegularFrame{Throw1: 2, Throw2: 2},
			bowling.SpareFrame{Throw: 5},
		}))

		Expect(bowling.Throws2Frames("X")).To(Equal([]bowling.Frame{
			bowling.StrikeFrame{},
		}))
		Expect(bowling.Throws2Frames("XX")).To(Equal([]bowling.Frame{
			bowling.StrikeFrame{},
			bowling.StrikeFrame{},
		}))
		Expect(bowling.Throws2Frames("X72X")).To(Equal([]bowling.Frame{
			bowling.StrikeFrame{},
			bowling.RegularFrame{Throw1: 7, Throw2: 2},
			bowling.StrikeFrame{},
		}))
	})

	It("Deals with a single number in a frame", func() {
		Expect(bowling.NewFrame("5")).To(Equal(bowling.RegularFrame{
			Throw1: 5,
		}))
	})
})
