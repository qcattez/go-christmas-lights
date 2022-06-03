package grid

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {
	var grid Grid

	BeforeEach(func() {
		grid = New()
	})

	It("should create a grid with lightsOn turned off", func() {
		Expect(grid.LightsOn()).To(Equal(0))
	})

	It("should turn on 1 light", func() {
		grid.TurnOn(0, 0, 0, 0)
		Expect(grid.LightsOn()).To(Equal(1))
	})

	It("should turn off 1 light", func() {
		grid.TurnOn(0, 0, 0, 0)
		grid.TurnOff(0, 0, 0, 0)
		Expect(grid.LightsOn()).To(Equal(0))
	})

	It("should turn off the correct light", func() {
		grid.TurnOn(1, 1, 1, 1)
		grid.TurnOff(0, 0, 0, 0)
		grid.TurnOff(0, 1, 0, 1)
		grid.TurnOff(1, 0, 1, 0)
		Expect(grid.LightsOn()).To(Equal(1))
	})

	It("should turn on 1 row of 2 lights", func() {
		grid.TurnOn(0, 0, 1, 0)
		Expect(grid.LightsOn()).To(Equal(2))
	})

	It("should turn off 1 row of 2 lights", func() {
		grid.TurnOn(0, 0, 1, 0)
		grid.TurnOff(1, 0, 2, 0)
		Expect(grid.LightsOn()).To(Equal(1))
	})

	It("should turn on a row of several lights", func() {
		grid.TurnOn(1, 0, 5, 0)
		Expect(grid.LightsOn()).To(Equal(5))
	})

	It("should turn off a row of several lights", func() {
		grid.TurnOn(1, 0, 5, 0)
		grid.TurnOff(2, 0, 4, 0)
		Expect(grid.LightsOn()).To(Equal(2))
	})

	It("should turn on 1 column of several lights", func() {
		grid.TurnOn(0, 1, 0, 5)
		Expect(grid.LightsOn()).To(Equal(5))
	})

	It("should turn off 1 column of several lights", func() {
		grid.TurnOn(0, 1, 0, 5)
		grid.TurnOff(0, 2, 0, 4)
		Expect(grid.LightsOn()).To(Equal(2))
	})

	It("should toogle 1 light on if it was off", func() {
		grid.Toogle(0, 0, 0, 0)
		Expect(grid.LightsOn()).To(Equal(1))
	})

	It("should toogle 1 light off if it was on", func() {
		grid.Toogle(0, 0, 0, 0)
		grid.Toogle(0, 0, 0, 0)
		Expect(grid.LightsOn()).To(Equal(0))
	})

	It("should toogle several lights", func() {
		grid.Toogle(1, 2, 3, 4)
		Expect(grid.LightsOn()).To(Equal(9))
	})
})
