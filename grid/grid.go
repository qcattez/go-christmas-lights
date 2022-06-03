package grid

import "reflect"

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type grid struct {
	lightRow  []point
	lightIsOn bool
}

type point struct {
	absiss   int
	ordinate int
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	g.lightRow = createLightRow(absiss1, ordinate1, absiss2, ordinate2)
}

func createLightRow(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) []point {
	point1 := point{absiss1, ordinate1}
	point2 := point{absiss2, ordinate2}

	if point1 == point2 {
		return []point{point1}
	}

	return []point{point1, point2}
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	row := createLightRow(absiss1, ordinate1, absiss2, ordinate2)
	if reflect.DeepEqual(g.lightRow, row) {
		g.lightRow = []point{}
	}
}

func (g grid) LightsOn() int {
	return len(g.lightRow)
}

func New() Grid {
	return &grid{}
}
