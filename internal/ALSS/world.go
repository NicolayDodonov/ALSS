package ALSS

type cell struct {
	Agent         *agent
	Height        int
	localMinerals int
	localGrass    int
}

type wGlobal struct {
	Temperature  int
	Illumination int
	Pollution    int
	SeaLevel     int
}

type wStat struct {
	Year       int
	LiveAgent  int
	DeathAgent int

	MaxMinerals int
	AVGMinerals int
	TotMinerals int

	MaxPollution int
	AVGPollution int
	TotPollution int

	MaxGrass int
	AVGGrass int
	TotGrass int
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

func (w *world) init() {
	
}

func (w *world) clearWorld() {

}
