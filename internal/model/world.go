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
		},
	}
	w.Sync()
	return w
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

// newGeneration создаёт стартовую популяцию сущностей(Entity). Возращает массив ссылок на Entity.
func newGeneration(x, y, population int) []*Entity {
	entityArray := make([]*Entity, population)
	for i := 0; i < population; i++ {
		entityArray[i] = NewEntity(i, rand.Intn(x), rand.Intn(y), LengthDNA)
	}
	return entityArray
}

// RemoveDead очищает мир от умерших ботов, чтобы живые с ними не взаимодействовали.
func (w *World) RemoveDead() {
	for _, entity := range w.ArrayEntity {
		if !entity.Live {
			_ = w.SetCellEntity(entity.Coordinates, nil)
		}
	}
}

// Sync отвечает за синхронизацию World.ArrayEntity c World.Map.
// если несколько сущностей оказывается в одной клетке, для всех последующих
// создаёт новое расположение.
func (w *World) Sync() {
	w.RemoveDead()
	for _, entity := range w.ArrayEntity {
		if entity.Live {
			//если по коордитанам сущности расположена другая сущность
			cell, _ := w.GetCellData(entity.Coordinates)
			if cell.Entity != nil &&
				cell.Entity != entity {
				//ищем пустую клетку
				for {
					x := rand.Intn(w.Xsize)
					y := rand.Intn(w.Ysize)
					cell, _ := w.GetCellData(Coordinates{x, y})
					if cell.Entity == nil &&
						cell.Types == EmptyCell {
						//и записываем туда нащу сущность
						entity.Coordinates = Coordinates{x, y}
						cell.Entity = entity
						break
					}
				}
			}
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
}

// Update обновляет состояние всех клеток(Cell) вызвавщего функцию мира(World)
// создава новые ресурсы, удаля ресурсы из за отравления.
func (w *World) Update(percent int) {
	maxFood := percent * (w.Xsize * w.Ysize) / 100
	if w.CountFood >= maxFood {
		return
	}

	//Пройдёмся по всем клеткам мира
	for countAdd := w.CountFood; countAdd < maxFood; {
		rX := rand.Intn(w.Xsize)
		rY := rand.Intn(w.Ysize)
		cell, _ := w.GetCellData(Coordinates{rX, rY})
		if cell.Types == EmptyCell {
			cell.Types = FoodCell
			countAdd++
		}
	}
}

// Execute выполняет генетический код для каждой сущности(Entity) вызвавщего
// функцию мира(World). Возвращает nil или ошибку исполнения сущности.
func (w *World) Execute() (err []error) {
	for _, entity := range w.ArrayEntity {
		if errEntity := entity.Run(w); errEntity != nil {
			err = append(err, errEntity)
		}
	}
	return err
}

// MoveEntity передвигает сущность(Entity) из старой клетки(Cell) в новую.
// Возвращает nil или ошибку.
func (w *World) MoveEntity(oldCord, newCord Coordinates, entity *Entity) error {
	if checkLimit(newCord, Coordinates{w.Xsize, w.Ysize}) {
		cell, _ := w.GetCellData(newCord)
		if cell.Entity == nil && cell.Types != WallCell {
			w.Map[oldCord.X][oldCord.Y].Entity = nil
			w.Map[newCord.X][newCord.Y].Entity = entity
			entity.Coordinates = newCord
			return nil
		} else if cell.Entity != nil {
			return fmt.Errorf("world move e in %v is fall - have entity №%v", newCord, cell.Entity.ID)
		} else {
			return fmt.Errorf("world move e in %v is fall - have wall", newCord)
		}
	}
	return fmt.Errorf("world move e n %v is fall - out of range", newCord)
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
		w.ArrayEntity[rand.Intn(length)].DNA.Mutation(rand.Intn(length))
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
	w.Sync()
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
		w.Map[cord.X][cord.Y].Poison = dPoison
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
	return "World №" + strconv.Itoa(w.ID) + "\n" +
		"Age:    " + strconv.Itoa(w.Age) + "\n" +
		"Entity: " + strconv.Itoa(w.CountEntity) + "\n" +
		"Food:   " + strconv.Itoa(w.CountFood) + "\n" +
		"Poison: " + strconv.Itoa(w.CountPoison)
}

// GetPrettyEntityInfo возвращает массив строк с лучших по возрасту
// сущностей(Entity). Одновременно с этим сортирует весь массив
// сущностей(Entity).
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

// neighbors - функция клеточного автомата, смотрит на состояние соседей клетки
// и определяет, становиться ли клетка едой или нет. True - еда, False - empty.
// Временно не используется.
func (w *World) neighbors(c Coordinates) bool {
	if w.Map[c.X][c.Y].Poison >= pLevel4 {
		//Если в самой клетке очень много яда, то она всегда пустая
		//Если была еда = она погибает!
		return false
	}

	if w.Map[c.X][c.Y].Types == FoodCell {
		//Если клетка уже еда то true
		return true
	}

	//смотрим на всех соседей во круг клетки и считаем колличество еды вокруг
	countFood := 0
	for i := 0; i < 8; i++ {
		//получаем клетку соседа
		cell, err := w.GetCellData(viewCell(turns(i)))
		if err != nil {
			//Выход за границу мира нас не волнует
			//Для нас это сверх ядовитые клетки
			continue
		}
		if cell.Types == FoodCell {
			if cell.Poison <= pLevel3 {
				//Если в клетке есть еда и уровень яда
				//меньше половины - для нас это нормально
				countFood++
			}
		}
	}

	if countFood >= 1 && countFood < 3 {
		return true
	}
	return false
}

// RandomFood случайным образом распологает в мире пищу.
// Временно не используется
func (w *World) randomFood(percent int) {
	for x := 0; x < w.Xsize; x++ {
		for y := 0; y < w.Ysize; y++ {
			if rand.Intn(100) < percent {
				w.Map[x][y].Types = FoodCell
			}
		}
	}
}
