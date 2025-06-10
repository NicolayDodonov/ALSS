package ALSS

// world структура описывающая структуру мира модели
// и реализующий основные методы работы с миром модели
type world struct {
	Map         Map
	MaxX        int
	MaxY        int
	CountCell   int
	CountUWCell int
	dynamicParameters
	userParameters
}

type dynamicParameters struct {
	Year         int
	Pollution    int
	PollutionFix int
}

type userParameters struct {
	Illumination int
	SeaLevel     int
}

type cell struct {
	Agent         *agent `json:"agent"`
	Height        int    `json:"height"`
	LocalMinerals int    `json:"mineral"`
}

type Map [][]cell

func (w *world) update() {
	w.Year++
	//Растворить часть яда в клетках ниже уровня моря
	solution := w.Pollution / (w.CountUWCell * 100)
	//Обновить компенсацию яда
	for _, cells := range w.Map {
		for _, cell := range cells {
			if cell.Height >= w.CountUWCell {
				continue
			}
			cell.LocalMinerals = (cell.LocalMinerals + solution) % 255
		}
	}
}

// init создаёт пустую карту мира, генерирует высотный ланшафт, настраивает уровень моря
// и проводит начальную настройку статистики.
func (w *world) init() {
	w.dynamicParameters = dynamicParameters{
		0, 0, 0,
	}
	//создать карту и заполнить её клетками
	w.newMap()
	//вызвать генератор высотности карты
	w.landscapeGenerator(mapGRADIENT)
	//подсчитать колличество клеток под водой
	w.countingUWCell()
}

func (w *world) newMap() {
	w.CountCell = w.MaxX * w.MaxY

	newMap := make([][]cell, w.MaxY)
	for y := range newMap {
		newMap[y] = make([]cell, w.MaxX)
		for x := range newMap[y] {
			newMap[y][x] = cell{
				Agent:         nil,
				Height:        0,
				LocalMinerals: 0,
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
			gradientLayer = w.MaxY / maxHeight
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

func (w *world) countingUWCell() {
	for _, cells := range w.Map {
		for _, c := range cells {
			if c.Height <= w.SeaLevel {
				w.CountUWCell++
			}
		}
	}
}
