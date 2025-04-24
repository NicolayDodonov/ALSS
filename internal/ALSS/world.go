package ALSS

type global struct {
	Temperature  int
	Illumination int
	Pollution    int
	SeaLevel     int
}

type worldStatistic struct {
	Year       int
	LiveAgent  int
	DeathAgent int

	countCell      int
	underwaterCell int

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

type cell struct {
	Agent         *agent
	Height        int
	localMinerals int //0 ... 255
	localGrass    int //0 ... 255
}

type Map [][]cell

// world структура описывающая структуру мира модели
// и реализующий основные методы работы с миром модели
type world struct {
	Map  Map
	MaxX int
	MaxY int
	global
	worldStatistic
}

func newWorld() *world {
	return &world{
		global: global{
			Temperature:  0,
			Illumination: 0,
			Pollution:    0,
			SeaLevel:     0,
		},
		worldStatistic: worldStatistic{
			Year:         0,
			LiveAgent:    0,
			DeathAgent:   0,
			MaxMinerals:  0,
			AVGMinerals:  0,
			TotMinerals:  0,
			MaxPollution: 0,
			AVGPollution: 0,
			TotPollution: 0,
			MaxGrass:     0,
			AVGGrass:     0,
			TotGrass:     0,
		},
	}
}

func (w *world) initMap(season string) {
	//создать карту и заполнить её клетками
	w.newMap()
	//вызвать генератор высотности карты
	w.landscapeGenerator()
	//запустить клеточный автомат травы
	w.grassHandler()
	//установить сезон
	w.changeSeason(season)
}

func (w *world) update() {
	w.grassHandler()
	w.mineralHandler()
	w.updateStat()
}

func (w *world) newMap() {
	newMap := make(Map, 0)
	for y := 0; y < w.MaxY; y++ {
		newMap = append(newMap, []cell{})
		for x := 0; x < w.MaxX; x++ {
			newMap[y] = append(newMap[y], cell{
				Agent:         nil,
				Height:        0,
				localMinerals: 0,
				localGrass:    0,
			})
		}
	}
	w.Map = newMap
}

func (w *world) landscapeGenerator() {
	//todo: создать шум перлина заданого размера
	//todo: задать каждой клетке мира высоту с учётом этого шума
}

// grassHandler оператор клеточного автомата для динамического изменения травы в world.Map.
func (w *world) grassHandler() {
	//todo: условия добавления новой растительности
	//todo: условие удаления травы из клетки
}

func (w *world) mineralHandler() {
	//todo: превратить часть загрязнения в минералы в затопленных клетках
}

func (w *world) updateStat() {
	//todo: пересчитать динамическую статистику мира
}

func (w *world) changeSeason(s string) {
	//todo: реализовать смену времён года
	switch s {
	case spring:
		w.Illumination = 20
		w.Temperature = 20
		w.SeaLevel = 10
	case summer:
		w.Illumination = 40
		w.Temperature = 30
		w.SeaLevel = 15
	case autumn:
		w.Illumination = 20
		w.Temperature = 10
		w.SeaLevel = 10
	case winter:
		w.Illumination = 15
		w.Temperature = 0
		w.SeaLevel = 5
	default:
		w.Illumination = 1
		w.Temperature = 1
		w.SeaLevel = 1
	}
}
