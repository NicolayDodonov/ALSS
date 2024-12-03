package simulation

import (
	"artificialLifeGo/internal/alModel"
	"artificialLifeGo/internal/console"
	l "artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/storage"
	"strconv"
)

// Simulation структура описывающая симуляцию
type Simulation struct {
	printer console.Console
	storage.Storage
}

func New(console console.Console, storage storage.Storage) (s *Simulation) {
	return &Simulation{
		console,
		storage,
	}
}

// Train производит обучение в заданных условиях ботов для получения лучших по геному ботов.
func (s *Simulation) Train() []string {
	l.Sim.Info("start train")
	defer l.Sim.Info("end train")
	//определяем стартовую популяцию как конечная популяция^2
	startPopulation := StartPopulation

	//создаёсм мир
	w := alModel.NewWorld(WorldSizeX, WorldSizeY, startPopulation, BasePoisonLevel)

	//выполняем цикл обучения
	for w.Age < FinalAgeTrain {
		l.Sim.Info("Start world№" + strconv.Itoa(w.ID))
		//очистить мир
		w.Age = 0

		w.Update(true)
		for {
			//обновить состояние ресурсов
			w.Update(w.Age%RecurseUpdateRate == 0)

			//выполнить генокод всех сущностей
			w.Execute()
			w.RemoveDead()

			//обновляем статистику
			w.StatisticUpdate()

			//отрисовываем кадр мира в консоле
			s.printer.Print(w)

			l.Sim.Debug("world " + strconv.Itoa(w.ID) + "age " + strconv.Itoa(w.Age) + "is done!\n" +
				"in world live now: " + strconv.Itoa(w.CountEntity))
			//проверить, живо ли больше endPopulation сущностей
			if w.CountEntity <= EndPopulation {
				break
			}
			w.Age++
		}
		//Сохранить
		err := s.WorldAgeSave(w.Age)
		if err != nil {
			l.Sim.Error(err.Error())
		}
		l.Sim.Info("world is dead! " +
			w.GetStatistic())
		l.Sim.Debug(strconv.Itoa(EndPopulation) + " best bot's DNA:\n" +
			w.GetPrettyEntityInfo(EndPopulation))

		w.Clear()
		w.SetGeneration(EndPopulation, MutationCount)
		//и обновить ID мира для следующей итерации
		w.ID++
	}
	//сохраняем данные итогов обучения
	err := s.TrainGenSave(w.GetEntityInfo(EndPopulation))
	if err != nil {
		l.Sim.Error(err.Error())
		return nil
	}
	return w.GetEntityInfo(EndPopulation)
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
