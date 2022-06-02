package grid

type Grid interface {
	LightsOn() int
	TurnOn(p1a int, p1o int, p2a int, p2o int)
}

type grid struct {
	lights int
}

func (g *grid) TurnOn(p1a int, p1o int, p2a int, p2o int) {
	g.lights = (p2a - p1a + 1) * (p2o - p1o + 1)
}

func (g grid) LightsOn() int {
	return g.lights
}

func New() Grid {
	return &grid{}
}
