package ALSS

type agent struct {
	ID     int
	Age    int
	Energy int
	Angle  angle
	coordinates
	Genome *genome
}

func newAgent(id int, age int, energy int, angle angle) *agent {
	return &agent{}
}

func (a *agent) run(c *Controller) error {
	if a.Energy <= 0 {
		return nil
	}

	if err := a.interpretationGenome(c); err != nil {
		return err
	}

	if err := a.pollution(c); err != nil {
		return err
	}

	return nil
}

func (a *agent) interpretationGenome(c *Controller) error {
	switch a.Genome.getGen() {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		a.move(a.Genome.Pointer, c)
	case 8:
		a.move(int(a.Angle), c)
	case 9:
		a.turnLeft()
	case 10:
		a.turnRight()
	case 11:
		a.eatSun()
	case 12:
		a.eatChemo()
	case 13:
		a.eatGrass()
	case 14:
		a.eatPollution()
	case 15:
		a.attack()
	case 16:
		a.look()
	case 17:
		a.lookHeightCell()
	case 18:
		a.friendOrFoe()
	case 19:
		a.getEnergy()
	}
	a.Energy--
	a.Genome.jumpPointer(1)
	return nil
}

func (a *agent) pollution(c *Controller) error {

	return nil
}

func (a *agent) death(c *Controller) {
	//todo: remove agent from c.agents
	//todo: add cell minerals
	//todo: add cell-neighbour minerals/2
}

func (a *agent) birth(c *Controller) {
	//todo: make new agent
	//todo: add agent before themself in c.agents

}
