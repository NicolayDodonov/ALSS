package ALSS

// файл содержит все обработчики действий агента

func (a *agent) move(angle angle, c *Controller) error {
	newCoord := offset(&a.coordinates, angle)

	newCell := c.world.getCell(newCoord)
	if newCell.Agent != nil {
		return nil
	}

	c.world.getCell(&a.coordinates).Agent = nil

	a.coordinates = *newCoord
	c.world.getCell(newCoord).Agent = a
	return nil
}

func (a *agent) turnLeft() {
	a.Angle.minus()
}

func (a *agent) turnRight() {
	a.Angle.minus()
}

func (a *agent) eatSun() {

}

func (a *agent) eatChemo() {

}

func (a *agent) eatPollution() {

}

func (a *agent) eatGrass() {

}

func (a *agent) attack() {

}

func (a *agent) look() {

}

func (a *agent) lookHeightCell() {

}

func (a *agent) friendOrFoe() {

}

func (a *agent) getEnergy() {

}
