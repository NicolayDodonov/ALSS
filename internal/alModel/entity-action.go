package alModel

import "fmt"

// move отвечает за передвижение сущности(Entity) из одной клетки(Cell) мира(World) в другую.
// Возвращает nil или ошибку.
func (e *Entity) move(w *World) error {
	// Получаем координаты, куда хотим переместиться
	newCord := sum(
		viewCell(e.turn),
		e.Coordinates,
	)
	newCord, _ = w.loopCord(newCord)
	// Перемещаемся в новые координаты
	if err := w.MoveEntity(newCord, e); err != nil {
		return err
	}
	e.Coordinates = newCord
	return nil
}

// look отвечает за получение данных из другой клетки(Cell). Возвращает номер
// сдвига Entity.DNA или ошибку.
func (e *Entity) look(w *World) (int, error) {
	// Константы ответов на что мы смотрим
	const (
		isError = iota
		isEmpty
		isFood
		isWall
		isEntity
	)

	// Получаем координаты, куда хотим посмотреть
	newCord := sum(
		viewCell(e.turn),
		e.Coordinates)
	newCord, _ = w.loopCord(newCord)
	// Смотрим что там
	cell, err := w.GetCellData(newCord)
	if err != nil {
		return isError, err
	}

	// Определяем тип возврата
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
	// Получаем координаты для взятия
	newCord := sum(
		viewCell(e.turn),
		e.Coordinates)
	newCord, _ = w.loopCord(newCord)
	// Смотрим что там
	cell, err := w.GetCellData(newCord)
	if err != nil {
		return err
	}
	// Совераем действие в зависимости от типа клетки
	switch cell.Types {
	case EmptyCell:
		// клетка пуста - выходим
		if err = e.attack(cell); err != nil {
			return err
		}
	case FoodCell:
		// Сначала меняем тип клетки
		cell.Types = EmptyCell
		// А потом увеличиваем энергию
		e.Energy += EnergyPoint
	case WallCell:
		// Стена - наказываем за удар
		e.Energy -= EnergyPoint
	default:
		return fmt.Errorf("cell type is %v, I dont't know this type", cell.Types)
	}
	return nil
}

// attack отвечает за убийство сущности(Entity) в клетке(Cell) и передачи энергии сущности(Entity),
// вы звавщей функцию. Ничего не возвращает.
func (e *Entity) attack(cell *Cell) error {
	if cell.Entity == nil {
		return fmt.Errorf("attack is fall - not entity")
	}
	energy := cell.Entity.Energy
	cell.Entity.Live = false
	cell.Entity = nil
	e.Energy = energy
	return nil
}

// rotation отвечает за смену угла взгляда на заданное число.
// Повороты зациклены.
func (e *Entity) rotation(turnCount turns) {
	e.turn = (e.turn + turnCount) % 8
}

// recycling отвечает за получение энергии из загрязнения окружающей среды.
// Возвращает nil или ошибку.
func (e *Entity) recycling(w *World) error {

	// Получаем координаты переработки
	newCord := viewCell(e.turn)
	newCord, _ = w.loopCord(newCord)
	// Смотрим что там
	cell, err := w.GetCellData(
		sum(
			newCord,
			e.Coordinates))
	if err != nil {
		return err
	}

	// Расчитываем размер очистки клетки
	var dPoison = 0
	if cell.Poison >= pLevel4 {
		dPoison = EnergyPoint * 2
	} else if cell.Poison >= pLevel3 {
		dPoison = EnergyPoint
	} else if cell.Poison >= pLevel2 {
		dPoison = EnergyPoint / 2
	} else if cell.Poison >= pLevel1 {
		dPoison = EnergyPoint / 5
	}

	// Очищаем клетку
	if err = w.SetCellPoison(newCord, cell.Poison-dPoison); err != nil {
		return err
	}

	return nil
}

// reproduction is todo!
func (e *Entity) reproduction() error {
	return nil
}
