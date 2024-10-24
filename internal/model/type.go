package model

type World struct {
	Xsize       int
	Ysize       int
	Map         [][]Cell
	ArrayEntity []Entity
	Statistic
}

type Statistic struct {
	CountEntity int
	CountFood   int
	CountPoison int
	Age         int
	ID          int
}

// Изменить структуру типов клеток
type Cell struct {
	*Entity
	Poison int
	Food   int
	Wall   bool
}

type Coordinates struct {
	X int
	Y int
}

type Entity struct {
	ID     int
	Age    int
	Energy int
	turn   turns
	DNA
}

type turns int

type DNA struct {
	ID      int
	Pointer int
	Array   []int
}
