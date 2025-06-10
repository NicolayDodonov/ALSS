package ALSS

import (
	"math/rand/v2"
)

// agent это структура описывающего отдельно взятого бота/агента модели.
type agent struct {
	ID     string
	Age    int
	Energy int
	Angle  angle
	coordinates
	Genome *genome
}

// newAgent возвращает одного нового агента без добавления его в какие либо списки.
func newAgent(c *Controller) *agent {
	a := &agent{
		ID:     newID(),
		Age:    0,
		Energy: c.Parameters.baseAgentEnergy,
		Angle:  0,
		coordinates: coordinates{
			X: rand.IntN(c.world.MaxX),
			Y: rand.IntN(c.world.MaxY),
		},
	}

	//выбираем тип заполнения генома
	switch c.Parameters.typeGenome {
	case genomeTypeRAND:
		//случайное заполнение
		a.Genome = newRandomGenome(c.Parameters.sizeGenome, c.Parameters.maxGen)
	case genomeTypeZERO:
		//заполнить нулями
		a.Genome = newZeroGenome(c.Parameters.sizeGenome)
	default:
		//стандартный набор команд
		a.Genome = newBaseGenome()
	}
	return a
}

// run основной метод агента. Для живых ботов проводит интерпритацию генокода,
// уменьшает энергию, увеличивает возраст, а так же проверяет может ли агент
// умереть или поделиться.
func (a *agent) run(c *Controller) error {
	//проверяем жив ли агент
	if a.Energy <= 0 {
		return nil
	}

	//взимаем плату с агента за интерпритацию кода
	a.Energy -= c.Parameters.energyCost
	a.Age++

	//интерпритируем генокод.
	if err := a.interpretationGenome(c); err != nil {
		return err
	}
	// обрабатываем отравление
	a.pollutionHandler(c)
	//смерть
	if err := a.deathHandler(c); err != nil {
		return err
	}
	//и рождение
	if err := a.birthHandler(c); err != nil {
		return err
	}

	return nil
}

// interpretationGenome - мозг бота. Функция отвечает за интерпритацию в реальном времени
// значений генокода агента и вызова соотвествующей функции.
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
		err = a.eatSun(c)
	case 12:
		err = a.eatMinerals(c)
	case 13:
		a.eatPollution(c)
	case 14:
		err = a.attack(c)
	case 15:
		err = a.look(c)
	case 16:
		err = a.friendOrFoe(c)
	case 17:
		err = a.lookHeightCell(c)
	case 18:
		err = a.getEnergy(c)
	default:
		a.Genome.jumpPointer(gen - 1)
	}
	a.Genome.jumpPointer(1)
	return err
}

// pollutionHandler обрабатывает увеличение загрязнения и урон от минерального
// перенасыщения месности.
func (a *agent) pollutionHandler(c *Controller) error {
	//увеличить общее загрянение воздуха
	c.world.addPollution(c.Parameters.madePollution)

	//нанести урон от загрязнения минералами
	cell, err := c.world.getCell(&a.coordinates)
	if err != nil {
		return err
	}

	if cell.LocalMinerals >= 200 {
		a.Energy -= c.Parameters.energyCost
	}
	return nil
}

// birthHandler обработчик размножения.
func (a *agent) birthHandler(c *Controller) error {
	//проверяем возможность рождения по энергии
	if a.Energy < c.Parameters.minEnergyToBirth {
		return nil
	}

	//и по пустому месту рядом
	var freeCoords *coordinates = nil
	for angle := angle(0); angle < 8; angle++ {
		cell, _ := c.world.getCell(offset(&a.coordinates, angle))
		if cell != nil && cell.Agent == nil {
			//выходим по первой пустой клетке
			freeCoords = offset(&a.coordinates, angle)
			break
		}
	}

	//если есть пустая клетка - размножаемся
	if freeCoords != nil {
		// отдаём половину энергии
		a.Energy /= 2

		//создаём нового агента
		newA := &agent{
			ID:          newID(),
			Age:         0,
			Energy:      a.Energy / 2,
			Angle:       a.Angle,
			Genome:      a.Genome,
			coordinates: *freeCoords,
		}
		newA.Genome.mutation(c.Parameters.countMutation)

		//распологаем его в пустой клетке
		cell, _ := c.world.getCell(freeCoords)
		cell.Agent = newA

		//добавляем в список агентов после себя
		if err := c.agents.addAfter(a, newA); err != nil {
			return err
		}
	}
	return nil
}

// deathHandler обработчик смерти. Проверяет возможность смерти.
// Для успешного вызова нужно или 0 энергии, или Parameters maxAgentEnergy, или Parameters maxAgentAge.
// Удаляет агента из клетки и списка активныъ агентов при правдивости условия, иначе nil.
// Возвращает ошибки удаляения из списка или при аномалии расположения агента.
func (a *agent) deathHandler(c *Controller) error {
	if a.Energy <= 0 || a.Energy > c.Parameters.maxAgentEnergy || a.Age >= c.Parameters.maxAgentAge {
		//удаляем из списка активных агентов
		if err := c.agents.del(a); err != nil {
			return err
		}
		//удаляем из клетки
		cell, err := c.world.getCell(&a.coordinates)
		if err != nil {
			return err
		}
		cell.Agent = nil

		//расчитываем добавление минералов
		var MineralUnit = 0
		if a.Energy < 10 {
			MineralUnit = 10
		} else {
			//если агент умер от переедания
			MineralUnit = a.Energy / 10
		}

		//добавляем минералы в площади
		for i := angle(0); i < 8; i++ {
			c.world.addMinerals(offset(&a.coordinates, i), MineralUnit)
		}
		c.world.addMinerals(&a.coordinates, MineralUnit*2)

		//стамив "мертвые" значения энергии.
		a.Energy = -1
		c.Stats.Deaths++
	}
	return nil
}

// checkSelfPosition проверяет правильность распложения агента в своих координатах.
// Возвращает true при соотвествии координат и клетки.
// Возвращает false при несоответсвии координат и клетки или присутсвии другого агента.
func (a *agent) checkSelfPosition(c *Controller) bool {
	cell, err := c.world.getCell(&a.coordinates)
	if err != nil {
		return false
	}
	if cell.Agent == a {
		return true
	}
	return false
}
