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

	Context("From a grid with turned off lights", func() {
		It("should create a grid with lights turned off", func() {
			Expect(grid.LightsOn()).To(Equal(0))
		})

		It("should turn on 1 light at 0,0", func() {
			grid.TurnOn(0, 0, 0, 0)
			Expect(grid.LightsOn()).To(Equal(1))
		})

		It("should turn on 2 lights at 0,0 and 0,1", func() {
			grid.TurnOn(0, 0, 0, 1)
			Expect(grid.LightsOn()).To(Equal(2))
		})

		It("should turn on 3 lights from 0,1 to 0,3", func() {
			grid.TurnOn(0, 1, 0, 3)
			Expect(grid.LightsOn()).To(Equal(3))
		})

		It("should turn on 2 lights at 0,0 and 1,0", func() {
			grid.TurnOn(0, 0, 1, 0)
			Expect(grid.LightsOn()).To(Equal(2))
		})

		It("should turn on 3 lights from 1,0 to 3,0", func() {
			grid.TurnOn(1, 0, 3, 0)
			Expect(grid.LightsOn()).To(Equal(3))
		})

		It("should turn on 6 lights from 2,1 to 3,3", func() {
			grid.TurnOn(2, 1, 3, 3)
			Expect(grid.LightsOn()).To(Equal(6))
		})
	})

	Context("From a grid with lights on", func() {
		var lightsOnBeforeTurningOff int

		BeforeEach(func() {
			grid.TurnOn(0, 0, 999, 999)
			lightsOnBeforeTurningOff = grid.LightsOn()
		})

		It("should turn off 1 light at 0,0", func() {
			grid.TurnOff(0, 0, 0, 0)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 1))
		})

		It("should turn off 2 lights from 0,0 to 1,0", func() {
			grid.TurnOff(0, 0, 1, 0)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 2))
		})

		It("should turn off 3 lights from 1,0 to 3,0", func() {
			grid.TurnOff(1, 0, 3, 0)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 3))
		})

		It("should turn off 2 lights from 0,0 to 0,1", func() {
			grid.TurnOff(0, 0, 0, 1)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 2))
		})

		It("should turn off 3 lights from 0,1 to 0,3", func() {
			grid.TurnOff(0, 1, 0, 3)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 3))
		})

		It("should turn off 6 lights from 2,1 to 3,3", func() {
			grid.TurnOff(2, 1, 3, 3)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 6))
		})

		It("turn off should be idempotent on points", func() {
			grid.TurnOff(0, 0, 0, 0)
			grid.TurnOff(0, 0, 0, 0)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 1))
		})

		It("should still turn off several points consecutively", func() {
			grid.TurnOff(0, 0, 0, 0)
			grid.TurnOff(1, 1, 1, 1)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 2))
		})

		It("turn off should be idempotent on squares", func() {
			c := make(map[Point]int)
			c[Point{}] = 3
			print(c[Point{0, 1}])

			grid.TurnOff(1, 1, 3, 3)
			grid.TurnOff(2, 2, 2, 2)
			Expect(grid.LightsOn()).To(Equal(lightsOnBeforeTurningOff - 9))
		})
	})

})
