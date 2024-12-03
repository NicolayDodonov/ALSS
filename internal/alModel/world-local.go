package alModel

import (
	"fmt"
	"math/rand"
)

// newMap возвращает пустую карту мира.
func newMap(Xsize, Ysize, Poison int) Map {
	//создаём массив карты (содержащий строки клеток)
	Map := make(Map, Xsize)
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
				newCord := Coordinates{
					rand.Intn(w.Xsize),
					rand.Intn(w.Ysize),
				}
				cell, _ := w.GetCellData(newCord)
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

// loopCord - отвечает за перенос координат, выходящих за границу мира.
func (w *World) loopCord(old Coordinates) (Coordinates, error) {
	//todo: просто сделать лучше и чтобы работало

	if LoopX {
		if old.X < 0 {
			old.X = w.Xsize + (old.X % w.Xsize)
		} else if old.X > w.Xsize {
			old.X = old.X % w.Xsize
		}
	}
	if LoopY {
		if old.Y < 0 {
			old.Y = w.Ysize + (old.Y % w.Ysize)
		} else if old.Y > w.Ysize {
			old.Y = old.X % w.Ysize
		}
	}
	if !checkLimit(old, Coordinates{w.Xsize, w.Ysize}) {
		return old, fmt.Errorf("У нас тут ошибка!")
	}
	return old, nil
}
