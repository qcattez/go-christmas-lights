package grid

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {
	It("should create a grid with lights turned off", func() {
		grid := New()
		Expect(grid.LightsOn()).To(Equal(0))
	})

	It("should turn on 1 light at 0,0", func() {
		grid := New()
		grid.TurnOn(0, 0, 0, 0)
		Expect(grid.LightsOn()).To(Equal(1))
	})

	It("should turn on 2 lights at 0,0 and 0,1", func() {
		grid := New()
		grid.TurnOn(0, 0, 0, 1)
		Expect(grid.LightsOn()).To(Equal(2))
	})

	It("should turn on 3 lights from 0,1 to 0,3", func() {
		grid := New()
		grid.TurnOn(0, 1, 0, 3)
		Expect(grid.LightsOn()).To(Equal(3))
	})

	It("should turn on 2 lights at 0,0 and 1,0", func() {
		grid := New()
		grid.TurnOn(0, 0, 1, 0)
		Expect(grid.LightsOn()).To(Equal(2))
	})

	It("should turn on 3 lights from 1,0 to 3,0", func() {
		grid := New()
		grid.TurnOn(1, 0, 3, 0)
		Expect(grid.LightsOn()).To(Equal(3))
	})

	It("should turn on 6 lights from 2,1 to 3,3", func() {
		grid := New()
		grid.TurnOn(2, 1, 3, 3)
		Expect(grid.LightsOn()).To(Equal(6))
	})
})
