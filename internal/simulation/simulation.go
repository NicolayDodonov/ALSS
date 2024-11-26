package simulation

import (
	"artificialLifeGo/internal/console"
	l "artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/model"
	"strconv"
)

type Simulation struct {
	printer console.Console
}

func New(console console.Console) (s *Simulation) {
	return &Simulation{
		printer: console,
	}
}

// Train производит обучение в заданных условиях ботов для получения лучших по геному ботов.
func (s *Simulation) Train() []string {
	l.Sim.Info("start train")
	defer l.Sim.Info("end train")
	//определяем стартовую популяцию как конечная популяция^2
	startPopulation := endPopulation * endPopulation

	//создаёсм мир
	w := model.NewWorld(worldSizeX, worldSizeY, startPopulation, basePoisonLevel)

	//выполняем цикл обучения
	for w.Age < finalAgeTrain {
		l.Sim.Debug("start new cycle")
		//очистить мир
		w.Age = 0
		w.Clear()

		w.Update(30)
		for {
			//обновить состояние ресурсов
			if w.Age%recurseUpdateRate == 0 {
				w.Update(30)
			}

			//выполнить генокод всех сущностей
			w.Execute()
			w.RemoveDead()

			//обновляем статистику
			w.UpdateStat()

			//отрисовываем кадр мира в консоле
			s.printer.Print(w)

			l.Sim.Debug("world " + strconv.Itoa(w.ID) + "age " + strconv.Itoa(w.Age) + "is done!\n" +
				"in world live now: " + strconv.Itoa(w.CountEntity))
			//проверить, живо ли больше endPopulation сущностей
			if w.CountEntity <= endPopulation {
				break
			}
			w.Age++
		}
		//Вывести информацию о мире
		l.Sim.Info("world is dead! " +
			w.GetStatistic())
		l.Sim.Debug(strconv.Itoa(endPopulation) + " best bot's DNA:\n" +
			w.GetPrettyEntityInfo(endPopulation))
		w.SetGeneration(endPopulation, mutationCount)
		//и обновить ID мира для следующей итерации
		w.ID++
	}

	return w.GetEntityInfo(endPopulation)
}

func (s *Simulation) Run() {
	//todo: set dna in population

	//todo: execute worlds

	//todo: get data from worlds
}

func (s *Simulation) Experiment() {
	//todo: train 3 type population

	//todo: set parameters experiments

	//todo: make big world with all type population

	//todo: execute experiments

	//todo: get data from experiments
	//todo: save data to .cvs file
}
