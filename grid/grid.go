package grid

type Grid interface {
	Brightness() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	Toogle(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type coordinate struct {
	absiss   int
	ordinate int
}

type grid struct {
	brightnessPerLight map[coordinate]int
}

func New() Grid {
	grid := grid{}
	grid.brightnessPerLight = make(map[coordinate]int)
	return &grid
}

func (g grid) Brightness() int {
	totalBrightness := 0
	for _, brightness := range g.brightnessPerLight {
		totalBrightness += brightness
	}
	return totalBrightness
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, coordinate := range allCoordinatesInTheSquare(coordinate{absiss1, ordinate1}, coordinate{absiss2, ordinate2}) {
		g.increaseBrightnessBy1(coordinate)
	}
}

func (g *grid) increaseBrightnessBy1(coordinate coordinate) {
	g.brightnessPerLight[coordinate]++
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, coordinate := range allCoordinatesInTheSquare(coordinate{absiss1, ordinate1}, coordinate{absiss2, ordinate2}) {
		g.decreaseBrightnessBy1(coordinate)
	}
}

func (g *grid) decreaseBrightnessBy1(coordinate coordinate) {
	if g.brightnessPerLight[coordinate] > 0 {
		g.brightnessPerLight[coordinate]--
	}
}

func (g *grid) Toogle(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	for _, coordinate := range allCoordinatesInTheSquare(coordinate{absiss1, ordinate1}, coordinate{absiss2, ordinate2}) {
		g.increaseBrightnessBy1(coordinate)
		g.increaseBrightnessBy1(coordinate)
	}
}

func allCoordinatesInTheSquare(firstPoint coordinate, lastPoint coordinate) []coordinate {
	var rowPoints []coordinate

	for absiss := firstPoint.absiss; absiss <= lastPoint.absiss; absiss++ {
		for ordinate := firstPoint.ordinate; ordinate <= lastPoint.ordinate; ordinate++ {
			rowPoints = append(rowPoints, coordinate{absiss, ordinate})
		}
	}

	return rowPoints
}
