package ALSS

import (
	"container/list"
	"math/rand/v2"
)

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
		Energy: c.Parameters.baseAgentEnergy,
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

func (a *agent) run(c *Controller, me *list.Element) error {
	if a.Energy <= 0 {
		return nil
	}

	a.Energy -= c.Parameters.energyCost
	a.Age++

	if err := a.interpretationGenome(c, me); err != nil {
		return err
	}

	a.pollutionHandler(c)

	a.deathHandler(c, me)

	a.birthHandler(c, me)

	return nil
}

func (a *agent) interpretationGenome(c *Controller, me *list.Element) error {
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
		err = a.attack(c, me)
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

func (a *agent) pollutionHandler(c *Controller) {
	// увеличить общее загрянение воздуха
	c.world.Pollution += c.Parameters.madePollution

	//нанести урон от загрязнения минералами
	cell, _ := c.world.getCell(&a.coordinates)
	if cell.localMinerals >= 200 {
		a.Energy -= c.Parameters.energyCost
	}
}

func (a *agent) birthHandler(c *Controller, me *list.Element) {
	//проверяем возможность рождения
	if a.Energy < c.Parameters.minEnergyToBirth {
		return
	}

	var freeCoords *coordinates = nil
	for angle := angle(0); angle < 8; angle++ {
		cell, _ := c.world.getCell(offset(&a.coordinates, angle))
		if cell != nil && cell.Agent == nil {
			freeCoords = offset(&a.coordinates, angle)
			break
		}
	}

	if freeCoords != nil {
		a.Energy /= 2

		//make new agent
		newA := agent{
			Age:         0,
			Energy:      a.Energy / 2,
			Angle:       a.Angle,
			Genome:      a.Genome,
			coordinates: *freeCoords,
		}
		makeID(newA)
		newA.Genome.mutation(c.Parameters.countMutation)

		_ = c.agents.InsertBefore(newA, me)
	}
}

func (a *agent) deathHandler(c *Controller, me *list.Element) {
	if a.Energy <= 0 || a.Energy > c.Parameters.maxAgentEnergy || a.Age >= c.Parameters.maxAgentAge {
		_ = c.agents.Remove(me)

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

func (a *agent) checkSelfPosition(c *Controller) bool {
	cell, err := c.world.getCell(&a.coordinates)
	if err != nil {
		return true
	}
	if cell.Agent == a {
		return true
	} else {
		return false
	}
}
