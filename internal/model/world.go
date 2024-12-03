package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func NewWorld(x, y, population, poison int) *World {
	w := &World{
		Xsize:       x,
		Ysize:       y,
		Map:         newMap(x, y, poison),
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
		//если клетка не жива
		//если у неё кончилась энергия
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
			w.Map[x][y].Poison = 0
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
		//Считаем, сколько еды максимум может быть в мире
		maxFood := (w.Xsize * w.Ysize) * MaxFoodPercent / 100
		//Если больше - выходим
		if w.CountFood >= maxFood {
			return
		}
		//Добавляем в мир еды
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
			//Если там есть еда - следующий цикл
			if cell.Types == FoodCell {
				continue
			}
			//Если нет бота, она пустая и не сильно отравлена, то добавляем еду
			if cell.Entity != nil && cell.Types == EmptyCell && cell.Poison < pLevel4 {
				cell.Types = FoodCell
				w.CountFood++
			}
			//Если еды много, выходим
			if w.CountFood >= maxFood {
				break
			}
		}
	}
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
func (w *World) MoveEntity(oldCord, newCord Coordinates, entity *Entity) (err error) {
	newCord, err = w.loopCoord(newCord)

	//Смотрим что в целевой клетке
	cell, err := w.GetCellData(newCord)
	if err != nil {
		//Если не можем посмотреть на клетку - выходим с ошибкой
		return err
	}
	if cell.Entity != nil {
		//Если в другой клетке есть сущность - мы не можем двигаться
		return fmt.Errorf("world move e in %v is fall - have entity №%v", newCord, cell.Entity.ID)
	}
	//Смотрим что в клетке
	switch cell.Types {
	case EmptyCell:
		//todo добавить проверку на ошибку
		_ = w.SetCellEntity(oldCord, nil)
		_ = w.SetCellEntity(newCord, entity)
	case FoodCell:
		//todo добавить проверку на ошибку
		_ = w.SetCellEntity(oldCord, nil)
		_ = w.SetCellEntity(newCord, entity)
		//Уничтожаем еду в клетке - сущность её затоплато
		_ = w.SetCellType(newCord, EmptyCell)
	case WallCell:
		return fmt.Errorf("world move e in %v is fall - wall", newCord)
	}

	return nil
}

// UpdateStat обновляет значение World Statistic высчитывая все живые сущности(Entity),
// подсчитывая клетки с едой и собирая общее коллличество яда в мире.
func (w *World) UpdateStat() {
	//собрать данные по колличеству сущностей
	Count := 0
	for _, entity := range w.ArrayEntity {
		if entity.Live {
			Count++
		}
	}
	w.CountEntity = Count

	//Собрать данные по пище
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

	//Собрать данные по отравлению
	Count = 0
	for x := 0; x < len(w.Map); x++ {
		for y := 0; y < len(w.Map[x]); y++ {
			cell, _ := w.GetCellData(Coordinates{x, y})
			Count += cell.Poison
		}
	}
	w.CountPoison = Count

	//Рассчитаем на сколько мир отравлен
	Count = w.Xsize * w.Ysize * PLevelMax
	w.PercentPoison = float32(w.CountPoison * 100.0 / Count)
}

// SetGeneration приводит отработавщую популяцию к стартовому состоянию с заменой генома.
// разбрасывает сущности по карте в случайном порядке.
func (w *World) SetGeneration(endPopulation, mutationCount int) {
	//отсортируем сущности мо возрасту
	//определив лучшие сущности
	w.sortAge()
	//присвоим их геном остальным ботам
	for i := 0; i < endPopulation; i++ {
		for j := 0; j < endPopulation; j++ {
			w.ArrayEntity[i*endPopulation+j].DNA.Set(w.ArrayEntity[i].DNA)
		}
	}
	//случайно произведём мутации в генокоде
	length := len(w.ArrayEntity)
	for i := 0; i < mutationCount; i++ {
		w.ArrayEntity[rand.Intn(length)].DNA.Mutation(rand.Intn(MaxGen))
	}
	for _, entity := range w.ArrayEntity {
		entity.Energy = rand.Intn(10) + 90
		entity.Age = 0
		entity.Live = true
		entity.Coordinates = Coordinates{
			rand.Intn(w.Xsize),
			rand.Intn(w.Ysize),
		}
	}
	w.sync()
}

// SetCellType изменяет тип клетки(Cell), на указанный. Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellType(cord Coordinates, types CellTypes) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Types = types
		return nil
	}
	return fmt.Errorf("set cell.Types in %v is fall - out of range", cord)
}

// SetCellPoison изменяет уровень яда в клетке(Cell). Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellPoison(cord Coordinates, dPoison int) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Poison += dPoison
		if w.Map[cord.X][cord.Y].Poison == PLevelMax {
			w.Map[cord.X][cord.Y].Poison = PLevelMax
		}
		return nil
	}
	return fmt.Errorf("set cell.Poison in %v is fall - out of range", cord)
}

// SetCellEntity изменяет сущность(Entity) в клетке(Cell). Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellEntity(cord Coordinates, entity *Entity) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Entity = entity
		return nil
	}
	return fmt.Errorf("set cell.Entity in %v is fall - out of range", cord)
}

