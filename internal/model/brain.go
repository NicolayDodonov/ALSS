package model

import (
	l "artificialLifeGo/internal/logger"
	"strconv"
)

type brain16 struct{}

func (brain16) run(e *Entity, w *World) error {
	//выполняем генетический код
	//не все команды равноценны по сложности, по этому
	//выполняем их со счётчиком frameCount. Это создёт
	//более сложное поведение ботов.
	for frameCount := 0; frameCount < 10; {
		//создаём переменную некретической ошибки
		var errGen error

		//считываем генокод по указателю
		switch e.Array[e.Pointer] {
		case move:
			errGen = e.move(w)
			frameCount += 5

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " move")
		case look:
			//функционал логического перехода
			var dPointer int
			dPointer, errGen = e.look(w)
			e.Pointer += dPointer - 1
			frameCount += 2

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " look")
		case get:
			errGen = e.get(w)
			frameCount += 5

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " get")
		case rotatedLeft:
			e.rotation(left)
			frameCount++

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " tunrs left")
		case rotatedRight:
			e.rotation(right)
			frameCount++

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " tunrs right")
		case recycling:
			errGen = e.recycling(w)
			frameCount += 5

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " recycling")
		case reproduction:
			errGen = e.reproduction()
			frameCount += 12

			l.Ent.Debug("id " + strconv.Itoa(e.ID) + " make new bot")
		default:
			e.jump()
			frameCount++
		}
		//Логгируем некретические ошибки генокода
		if errGen != nil {
			l.Ent.Debug("id:" + strconv.Itoa(e.ID) + " " + errGen.Error())
		}

		//увеличиваем программно-генетический счётчик
		e.Pointer++
		e.loopPointer()

		//добавляем отравление на клетку с ботом
		if err := w.SetCellPoison(e.Coordinates, pLevel1+1); err != nil {
			return err
		}
	}
	return nil
}

type brain64 struct{}

func (brain64) run(e *Entity, w *World) error {
	return nil
}

type brainNero struct{}

func (brainNero) run(e *Entity, w *World) error {
	return nil
}
