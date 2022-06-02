package grid

type Grid interface {
	LightsOn() int
	TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int)
}

type grid struct {
	lights int
}

func (g *grid) TurnOn(absiss1 int, ordinate1 int, absiss2 int, ordinate2 int) {
	g.lights = (absiss2 - absiss1 + 1) * (ordinate2 - ordinate1 + 1)
}

func (g grid) LightsOn() int {
	return g.lights
}

func New() Grid {
	return &grid{}
}
