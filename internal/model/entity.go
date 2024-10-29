package model

import (
	l "artificialLifeGo/internal/logger"
	"fmt"
	"strconv"
)

// NewEntity возвращает живую сущность(Entity) с координатами x, y.
func NewEntity(ID, x, y, longDNA int) *Entity {
	return &Entity{
		ID,
		0,
		100,
		true,
		0,
		Coordinates{
			x,
			y,
		},
		*NewDNA(longDNA),
	}
}

// Run отвечает за исполнение генетического кода в DNA.Array.
// Возвращает nil или ошибку.
func (e *Entity) Run(w *World) (err error) {
	l.App.Debug("id " + strconv.Itoa(e.ID) + " is run his genocode")
	//если бот мёрт, вылетаем с ошибкой
	if !e.Live {
		l.App.Error("id " + strconv.Itoa(e.ID) + " is dead!")
		return nil
	}

	//уменьшаем энергию бота перед выполнение генокода
	// "Деньги в перёд"
	e.Energy--
	e.Age++

	//выполняем генетический код
	//не все команды равноценны по сложности, по этому
	//выполняем их со счётчиком frameCount. Это создёт
	//более сложное поведение ботов.
	for frameCount := 0; frameCount < 10; {
		switch e.Array[e.Pointer] {
		case move:
			err = e.move(w)
			frameCount += 5

			l.App.Debug("id " + strconv.Itoa(e.ID) + " move")
		case look:
			//функционал логического перехода
			var dPointer int
			dPointer, err = e.look(w)
			e.Pointer += dPointer - 1
			frameCount += 2

			l.App.Debug("id " + strconv.Itoa(e.ID) + " look")
		case get:
			err = e.get(w)
			frameCount += 5

			l.App.Debug("id " + strconv.Itoa(e.ID) + " get")
		case rotatedLeft:
			e.rotation(left)
			frameCount++

			l.App.Debug("id " + strconv.Itoa(e.ID) + " tunrs left")
		case rotatedRight:
			e.rotation(right)
			frameCount++

			l.App.Debug("id " + strconv.Itoa(e.ID) + " tunrs right")
		case recycling:
			err = e.recycling(w)
			frameCount += 5

			l.App.Debug("id " + strconv.Itoa(e.ID) + " recycling")
		case reproduction:
			err = e.reproduction()
			frameCount += 12

			l.App.Debug("id " + strconv.Itoa(e.ID) + " make new bot")
		default:
			e.jump()
			frameCount++
		}
		//увеличиваем программно-генетический счётчик
		e.Pointer++
		e.loopPointer()

		//если получили ошибку - вылетаем с ошибкой
		if err != nil {
			l.App.Error(err.Error())
		}
	}

	//проверяем, умер ли бот
	if e.Energy <= 0 {
		e.Live = false
		l.App.Info("id " + strconv.Itoa(e.ID) + " die!")
		return fmt.Errorf("I  die")
	}
	return nil
}

// move отвечает за передвижение сущности(Entity) из одной клетки(Cell) мира(World) в другую.
// Возвращает nil или ошибку.
func (e *Entity) move(w *World) error {
	//получаем координаты, куда хотим переместиться
	newCord := viewCell(e.turn)
	//смотрим что там
	cell, err := w.GetCellData(
		Sum(
			newCord,
			e.Coordinates))
	if err != nil {
		return err
	}
	//мы не двигаемся в клетку с другим ботом
	if cell.Entity != nil {
		return fmt.Errorf("is another entity")
	}
	//мы не двигаемся в клетку со стеной
	if cell.Types == WallCell {
		return fmt.Errorf("is wall")
	}
	//двигаемся в клетку
	if err = w.MoveEntity(e.Coordinates, newCord, e); err != nil {
		return err
	}

	return nil
}

// look отвечает за получение данных из другой клетки(Cell). Возвращает номер
// сдвига Entity.DNA.Pointer или ошибку.
func (e *Entity) look(w *World) (int, error) {
	//константы ответов на что мы смотрим
	const (
		isError = iota
		isEmpty
		isFood
		isWall
		isEntity
	)

	//получаем координаты, куда хотим посмотреть
	newCord := viewCell(e.turn)
	//смотрим что там
	cell, err := w.GetCellData(
		Sum(
			newCord,
			e.Coordinates))
	if err != nil {
		return isError, err
	}

	//Определяем тип возврата
	switch cell.Types {
	case EmptyCell:
		if cell.Entity != nil {
			return isEntity, nil
		} else {
			return isEmpty, nil
		}
	case FoodCell:
		return isFood, nil
	case WallCell:
		return isWall, nil
	default:
		return isError, fmt.Errorf("cell type is %v, I dont't know this type", cell.Types)
	}
}

// get отвечает за взаимодействие сущности(Entity) с окружением
// таким как: взять, съесть и тп. Возвращает nil или ошибку.
func (e *Entity) get(w *World) error {
	//получаем координаты для взятия
	newCord := viewCell(e.turn)
	//смотрим что там
	cell, err := w.GetCellData(
		Sum(
			newCord,
			e.Coordinates))
	if err != nil {
		return err
	}
	//совераем действие в зависимости от типа клетки
	switch cell.Types {
	case EmptyCell:
		if cell.Entity != nil {
			e.attack(cell)
		}
	case FoodCell:
		//сначала меняем тип клетки
		if err = w.SetCellType(newCord, EmptyCell); err != nil {
			return err
		}
		//а потом увеличиваем энергию
		e.Energy += energyPoint
	case WallCell:
		e.Energy -= energyPoint
	default:
		return fmt.Errorf("cell type is %v, I dont't know this type", cell.Types)
	}
	return nil
}

// attack отвечает за убийство сущности(Entity) в клетке(Cell) и передачи энергии сущности(Entity),
// вы звавщей функцию. Ничего не возвращает.
func (e *Entity) attack(cell *Cell) {
	energy := cell.Entity.Energy
	cell.Entity.Live = false
	cell.Entity = nil
	e.Energy = energy
}

// rotation отвечает за смену угла взгляда на заданное число.
// Повороты зациклены.
func (e *Entity) rotation(turnCount turns) {
	e.turn = (e.turn + turnCount) % 8
}

// recycling отвечает за получение энергии из загрязнения окружающей среды.
// Возвращает nil или ошибку.
func (e *Entity) recycling(w *World) error {
	const (
		level1 = 5
		level2 = 25
		level3 = 50
		level4 = 75
	)

	//получаем координаты переработки
	newCord := viewCell(e.turn)
	//смотрим что там
	cell, err := w.GetCellData(
		Sum(
			newCord,
			e.Coordinates))
	if err != nil {
		return err
	}

	//Расчитываем размер очистки клетки
	var dPoison = 0
	if cell.Poison >= level4 {
		dPoison = energyPoint * 2
	} else if cell.Poison >= level3 {
		dPoison = energyPoint
	} else if cell.Poison >= level2 {
		dPoison = energyPoint / 2
	} else if cell.Poison >= level1 {
		dPoison = energyPoint / 5
	}

	//очищаем клетку
	if err = w.SetCellPoison(newCord, cell.Poison-dPoison); err != nil {
		return err
	}

	return nil
}

// reproduction is todo!
func (e *Entity) reproduction() error {
	return nil
}

// jump обеспечивает зацикленный прыжок по DNA.Array.
func (e *Entity) jump() {
	e.Pointer += (e.Pointer + e.Array[e.Pointer]) % lengthDNA
}

// loopPointer обеспечивает зацикленность DNA.Pointer.
func (e *Entity) loopPointer() {
	e.Pointer = e.Pointer % lengthDNA
}
