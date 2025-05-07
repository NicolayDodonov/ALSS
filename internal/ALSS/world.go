package ALSS

// world структура описывающая структуру мира модели
// и реализующий основные методы работы с миром модели
type world struct {
	Map  Map
	MaxX int
	MaxY int
	global
	worldStatistic
}

type global struct {
	Illumination int
	Pollution    int
	PollutionFix int
	SeaLevel     int
}

type worldStatistic struct {
	Year       int
	LiveAgent  int
	DeathAgent int

	countCell      int
	underwaterCell int

	AVGMinerals int
	TotMinerals int
}

type cell struct {
	Agent         *agent
	Height        int
	localMinerals int //0 ... 255
}

type Map [][]cell

func newWorld(x, y int) *world {
	return &world{
		MaxX: x,
		MaxY: y,
		Map:  nil,
		global: global{
			Illumination: 0,
			Pollution:    0,
			SeaLevel:     0,
		},
		worldStatistic: worldStatistic{
			Year:        0,
			LiveAgent:   0,
			DeathAgent:  0,
			AVGMinerals: 0,
			TotMinerals: 0,
		},
	}
}

// initMap создаёт пустую карту мира, генерирует высотный ланшафт, настраивает уровень моря
// и проводит начальную настройку статистики.
func (w *world) initMap() {
	//создать карту и заполнить её клетками
	w.newMap()
	//вызвать генератор высотности карты
	w.landscapeGenerator(mapGRADIENT)
	//установить сезон
	w.changeSeason(summer)
	//собрать начальную статистику
	w.initStat()
}

func (w *world) newMap() {
	newMap := make([][]cell, w.MaxY)
	for y := range newMap {
		newMap[y] = make([]cell, w.MaxX)
		for x := range newMap[y] {
			newMap[y][x] = cell{
				Agent:         nil,
				Height:        0,
				localMinerals: 0,
			}
		}
	}
	w.Map = newMap
}

func (w *world) landscapeGenerator(landType string) {
	switch landType {
	case mapGRADIENT:
		var gradientLayer int
		if w.MaxY > maxHeight {
			gradientLayer = maxHeight
		} else {
			gradientLayer = 1
		}

		height := maxHeight
		layerCount := gradientLayer
		for y := 0; y < w.MaxX; y++ {
			if layerCount == 0 {
				layerCount = gradientLayer
				height--
			}
			for x := 0; x < w.MaxX; x++ {
				w.Map[y][x].Height = height
			}
			layerCount--
		}
	case mapRANDOM:
	}
}

func (w *world) mineralHandler() {
	dPollution := w.Pollution / 1000
	dMinerals := dPollution / w.underwaterCell
	w.Pollution -= dPollution

	for _, cells := range w.Map {
		for _, cell := range cells {
			if cell.Height <= w.SeaLevel {
				cell.localMinerals += dMinerals
			}
		}
	}
}

func (w *world) initStat() {
	w.countCell = w.MaxX * w.MaxY

	for _, cells := range w.Map {
		for _, cell := range cells {
			if cell.Height <= w.SeaLevel {
				w.underwaterCell++
			}
		}
	}
}

func (w *world) updateStat() {
	w.Year++
	totalMinerals := 0

	for _, cells := range w.Map {
		for _, cell := range cells {
			totalMinerals += cell.localMinerals
		}
	}
	w.TotMinerals = totalMinerals
	w.AVGMinerals = totalMinerals / w.countCell
}

func (w *world) changeSeason(s string) {
	switch s {
	case spring:
		w.Illumination = 20
		w.SeaLevel = 10
	case summer:
		w.Illumination = 40
		w.SeaLevel = 15
	case autumn:
		w.Illumination = 20
		w.SeaLevel = 10
	case winter:
		w.Illumination = 15
		w.SeaLevel = 5
	default:
		w.Illumination = 1
		w.SeaLevel = 1
	}
}
