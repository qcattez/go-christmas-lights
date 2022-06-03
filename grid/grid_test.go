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

	It("should create a grid with lights turned off", func() {
		Expect(grid.Brightness()).To(Equal(0))
	})

	It("should increase the brightness of 1 light", func() {
		grid.TurnOn(0, 0, 0, 0)
		Expect(grid.Brightness()).To(Equal(1))
	})

	It("should decrease the brightness of 1 light", func() {
		grid.TurnOn(0, 0, 0, 0)
		grid.TurnOn(0, 0, 0, 0)
		grid.TurnOff(0, 0, 0, 0)
		Expect(grid.Brightness()).To(Equal(1))
	})

	It("should decrease the brightness of the correct light", func() {
		grid.TurnOn(1, 1, 1, 1)
		grid.TurnOff(0, 0, 0, 0)
		grid.TurnOff(0, 1, 0, 1)
		grid.TurnOff(1, 0, 1, 0)
		Expect(grid.Brightness()).To(Equal(1))
	})

	It("should increase the brightness of 1 row of 2 lights", func() {
		grid.TurnOn(0, 0, 1, 0)
		Expect(grid.Brightness()).To(Equal(2))
	})

	It("should decrease the brightness of 1 row of 2 lights", func() {
		grid.TurnOn(0, 0, 1, 0)
		grid.TurnOff(1, 0, 2, 0)
		Expect(grid.Brightness()).To(Equal(1))
	})

	It("should increase the brightness of a row of several lights", func() {
		grid.TurnOn(1, 0, 5, 0)
		Expect(grid.Brightness()).To(Equal(5))
	})

	It("should decrease the brightness of a row of several lights", func() {
		grid.TurnOn(1, 0, 5, 0)
		grid.TurnOff(2, 0, 4, 0)
		Expect(grid.Brightness()).To(Equal(2))
	})

	It("should increase the brightness of 1 column of several lights", func() {
		grid.TurnOn(0, 1, 0, 5)
		Expect(grid.Brightness()).To(Equal(5))
	})

	It("should decrease the brightness of 1 column of several lights", func() {
		grid.TurnOn(0, 1, 0, 5)
		grid.TurnOff(0, 2, 0, 4)
		Expect(grid.Brightness()).To(Equal(2))
	})

	It("should increase the brightness twice of 1 light", func() {
		grid.Toogle(0, 0, 0, 0)
		Expect(grid.Brightness()).To(Equal(2))
	})

	It("should increase the brightness twice of several lights", func() {
		grid.Toogle(1, 2, 3, 4)
		Expect(grid.Brightness()).To(Equal(9 * 2))
	})
})
