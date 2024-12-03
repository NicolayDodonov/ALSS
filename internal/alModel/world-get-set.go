package alModel

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

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

// GetEntityInfo возвращает массив строк лучших по возрасту сущностей(Entity)
// Сортирует массив сущностей(Entity) у вызвающего мира (World).
func (w *World) GetEntityInfo(countEntity int) []string {
	w.sortAge()
	s := make([]string, countEntity)
	for i := 0; i < countEntity; i++ {
		s = append(s, w.ArrayEntity[i].DNA.String())
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
				w.ArrayEntity[i].DNA.String() + " \n")
	}

	return s.String()
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
