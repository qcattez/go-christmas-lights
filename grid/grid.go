package grid

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	Toogle(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type point struct {
	absiss   int
	ordinate int
}

type grid struct {
	lightsOn map[point]bool
}

func New() Grid {
	grid := grid{}
	grid.lightsOn = make(map[point]bool)
	return &grid
}

func (g grid) LightsOn() int {
	return len(g.lightsOn)
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, point := range allPointsBetween(point{absiss1, ordinate1}, point{absiss2, ordinate2}) {
		g.turnOnLight(point)
	}
}

func (g *grid) turnOnLight(point point) {
	g.lightsOn[point] = true
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, point := range allPointsBetween(point{absiss1, ordinate1}, point{absiss2, ordinate2}) {
		g.turnOffLight(point)
	}
}

func (g *grid) turnOffLight(point point) {
	delete(g.lightsOn, point)
}

func (g *grid) Toogle(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, point := range allPointsBetween(point{absiss1, ordinate1}, point{absiss2, ordinate2}) {
		if g.lightIsOn(point) {
			g.turnOffLight(point)
			continue
		}
		g.turnOnLight(point)
	}
}

func (g grid) lightIsOn(point point) bool {
	return g.lightsOn[point]
}

func allPointsBetween(firstPoint point, lastPoint point) []point {
	var rowPoints []point

	for absiss := firstPoint.absiss; absiss <= lastPoint.absiss; absiss++ {
		for ordinate := firstPoint.ordinate; ordinate <= lastPoint.ordinate; ordinate++ {
			rowPoints = append(rowPoints, point{absiss, ordinate})
		}
	}

	return rowPoints
}
