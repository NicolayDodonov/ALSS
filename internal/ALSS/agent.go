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
	a := &agent{
		Age:    0,
		Energy: c.Parameters.baseEnergy,
		Angle:  0,
		coordinates: coordinates{
			X: rand.IntN(c.world.MaxX),
			Y: rand.IntN(c.world.MaxY),
		},
	}
	makeID(a)

	switch c.Parameters.typeGenome {
	case genomeTypeRAND:
		a.Genome = newRandomGenome(c.Parameters.sizeGenome)
	case genomeTypeZERO:
		a.Genome = newZeroGenome(c.Parameters.sizeGenome)
	default:
		a.Genome = newBaseGenome()
	}
	return a
}

func (a *agent) run(c *Controller) error {
	if a.Energy <= 0 {
		return nil
	}

	a.Energy -= c.Parameters.energyCost
	a.Age++

	if err := a.interpretationGenome(c); err != nil {
		return err
	}

	a.pollution(c)

	a.deathHandler(c)

	a.birthHandler(c)

	return nil
}

func (a *agent) interpretationGenome(c *Controller) error {
	var err error = nil
	gen := a.Genome.getGen()
	switch gen {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		err = a.move(angle(gen), c)
	case 8:
		err = a.move(a.Angle, c)
	case 9:
		a.turnLeft()
	case 10:
		a.turnRight()
	case 11:
		a.eatSun(c)
	case 12:
		a.eatMinerals(c)
	case 13:
		a.eatGrass(c)
	case 14:
		a.eatPollution(c)
	case 15:
		err = a.attack(c)
	case 16:
		err = a.look(c)
	case 17:
		err = a.lookHeightCell(c)
	case 18:
		err = a.friendOrFoe(c)
	case 19:
		err = a.getEnergy(c)
	}
	a.Genome.jumpPointer(1)
	return err
}

func (a *agent) pollution(c *Controller) {
	c.world.Pollution += c.Parameters.pollutionCost
}

func (a *agent) birthHandler(c *Controller) {
	//make new agent
	newA := agent{
		Age:         0,
		Energy:      a.Energy / 2,
		Angle:       a.Angle,
		coordinates: a.coordinates,
		Genome:      a.Genome,
	}
	makeID(newA)
	newA.Genome.mutation(c.Parameters.countMutation)

	//todo: add newAgent before themself in c.agents

}

func (a *agent) deathHandler(c *Controller) {
	if a.Energy <= 0 || a.Age >= c.Parameters.maxAge {
		//todo: remove agent from c.agents

		//add cell minerals
		MineralUnit := a.Energy / 10
		if a.Energy <= 0 {
			MineralUnit = 1
		}
		for i := angle(0); i < 8; i++ {
			_ = c.world.addLocalMinerals(offset(&a.coordinates, i), MineralUnit)
		}
		_ = c.world.addLocalMinerals(&a.coordinates, MineralUnit*2)

		a.Energy = 0
		a.Age = -1
	}
}
