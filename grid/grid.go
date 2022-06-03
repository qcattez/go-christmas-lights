package grid

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type grid struct {
	lightsOn map[point]bool
}

type point struct {
	absiss   int
	ordinate int
}

func (p point) next() point {
	return point{p.absiss + 1, p.ordinate}
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, point := range createRow(point{absiss1, ordinate1}, point{absiss2, ordinate2}) {
		g.lightsOn[point] = true
	}
}

func createRow(point1 point, point2 point) []point {
	var rowPoints []point

	for p := point1; p != point2; p = p.next() {
		rowPoints = append(rowPoints, p)
	}

	return append(rowPoints, point2)
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, point := range createRow(point{absiss1, ordinate1}, point{absiss2, ordinate2}) {
		delete(g.lightsOn, point)
	}
}

func (g grid) LightsOn() int {
	return len(g.lightsOn)
}

func New() Grid {
	grid := grid{}
	grid.lightsOn = make(map[point]bool)
	return &grid
}
