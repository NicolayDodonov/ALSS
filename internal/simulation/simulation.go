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

func (s *Simulation) Run(entAge int) {
	//определяем стартовую популяцию как конечная популяция^2
	startPopulation := s.endPopulation * s.endPopulation

	w := model.NewWorld(30, 30, startPopulation)
	err := w.NewGeneration()
	if err != nil {
		//todo:Логирование ошибки
	}

	for w.Age < entAge {
		//todo: настройка мира
		w.Age = 0
		err = w.Clear()

		//todo: добновить ресурсы мира
		err = w.Update()
		//todo: выполняем код для каждого бота
		err = w.Execute()
		//todo: отрисовываем кадр
		s.printer.Print(w)
		//todo: собираем данные
		stat, err := w.GetStatistic()
		if err != nil {

		}
		//todo: создать новое полколение
		err = w.SetGeneration()
		//todo: увеличить счётчик мира
		w.ID++
		//todo: провести логирование конца цикла

	}
	//todo: провести логирование конца симуляции
}