// GetCellData возвращает указатель на клетку(*Cell) по заданным координатам или ошибку,
// если координаты оказались вне мира.
func (w *World) GetCellData(cord Coordinates) (*Cell, error) {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		return w.Map[cord.X][cord.Y], nil
	}
	return nil, fmt.Errorf("get cell data in %v is fall - out of range", cord)
}

// GetStatistic возвращает строку статистики типа:
//
// [STS] id: *** age: *** e_c: *** f_c: *** p_c: ***
func (w *World) GetStatistic() string {
	return "[STS] id:" + strconv.Itoa(w.ID) +
		" age:" + strconv.Itoa(w.Age) +
		" e_c:" + strconv.Itoa(w.CountEntity) +
		" f_c:" + strconv.Itoa(w.CountFood) +
		" p_c:" + strconv.Itoa(w.CountPoison)
}

// GetPrettyStatistic возвращает статистически данные
// в удобном форматировании.
func (w *World) GetPrettyStatistic() string {
	return "World №" + strconv.Itoa(w.ID) + "    \n" +
		"Age:    " + strconv.Itoa(w.Age) + "    \n" +
		"Entity: " + strconv.Itoa(w.CountEntity) + "    \n" +
		"Food:   " + strconv.Itoa(w.CountFood) + "    \n" +
		"Poison: " + strconv.Itoa(int(w.PercentPoison))
}

// GetEntityInfo возвращает массив строк лучших по возрасту сущностей(Entity)
// Сортирует массив сущностей(Entity) у вызвающего мира (World).
func (w *World) GetEntityInfo(countEntity int) []string {
	w.sortAge()
	s := make([]string, countEntity)
	for i := 0; i < countEntity; i++ {
		s = append(s, w.ArrayEntity[i].DNA.GetDNAString())
	}
	return s
}

// GetPrettyEntityInfo возвращает форматированную строку лучших по возрасту сущностей(Entity)
// Сортирует массив сущностей(Entity) у вызвающего мира (World).
func (w *World) GetPrettyEntityInfo(countEntity int) string {
	w.sortAge()

	var s strings.Builder

	for i := 0; i < countEntity; i++ {
		s.WriteString(
			"ID:" + strconv.Itoa(w.ArrayEntity[i].Age) + " " +
				"Age:" + strconv.Itoa(w.ArrayEntity[i].Age) + " \n" +
				w.ArrayEntity[i].DNA.GetDNAString() + " \n")
	}

	return s.String()
}

// newMap возвращает пустую карту мира.
func newMap(Xsize, Ysize, Poison int) [][]*Cell {
	//создаём массив карты (содержащий строки клеток)
	Map := make([][]*Cell, Xsize)
	for x := 0; x < Xsize; x++ {
		//создаём массив строки (содеижит клетки)
		Map[x] = make([]*Cell, Ysize)
		for y := 0; y < Ysize; y++ {
			Map[x][y] = &Cell{
				nil,
				EmptyCell,
				Poison,
			}
		}
	}
	return Map
}

// sync - функция базовой синхронизации пустого! мира(World) и массива сущностей(Entity[]).
// Если сущность(Entity) оказалась за краем мира - рандомного размещает в мире. Ничего не возвращает
func (w *World) sync() {
	for _, entity := range w.ArrayEntity {
		err := w.SetCellEntity(entity.Coordinates, entity)
		if err != nil {
			for {
				newCoord := Coordinates{
					rand.Intn(w.Xsize),
					rand.Intn(w.Ysize),
				}
				cell, _ := w.GetCellData(newCoord)
				if cell.Entity == nil {
					_ = w.SetCellEntity(entity.Coordinates, entity)
					break
				}
			}
		}
	}
}

// newGeneration создаёт стартовую популяцию сущностей(Entity). Возращает массив ссылок на Entity.
func newGeneration(x, y, population int) []*Entity {
	entityArray := make([]*Entity, population)
	for i := 0; i < population; i++ {
		entityArray[i] = NewEntity(i, rand.Intn(x), rand.Intn(y), LengthDNA)
	}
	return entityArray
}

// sortAge сортирует сущности(Entity) по возрасту в вызывающем мире(World).
func (w *World) sortAge() {

	n := len(w.ArrayEntity)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if w.ArrayEntity[j].Age < w.ArrayEntity[j+1].Age {
				w.ArrayEntity[j], w.ArrayEntity[j+1] = w.ArrayEntity[j+1], w.ArrayEntity[j]
			}
		}
	}
}

// loopCoord - отвечает за перенос координат, выходящих за границу мира.
func (w *World) loopCoord(old Coordinates) (Coordinates, error) {
	//todo: просто сделать лучше и чтобы работало
	new := Coordinates{
		old.X,
		old.Y,
	}
	if LoopX {
		if old.X < 0 {
			new.X = w.Xsize + (old.X % w.Xsize)
		} else if old.X > w.Xsize {
			new.X = old.X % w.Xsize
		}
	}
	if LoopY {
		if old.Y < 0 {
			new.Y = w.Ysize + (old.Y % w.Ysize)
		} else if old.Y > w.Ysize {
			new.Y = old.X % w.Ysize
		}
	}
	if !checkLimit(new, Coordinates{w.Xsize, w.Ysize}) {
		return old, fmt.Errorf("У нас тут ошибка!")
	}
	return new, nil
}
