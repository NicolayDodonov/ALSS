package ALSS

type cell struct {
	Agent *agent
	//todo: other info
}

type wGlobal struct {
	Temperature  int
	Illumination int
	Pollution    int
	SeaLevel     int
}

type wStat struct {
	Number     int
	Year       int
	LiveAgent  int
	DeathAgent int
	AllAgent   int
	Minerals   int
	Grass      int
}

// world структура описывающая структуру мира модели
// и реализующий основные методы работы с миром модели
type world struct {
	Map  Map
	MaxX int
	MaxY int
	wGlobal
	wStat
}

type Map [][]cell

func newWorld() *world {
	return &world{}
}

func (w *world) clearWorld() {

}

func (w *world) getCell(c coordinates) *cell {
	//todo: fix out of bonding
	return &w.Map[c.X][c.Y]
}
