package grid

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
	TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type grid struct {
	lights int
}

func (g *grid) TurnOff(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	g.lights = g.lights - squareArea(absiss1, ordinate1, absiss2, ordinate2)
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	g.lights = squareArea(absiss1, ordinate1, absiss2, ordinate2)
}

func squareArea(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) int {
	return (absiss2 - absiss1 + 1) * (ordinate2 - ordinate1 + 1)
}

func (g grid) LightsOn() int {
	return g.lights
}

func New() Grid {
	return &grid{}
}
