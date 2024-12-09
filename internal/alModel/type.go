package alModel

// World описывает базисную структуру мира, его линейные размеры,
// а так же содержит массивы карты, сущностей мира и статистики.
type World struct {
	Xsize       int
	Ysize       int
	Map         Map
	ArrayEntity []*Entity
	Statistic
}

// Map это тип карты. Слайс слайса указателей клеток(*Cell)
type Map [][]*Cell

// Statistic позволяет сохранять статистические данные о мире.
type Statistic struct {
	CountEntity   int
	CountFood     int
	CountPoison   int
	PercentPoison int
	Age           int
	ID            int
}

// Cell описывает базовый эллемент карты мира - его клетку.
type Cell struct {
	*Entity
	Types  CellTypes
	Poison int
}

// CellTypes это тип клетки.
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

// turns это угол поворота.
type turns int

// DNA - генокод. Array - область программного кода,
// Pointer - указатель на ячейку программного кода.
// ID - уникальный идентификатор генокода.
type DNA struct {
	ID      int
	Pointer int
	Array   []int
}
