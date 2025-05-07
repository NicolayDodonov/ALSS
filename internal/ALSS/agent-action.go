package ALSS

import (
	"container/list"
	"fmt"
)

// файл содержит все обработчики действий агента

func (a *agent) move(angle angle, c *Controller) error {
	//take new coordinates to move
	newCoord := offset(&a.coordinates, angle)

	//get cell from new coordinates
	newCell, err := c.world.getCell(newCoord)
	if err != nil {
		return err
	}
	if newCell.Agent != nil {
		return nil
	}

	//get out from old cell
	oldCell, _ := c.world.getCell(&a.coordinates)
	oldCell.Agent = nil

	//and go to new cell!
	newCell.Agent = a
	a.coordinates = *newCoord

	//if cell under sea level - energyCost!
	if newCell.Height <= c.world.SeaLevel {
		a.Energy -= c.Parameters.energyCost
	}
	return nil
}

func (a *agent) turnLeft() {
	a.Angle.minus()
}

func (a *agent) turnRight() {
	a.Angle.minus()
}

func (a *agent) eatSun(c *Controller) {
	//get cell data
	cell, _ := c.world.getCell(&a.coordinates)
	if cell.Height > c.world.SeaLevel {
		a.Energy += ((c.world.Illumination * (cell.Height - c.world.SeaLevel)) /
			c.world.PollutionFix) * ((cell.Height / maxHeight) + 1)
	}

}

func (a *agent) eatMinerals(c *Controller) {
	cell, _ := c.world.getCell(&a.coordinates)

	var dMinerals int
	if cell.localMinerals > 10 {
		dMinerals = cell.localMinerals / 10
	} else {
		dMinerals = cell.localMinerals
	}

	cell.localMinerals -= dMinerals
	a.Energy += dMinerals
}

func (a *agent) eatPollution(c *Controller) {
	dPollution := c.world.PollutionFix //todo: надо будет настроить эту строчку
	c.world.Pollution -= dPollution
	dMinerals := dPollution / 2
	cell, _ := c.world.getCell(&a.coordinates)
	cell.localMinerals += dMinerals
	a.Energy += dMinerals
}

func (a *agent) attack(c *Controller, me *list.Element) error {
	cell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return fmt.Errorf("cant attack: %w", err)
	}

	if cell.Agent == nil {
		return nil
	}
	profit := cell.Agent.Energy * c.Parameters.attackProfitPercent / 100
	cell.Agent.Energy = -1
	cell.Agent.deathHandler(c, me)

	a.Energy += profit
	return nil
}

// look check what is located in the looked cell: other agent, grass, mineral or nothing
func (a *agent) look(c *Controller) error {
	lookedCell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return err
	}
	//todo: исправить взгляд
	if lookedCell.Agent != nil {
		//проверяем есть ли там вообще агент
		a.Genome.jumpPointer(1)
		return nil
	}
	if lookedCell.localMinerals >= baseMinerals {
		a.Genome.jumpPointer(2)
		return nil
	}
	a.Genome.jumpPointer(3)
	return nil
}

func (a *agent) lookHeightCell(c *Controller) error {
	lookedCell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	myCell, _ := c.world.getCell(&a.coordinates)
	if err != nil {
		return err
	}
	if lookedCell.Height > myCell.Height {
		a.Genome.jumpPointer(1)
		return nil
	}
	if lookedCell.Height < myCell.Height {
		a.Genome.jumpPointer(2)
		return nil
	}
	if lookedCell.Height == myCell.Height {
		a.Genome.jumpPointer(3)
		return nil
	}
	return nil
}

func (a *agent) friendOrFoe(c *Controller) error {
	lookedCell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return err
	}
	if lookedCell.Agent != nil {
		if equals(lookedCell.Agent.Genome, a.Genome) {
			// friend
			a.Genome.jumpPointer(1)
			return nil
		} else {
			// foe
			a.Genome.jumpPointer(2)
			return nil
		}
	} else {
		// looked cell is empty
		a.Genome.jumpPointer(3)
		return err
	}
}

func (a *agent) getEnergy(c *Controller) error {
	lookedCell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return err
	}
	if lookedCell.Agent != nil {
		if a.Energy > livingSurviveLevel {
			a.Energy -= energyTransferPacket
			lookedCell.Agent.Energy += energyTransferPacket
		}
	}
	return nil
}
