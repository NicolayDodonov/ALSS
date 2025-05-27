package ALSS

// файл содержит все обработчики действий агента

// move перемещает агента по углу от его координат
func (a *agent) move(angle angle, c *Controller) error {
	//todo: добавить обработку ошибок
	//определяем его координаты взгляда
	newCoord := offset(&a.coordinates, angle)

	//получаем клетку перемещенияя
	newCell := c.world.getCell(newCoord)
	if newCell == nil {
		return nil
	}

	if newCell.Agent != nil {
		return nil
	}

	//удаляем агента из старой клетки
	oldCell := c.world.getCell(&a.coordinates)
	oldCell.Agent = nil

	//и записываем в новую
	newCell.Agent = a
	a.coordinates = *newCoord

	//если оказался под водой - двойная оплата передвижения
	if newCell.Height <= c.world.SeaLevel {
		a.Energy -= c.Parameters.energyCost
	}
	return nil
}

// Поворот взгляда на лево.
func (a *agent) turnLeft() {
	a.Angle.minus()
}

// Поворот взгляда на право.
func (a *agent) turnRight() {
	a.Angle.minus()
}

// Команда фотосинтеза.
func (a *agent) eatSun(c *Controller) {
	//get cell data
	cell := c.world.getCell(&a.coordinates)
	if cell.Height > c.world.SeaLevel {
		a.Energy += ((c.world.Illumination * (cell.Height - c.world.SeaLevel)) /
			(c.world.PollutionFix + 1)) * ((cell.LocalMinerals + 1/maxMineral) + 1)
	}

}

// Команда Хемосинтеза.
func (a *agent) eatMinerals(c *Controller) {
	cell := c.world.getCell(&a.coordinates)

	var dMinerals int
	if cell.LocalMinerals > 10 {
		dMinerals = cell.LocalMinerals / 10
	} else {
		dMinerals = cell.LocalMinerals
	}

	cell.LocalMinerals -= dMinerals
	a.Energy += dMinerals
}

// Команда очистки атмосферы
func (a *agent) eatPollution(c *Controller) {
	dPollution := c.world.PollutionFix //todo: надо будет настроить эту строчку
	c.world.Pollution -= dPollution
	dMinerals := dPollution / 2
	cell := c.world.getCell(&a.coordinates)
	cell.LocalMinerals += dMinerals
	a.Energy += dMinerals
}

// Команда Охоты
func (a *agent) attack(c *Controller) error {
	cell := c.world.getCell(offset(&a.coordinates, a.Angle))
	if cell == nil {
		return nil
	}

	if cell.Agent == nil {
		return nil
	}
	profit := cell.Agent.Energy * c.Parameters.attackProfitPercent / 100
	cell.Agent.Energy = -1

	if err := cell.Agent.deathHandler(c); err != nil {
		return err
	}

	a.Energy += profit
	return nil
}

// Команда взгляда на клетку. Оценивает содержимое клетки и совершает переход указателя.
func (a *agent) look(c *Controller) error {
	lookedCell := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}

	//todo: исправить взгляд
	if lookedCell.Agent != nil {
		//проверяем есть ли там вообще агент
		a.Genome.jumpPointer(1)
		return nil
	}
	if lookedCell.LocalMinerals >= baseMinerals {
		a.Genome.jumpPointer(2)
		return nil
	}
	a.Genome.jumpPointer(3)
	return nil
}

// Команда взгляда высоты. Определяет высоту клетки в переди и совершает переход.
func (a *agent) lookHeightCell(c *Controller) error {
	lookedCell := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}
	myCell := c.world.getCell(&a.coordinates)

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

// Команда свой-чужой.
func (a *agent) friendOrFoe(c *Controller) error {
	lookedCell := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
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
		return nil
	}
}

// Команда передачи энергии.
func (a *agent) getEnergy(c *Controller) error {
	lookedCell := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}
	if lookedCell.Agent != nil {
		if a.Energy > livingSurviveLevel {
			a.Energy -= energyTransferPacket
			lookedCell.Agent.Energy += energyTransferPacket
		}
	}
	return nil
}
