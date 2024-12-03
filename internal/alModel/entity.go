package alModel

import (
	l "artificialLifeGo/internal/logger"
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
		*newDNA(longDNA),
		newBrain(),
	}
}

// Run отвечает за исполнение генетического кода в DNA.Array.
// Возвращает nil или критическую ошибку
func (e *Entity) Run(w *World) {
	l.Ent.Debug("id " + strconv.Itoa(e.ID) + " is run his genocode")
	// Если бот мёрт, вылетаем с ошибкой
	if !e.Live {
		l.Ent.Debug("ID:" + strconv.Itoa(e.ID) + "cant run - dead")
		return
	}

	// Уменьшаем энергию бота перед выполнение генокода
	// "Деньги в перёд"
	e.Energy--
	e.Age++

	// Вызываем мозг бота для исполнения команды
	err := e.run(e, w)
	if err != nil {
		l.Ent.Error("id" + strconv.Itoa(e.ID) + " " + err.Error())
		return
	}

	// Проверяем клетку бота на уровень яда
	// Тут бот может умереть, но код ниже это всё равно обработает
	err = e.poisonHandler(w)

	// Если энергии не осталось - умираем
	if e.Energy <= 0 || !e.Live {
		e.die(w)
		l.Ent.Info("ID:" + strconv.Itoa(e.ID) + " die without energy")
		return
	}
}

// poisonHandler обрабатывает все условия взаимодействия бота с отравлением на месности.
// Возвращает nil или ошибку выхода за границы мира
func (e *Entity) poisonHandler(w *World) error {
	// Добавляем отравление на клетку, где находиться бот
	if err := w.SetCellPoison(e.Coordinates, pLevel1+1); err != nil {
		return err
	}

	// Получаем исчерпывающую информацию о клетке, где находится бот
	cell, err := w.GetCellData(e.Coordinates)
	if err != nil {
		return err
	}

	// Проверяем уровень яда
	if cell.Poison >= pLevelDed {
		e.Live = false
		return nil
	} else if cell.Poison >= pLevel3 {
		e.Energy -= pLevel1 * 5
		return nil
	}
	return nil
}

// loopPointer обеспечивает зацикленность DNA.Pointer.
func (e *Entity) loopPointer() {
	e.Pointer = e.Pointer % LengthDNA
}

// die устанавливает бота в умершее состояние.
// Удаляет ссылку на бота из мира. Первый уровень защиты от умерших ботов.
func (e *Entity) die(w *World) {
	e.Live = false
	e.Energy = 0
	// Очищаем клетку от сущности
	_ = w.SetCellEntity(e.Coordinates, nil)
}
