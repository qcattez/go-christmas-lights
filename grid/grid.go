package grid

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type grid struct {
	lights                   int
	turnOffAlreadyCalledOnce bool
	lastSquareTurnedOff      square
	bulbs                    map[Point]bool
}

type square struct {
	bottomRightPoint Point
	topLeftPoint     Point
}

type Point struct {
	absiss   int
	ordinate int
}

func (s square) include(otherSquare square) bool {
	return otherSquare.bottomRightPoint.absiss >= s.bottomRightPoint.absiss && otherSquare.bottomRightPoint.ordinate >= s.bottomRightPoint.ordinate && s.topLeftPoint.absiss >= otherSquare.topLeftPoint.absiss && s.topLeftPoint.absiss >= otherSquare.topLeftPoint.ordinate
}

func (s square) area() int {
	return len(s.points())
}

func (s square) points() []Point {
	minimumAbsiss := s.bottomRightPoint.absiss
	minimumOrdinate := s.bottomRightPoint.ordinate
	maximumAbsiss := s.topLeftPoint.absiss
	maximumOrdinate := s.topLeftPoint.ordinate

	var points []Point
	for j := minimumOrdinate; j <= maximumOrdinate; j++ {
		for i := minimumAbsiss; i <= maximumAbsiss; i++ {
			points = append(points, Point{i, j})
		}
	}

	return points
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	currentSquare := square{Point{absiss1, ordinate1}, Point{absiss2, ordinate2}}
	points := currentSquare.points()

	for i := 0; i < len(points); i++ {
		g.bulbs[points[i]] = false
	}
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	currentSquare := square{Point{absiss1, ordinate1}, Point{absiss2, ordinate2}}
	points := currentSquare.points()

	for i := 0; i < len(points); i++ {
		g.bulbs[points[i]] = true
	}
}

func (g grid) LightsOn() int {
	bulbsOn := 0
	for _, bulbIsOn := range g.bulbs {
		if bulbIsOn {
			bulbsOn++
		}
	}

	return bulbsOn
}

func New() Grid {
	grid := grid{}
	grid.bulbs = make(map[Point]bool)
	return &grid
}
