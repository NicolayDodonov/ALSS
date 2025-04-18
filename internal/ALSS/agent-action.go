package ALSS

import "fmt"

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

	//calculate energy profit
	//todo: add coefficient
	a.Energy += cell.Height * cell.localMinerals * c.world.Illumination * c.Parameters.baseSunCost / c.world.Pollution
}

func (a *agent) eatMinerals(c *Controller) {
	cell, _ := c.world.getCell(&a.coordinates)

	a.Energy += cell.localMinerals * c.Parameters.baseMineralCost
}

func (a *agent) eatPollution(c *Controller) {
	a.Energy += c.world.Pollution / c.Parameters.basePollutionPart
}

func (a *agent) eatGrass(c *Controller) {
	cell, _ := c.world.getCell(offset(&a.coordinates, a.Angle))

	a.Energy += cell.localGrass * c.Parameters.baseGrassCost
}

func (a *agent) attack(c *Controller) error {
	cell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return fmt.Errorf("cant attack: %w", err)
	}

	if cell.Agent == nil {
		return nil
	}
	profit := cell.Agent.Energy * c.Parameters.baseAttackPart / 100
	cell.Agent.Energy = -1
	cell.Agent.deathHandler(c)

	a.Energy += profit
	return nil
}

func (a *agent) look(c *Controller) {

}

func (a *agent) lookPoison(c *Controller) {

}

func (a *agent) friendOrFoe() {

}

func (a *agent) getEnergy() {

}
