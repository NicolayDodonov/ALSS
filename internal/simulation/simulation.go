package simulation

import (
	"artificialLifeGo/internal/console"
	l "artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/model"
	"strconv"
)

type Simulation struct {
	printer       console.Console
	endPopulation int
}

func New(console console.Console, endPop int) (s *Simulation) {
	return &Simulation{
		printer:       console,
		endPopulation: endPop,
	}
}

func (s *Simulation) Train(world_X, world_Y, endAge, mutation, baseLevelPoison int) {
	l.Sim.Debug("start train")
	//определяем стартовую популяцию как конечная популяция^2
	startPopulation := s.endPopulation * s.endPopulation

	//создаёсм мир
	w := model.NewWorld(world_X, world_Y, startPopulation, baseLevelPoison)

	//выполняем цикл обучения
	for w.Age < endAge {
		l.Sim.Debug("start new cycle")
		//очистить мир
		w.Age = 0
		w.Clear()

		for {
			//обновить состояние ресурсов
			w.Update()

			//выполнить генокод всех сущностей
			_ = w.Execute()

			//обновляем статистику
			w.UpdateStat()

			//отрисовываем кадр мира в консоле
			s.printer.Print(w)

			l.Sim.Debug("world " + strconv.Itoa(w.ID) + "age " + strconv.Itoa(w.Age) + "is done!\n" +
				"in world live now: " + strconv.Itoa(w.CountEntity))
			//проверить, живо ли больше endPopulation сущностей
			if w.CountEntity <= s.endPopulation {
				break
			}
			w.Age++
		}
		l.Sim.Info("world №" + strconv.Itoa(w.ID) + " is dead!\n" +
			w.GetPrettyStatistic())
		l.Sim.Info(strconv.Itoa(s.endPopulation) + " best bot's DNA:\n" +
			w.GetPrettyEntityInfo(s.endPopulation))
		w.SetGeneration(s.endPopulation, mutation)
		w.ID++

		//todo: Логгирование
	}
}
