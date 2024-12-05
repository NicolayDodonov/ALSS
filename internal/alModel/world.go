package alModel

import (
	"fmt"
	"math/rand"
)

// NewWorld возвращает структуру мира(World).
func NewWorld(x, y, population int) *World {
	w := &World{
		Xsize:       x,
		Ysize:       y,
		Map:         newMap(x, y, BasePoisonLevel),
		ArrayEntity: newGeneration(x, y, population),
		Statistic: Statistic{
			population,
			0,
			0,
			0,
			0,
			0,
		},
	}
	w.sync()
	return w
}

// RemoveDead очищает мир от умерших сущностей(Entity), чтобы живые с ними не взаимодействовали.
// Является вторым уровнем защиты от умерших сущностей(Entity).
func (w *World) RemoveDead() {
	for _, entity := range w.ArrayEntity {
		// Если клетка не жива
		// Если у неё кончилась энергия
		if !entity.Live ||
			entity.Energy <= 0 {
			_ = w.SetCellEntity(entity.Coordinates, nil)
		}
	}
}

// Clear приводит все клетки(Cell) вызвавщего функцию мира(World) в стандартное состояние.
func (w *World) Clear() {
	for x := 0; x < len(w.Map); x++ {
		for y := 0; y < len(w.Map[x]); y++ {
			w.Map[x][y].Entity = nil
			w.Map[x][y].Types = EmptyCell
			w.Map[x][y].Poison = BasePoisonLevel
		}
	}
	w.CountFood = 0
	w.CountEntity = 0
	w.CountPoison = 0
	w.PercentPoison = 0
}

// Update обновляет состояние всех клеток(Cell) вызвавщего функцию мира(World)
// создава новые ресурсы, удаля ресурсы из за отравления.
func (w *World) Update(addFood bool) {
	// Проход по всем клеткам удаляя еду из за отравления или ботов
	for _, cells := range w.Map {
		for _, cell := range cells {
			if cell.Poison >= pLevel4 && cell.Types == FoodCell {
				cell.Types = EmptyCell
				w.CountFood--
			}
			if cell.Entity != nil && cell.Types == FoodCell {
				cell.Types = EmptyCell
				cell.Entity.Energy++
				w.CountFood--
			}
		}
	}
	//Если надо добавить в мир еды, то:
	if addFood {
		// Считаем, сколько еды максимум может быть в мире
		maxFood := (w.Xsize * w.Ysize) * MaxFoodPercent / 100
		// Если больше - выходим
		if w.CountFood >= maxFood {
			return
		}
		// Добавляем в мир еды
		for attempt := 0; attempt < maxFood*2; attempt++ {
			// Берём случайную клетку
			cell, err := w.GetCellData(
				Coordinates{
					rand.Intn(w.Xsize),
					rand.Intn(w.Ysize),
				})
			if err != nil {
				break
			}
			// Если там есть еда - следующий цикл
			if cell.Types == FoodCell {
				continue
			}
			// Если нет бота, она пустая и не сильно отравлена, то добавляем еду
			if cell.Entity == nil && cell.Types == EmptyCell && cell.Poison < pLevel4 {
				cell.Types = FoodCell
				w.CountFood++
			}
			// Если еды много, выходим
			if w.CountFood >= maxFood {
				break
			}
		}
	}
}

// StatisticUpdate обновляет значение World Statistic высчитывая все живые сущности(Entity),
// подсчитывая клетки с едой и собирая общее коллличество яда в мире.
func (w *World) StatisticUpdate() {
	// Собрать данные по колличеству сущностей
	Count := 0
	for _, entity := range w.ArrayEntity {
		if entity.Live {
			Count++
		}
	}
	w.CountEntity = Count

	// Собрать данные по пище
	Count = 0
	for x := 0; x < len(w.Map); x++ {
		for y := 0; y < len(w.Map[x]); y++ {
			cell, _ := w.GetCellData(Coordinates{x, y})
			if cell.Types == FoodCell {
				Count++
			}
		}
	}
	w.CountFood = Count

	// Собрать данные по отравлению
	Count = 0
	for x := 0; x < len(w.Map); x++ {
		for y := 0; y < len(w.Map[x]); y++ {
			cell, _ := w.GetCellData(Coordinates{x, y})
			Count += cell.Poison
		}
	}
	w.CountPoison = Count

	// Рассчитаем на сколько мир отравлен
	Count = w.Xsize * w.Ysize * PLevelMax
	w.PercentPoison = w.CountPoison * 100.0 / Count
}

// Execute выполняет генетический код для каждой сущности(Entity) вызвавщего
// функцию мира(World). Возвращает nil или ошибку исполнения сущности.
func (w *World) Execute() {
	for _, entity := range w.ArrayEntity {
		entity.Run(w)
	}
}

// MoveEntity передвигает сущность(Entity) из старой клетки(Cell) в новую.
// Возвращает nil или ошибку перемещения.
func (w *World) MoveEntity(newCord Coordinates, entity *Entity) (err error) {
	// Смотрим что в целевой клетке
	cell, err := w.GetCellData(newCord)
	if err != nil {
		//Если не можем посмотреть на клетку - выходим с ошибкой
		return err
	}
	if cell.Entity != nil {
		//Если в другой клетке есть сущность - мы не можем двигаться
		return fmt.Errorf("world move e in %v is fall - have entity №%v", newCord, cell.Entity.ID)
	}
	// Смотрим что в клетке
	switch cell.Types {
	case EmptyCell:
		if err = w.SetCellEntity(entity.Coordinates, nil); err != nil {
			return err
		}
		if err = w.SetCellEntity(newCord, entity); err != nil {
			return err
		}
	case FoodCell:
		if err = w.SetCellEntity(entity.Coordinates, nil); err != nil {
			return err
		}
		if err = w.SetCellEntity(newCord, entity); err != nil {
			return err
		}
		// Уничтожаем еду в клетке - сущность её затоплато
		if err = w.SetCellType(newCord, EmptyCell); err != nil {
			return err
		}
	case WallCell:
		return fmt.Errorf("world move e in %v is fall - wall", newCord)
	}

	return nil
}

// Wall в случайном порядке создаёт стены. Принимает в качесте аргумента колличество стен.
func (w *World) Wall(count int) {
	for i := 0; i < count; i++ {
		direction := rand.Intn(10)
		cord := Coordinates{
			X: rand.Intn(w.Xsize),
			Y: rand.Intn(w.Ysize),
		}
		if direction%2 == 0 {
			long := rand.Intn(w.Xsize/2) + 5

			for j := 0; j < long; j++ {
				cord.X++
				if err := w.SetCellType(cord, WallCell); err != nil {
					break
				}
			}
		} else {
			long := rand.Intn(w.Ysize/2) + 5
			for j := 0; j < long; j++ {
				cord.Y++
				if err := w.SetCellType(cord, WallCell); err != nil {
					break
				}
			}
		}
	}
}
