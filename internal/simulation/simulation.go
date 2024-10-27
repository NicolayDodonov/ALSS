package simulation

import (
	"artificialLifeGo/internal/console"
	"artificialLifeGo/internal/model"
)

type Simulation struct {
	printer       *console.Console
	endPopulation int
}

func New(console *console.Console, endPop int) (s *Simulation) {
	return &Simulation{
		printer:       console,
		endPopulation: endPop,
	}
}

func (s *Simulation) Train(endAge, mutation int) {
	//определяем стартовую популяцию как конечная популяция^2
	startPopulation := s.endPopulation * s.endPopulation

	//создаёсм мир
	w := model.NewWorld(30, 30, startPopulation)

	//выполняем цикл обучения
	for w.Age < endAge {
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

			//проверить, живо ли больше endPopulation сущностей
			if w.CountEntity <= s.endPopulation {
				break
			}
			w.Age++
		}
		w.SetGeneration(s.endPopulation, mutation)
		w.ID++

		//todo: Логгирование
	}
}
