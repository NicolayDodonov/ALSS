package ALSS

import "math/rand/v2"

type agent struct {
	ID     string
	Age    int
	Energy int
	Angle  angle
	coordinates
	Genome *genome
}

func newAgent(c *Controller) *agent {
	return &agent{
		makeID(typeOfAgent),
		0,
		c.Parameters.baseHP,
		0,
		coordinates{
			X: rand.IntN(c.world.MaxX),
			Y: rand.IntN(c.world.MaxY),
		},
		newGenome(c.Parameters.typeGenome, c.Parameters.sizeGenome),
	}
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

	a.deathHandler(c)
	return nil
}

func (a *agent) interpretationGenome(c *Controller) error {
	switch a.Genome.getGen() {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		a.move(angle(a.Genome.Pointer), c)
	case 8:
		a.move(a.Angle, c)
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

func (a *agent) deathHandler(c *Controller) {
	if a.Energy <= 0 || a.Age >= c.Parameters.maxAge {
		//todo: remove agent from c.agents

		//add cell minerals
		pollutionUnit := a.Energy / 10
		if a.Energy <= 0 {
			pollutionUnit = 1
		}
		for i := angle(0); i < 8; i++ {
			_ = c.world.addLocalMinerals(offset(&a.coordinates, i), pollutionUnit)
		}
		_ = c.world.addLocalMinerals(&a.coordinates, pollutionUnit*2)

		a.Energy = 0
		a.Age = -1
	}
}

func (a *agent) birth(c *Controller) {
	//make new agent
	newA := agent{
		makeID(typeOfAgent),
		0,
		a.Energy / 2,
		a.Angle,
		a.coordinates,
		a.Genome,
	}
	newA.Genome.mutation(c.Parameters.countMutation)

	//todo: add newAgent before themself in c.agents

}
