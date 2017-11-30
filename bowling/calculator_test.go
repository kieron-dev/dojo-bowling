package bowling_test

import (
	"github.com/kieron-pivotal/bowling/bowling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	It("can calculate scores if you never frame more than 9", func() {
		Expect(bowling.Score("0")).To(Equal(0))
		Expect(bowling.Score("1")).To(Equal(1))
		Expect(bowling.Score("2")).To(Equal(2))
		Expect(bowling.Score("12")).To(Equal(3))
		Expect(bowling.Score("123")).To(Equal(6))
	})

})
