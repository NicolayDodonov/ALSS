package model

import (
	"fmt"
	"math/rand"
	"strconv"
)

func NewWorld(x, y, population int) *World {
	w := &World{
		Xsize:       x,
		Ysize:       y,
		Map:         newMap(x, y),
		ArrayEntity: newGeneration(population, x, y),
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
func newMap(Xsize, Ysize int) [][]*Cell {
	//создаём массив карты (содержащий строки клеток)
	Map := make([][]*Cell, Xsize)
	for x := 0; x < Xsize; x++ {
		//создаём массив строки (содеижит клетки)
		Map[x] = make([]*Cell, Ysize)
		for y := 0; y < Ysize; y++ {
			Map[x][y] = &Cell{
				nil,
				EmptyCell,
				0,
			}
		}
	}
	return Map
}

// newGeneration создаёт стартовую популяцию сущностей(Entity). Возращает массив ссылок на Entity.
func newGeneration(population, x, y int) []*Entity {
	entityArray := make([]*Entity, population)
	for i := 0; i < population; i++ {
		entityArray[i] = NewEntity(i, rand.Intn(x), rand.Intn(y), lengthDNA)
	}
	return entityArray
}

// Sync отвечает за синхронизацию World.ArrayEntity c World.Map.
// если несколько сущностей оказывается в одной клетке, для всех последующих
// создаёт новое расположение.
func (w *World) Sync() {
	for _, entity := range w.ArrayEntity {
		if entity.Live {
			//если по коордитанам сущности расположена другая сущность
			if w.Map[entity.X][entity.Y].Entity != nil &&
				w.Map[entity.X][entity.Y].Entity != entity {
				//ищем пустую клетку
				for {
					x := rand.Intn(w.Xsize)
					y := rand.Intn(w.Ysize)
					if w.Map[entity.X][entity.Y].Entity == nil &&
						w.Map[entity.X][entity.Y].Types == EmptyCell {
						//и записываем туда нащу сущность
						entity.Coordinates = Coordinates{x, y}
						break
					}
				}
			}
		}
	}
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
		entity.Energy = 100
		entity.Age = 0
		entity.Live = true
		entity.Coordinates = Coordinates{
			rand.Intn(w.Xsize),
			rand.Intn(w.Ysize),
		}
	}
	w.Sync()
}

// sortAge сортирует сущности(Entity) по возрасту в вызывающем мире(World).
func (w *World) sortAge() {
	for i := 0; i < len(w.ArrayEntity)-1; i++ {
		for j := 0; j < len(w.ArrayEntity)-i-1; j++ {
			if w.ArrayEntity[j].Age < w.ArrayEntity[j+1].Age {
				w.ArrayEntity[j], w.ArrayEntity[j+1] = w.ArrayEntity[j+1], w.ArrayEntity[j]
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
func (w *World) Update() {
	//todo: придумать нормальное обновление ресурсов с учётом типа яда в клетке
}

// Execute выполняет генетический код для каждой сущности(Entity) вызвавщего
// функцию мира(World). Возвращает nil или ошибку исполнения сущности.
func (w *World) Execute() error {
	for _, entity := range w.ArrayEntity {
		if err := entity.Run(w); err != nil {
			return err
		}
	}
	return nil
}

// GetCellData возвращает указатель на клетку(*Cell) по заданным координатам или ошибку,
// если координаты оказались вне мира.
func (w *World) GetCellData(cord Coordinates) (*Cell, error) {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		return w.Map[cord.X][cord.Y], nil
	}
	return nil, fmt.Errorf("[err] coordinate %v out of range", cord)
}

// SetCellType изменяет тип клетки(Cell), на указанный. Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellType(cord Coordinates, types CellTypes) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Types = types
		return nil
	}
	return fmt.Errorf("[err] coordinate %v out of range", cord)
}

// SetCellPoison изменяет уровень яда в клетке(Cell). Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellPoison(cord Coordinates, dPoison int) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Poison = dPoison
		return nil
	}
	return fmt.Errorf("[err] coordinate %v out of range", cord)
}

// SetCellEntity изменяет сущность(Entity) в клетке(Cell). Возвращает nil или
// ошибку выхода за границы мира.
func (w *World) SetCellEntity(cord Coordinates, entity *Entity) error {
	if checkLimit(cord, Coordinates{w.Xsize, w.Ysize}) {
		w.Map[cord.X][cord.Y].Entity = entity
		return nil
	}
	return fmt.Errorf("[err] coordinate %v out of range", cord)
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
		} else {
			return fmt.Errorf("[err] coordinate %v have another object", newCord)
		}
	}
	return fmt.Errorf("[err] coordinate %v out of range", newCord)
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
