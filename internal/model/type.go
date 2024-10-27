package model

type World struct {
	Xsize       int
	Ysize       int
	Map         [][]*Cell
	ArrayEntity []*Entity
	Statistic
}

type Statistic struct {
	CountEntity int
	CountFood   int
	CountPoison int
	Age         int
	ID          int
}

// todo: Изменить структуру типов клеток
type Cell struct {
	*Entity
	Types  CellTypes
	Poison int
}

type CellTypes int

type Coordinates struct {
	X int
	Y int
}

type Entity struct {
	ID     int
	Age    int
	Energy int
	Live   bool
	turn   turns
	Coordinates
	DNA
}

type turns int

type DNA struct {
	ID      int
	Pointer int
	Array   []int
}
