package ALSS

// файл содержит все обработчики действий агента

// move перемещает агента по углу от его координат
func (a *agent) move(angle angle, c *Controller) error {
	//определяем его координаты взгляда
	newCord := offset(&a.coordinates, angle)

	//получаем клетку перемещенияя
	newCell, err := c.world.getCell(newCord)
	if err != nil {
		return nil
	}
	//если в клетке есть кто то - выходим без ошибки
	if newCell.Agent != nil {
		return nil
	}

	//удаляем агента из старой клетки
	oldCell, err := c.world.getCell(&a.coordinates)
	if err != nil {
		//если есть ошибка - нужна синхронизация!
		return err
	}
	oldCell.Agent = nil

	//и записываем в новую
	newCell.Agent = a
	a.coordinates = *newCord

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
func (a *agent) eatSun(c *Controller) error {
	// Получаем клетку агента
	cell, err := c.world.getCell(&a.coordinates)
	if err != nil {
		// если получаем ошибку - нужна синхронизация
		return err
	}
	if cell.Height > c.world.SeaLevel {
		a.Energy += (c.world.Illumination*cell.Height)/10 + (cell.Height / 10) - c.world.PollutionFix
	}
	a.Ration = rationSun
	return nil
}

// Команда Минералосинтеза.
func (a *agent) eatMinerals(c *Controller) {
	// Получаем клетку взгляда агента
	cell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
		return
	}

	// расчитываем сколько съедим из клетки минералов
	if cell.Height > 10 {
		dMinerals := cell.LocalMinerals / 10

		c.world.addMinerals(offset(&a.coordinates, a.Angle), -dMinerals)

		a.Energy += dMinerals
	}
	a.Ration = rationMine
	return
}

// Команда очистки атмосферы - Хемосинтеза
func (a *agent) eatPollution(c *Controller) {
	// расчитываем уменьшение атмосферного яда
	dPollution := c.world.Pollution / (pollutionFixCoefficient * 10)

	c.world.addMinerals(&a.coordinates, dPollution)

	//уменьшаем загрязнение и добавляем энергии агенту
	c.world.addPollution(-dPollution)
	a.Energy += dPollution
	a.Ration = rationHemo
}

// Команда Охоты
func (a *agent) attack(c *Controller) error {
	// Тут не нужна обработка ошибки, так как агент может обращаться в пустую клетку за
	// краем мира. Просто ничего нему не даём и не наказываем за это
	cell, err := c.world.getCell(offset(&a.coordinates, a.Angle))
	if err != nil {
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
	a.Ration = rationHunt
	return nil
}

// Команда взгляда на клетку. Оценивает содержимое клетки и совершает переход указателя.
func (a *agent) look(c *Controller) error {
	// если смотрит за край мира, то клетка будет пустой, а это тоже результат!
	lookedCell, _ := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}

	//todo: исправить взгляд
	if lookedCell.Agent != nil {
		if equals(a.Genome, lookedCell.Agent.Genome) {
			a.Genome.jumpPointer(1)
			return nil
		} else {
			a.Genome.jumpPointer(2)
			return nil
		}
	}
	if lookedCell.LocalMinerals >= baseMinerals {
		a.Genome.jumpPointer(3)
		return nil
	}

	a.Genome.jumpPointer(3)
	return nil
}

// Команда взгляда высоты. Определяет высоту клетки в переди и совершает переход.
func (a *agent) lookHeightCell(c *Controller) error {
	//если смотрит куда то за край мира, то это тоже результат!
	//клетка взгляда будет пустой, что можно проверить
	lookedCell, _ := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}
	//но если своя клетка необнаружена - то это вопрос задуматься...
	myCell, err := c.world.getCell(&a.coordinates)
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

// Команда свой-чужой.
func (a *agent) friendOrFoe(c *Controller) error {
	lookedCell, _ := c.world.getCell(offset(&a.coordinates, a.Angle))
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
	//Если агент пытается отдать энергию краю мира - ничего не случиться!
	lookedCell, _ := c.world.getCell(offset(&a.coordinates, a.Angle))
	if lookedCell == nil {
		return nil
	}
	//Но если там есть агент - то будет передача энергии независимо от того, друг он
	//или враг
	if lookedCell.Agent != nil {
		if a.Energy > livingSurviveLevel {
			a.Energy -= energyTransferPacket
			lookedCell.Agent.Energy += energyTransferPacket
		}
	}
	return nil
}
