package model

// World описывает базисную структуру мира, его линейные размеры,
// а так же содержит массивы карты, сущностей мира и статистики.
type World struct {
	Xsize       int
	Ysize       int
	Map         [][]*Cell
	ArrayEntity []*Entity
	Statistic
}

// Statistic позволяет сохранять статистические данные о мире.
type Statistic struct {
	CountEntity int
	CountFood   int
	CountPoison int
	Age         int
	ID          int
}

// todo: Изменить структуру типов клеток

// Cell описывает базовый эллемент карты мира - его клетку.
type Cell struct {
	*Entity
	Types  CellTypes
	Poison int
}

type CellTypes int

// Coordinates - универсальная структура координат любого объекта в мире.
type Coordinates struct {
	X int
	Y int
}

// Entity описывает базисную структуру сущностей мира. Они обладают
// своим генокодом описанном в DNA.
type Entity struct {
	ID     int
	Age    int
	Energy int
	Live   bool
	turn   turns
	Coordinates
	DNA
	brain
}

type turns int

// DNA - генокод. Array - область программного кода,
// Pointer - указатель на ячейку программного кода.
// ID - уникальный идентификатор генокода.
type DNA struct {
	ID      int
	Pointer int
	Array   []int
}

// brain - интерфейс обработчика генокода в DNA
type brain interface {
	run(*Entity, *World) error
}
