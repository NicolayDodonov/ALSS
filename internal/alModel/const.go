package alModel

// глобальные константы модели
var (
	MaxGen         = 1     // Максимальное значение числового гена
	MaxFoodPercent = 1     // Процент содержания в мире еды
	LengthDNA      = 1     // Длинна генокода
	EnergyPoint    = 1     // Базовый шаг увеличения энергии
	TypeBrain      = ""    // Тип мозга бота используемый в моделе
	LoopX          = false // зацикленность мира по X координате
	LoopY          = false // зацикленность мира по Y координате
)

const (
	EmptyCell CellTypes = iota // Клетка пуста
	FoodCell                   // Клетка с едой
	WallCell                   // Клетка со стеной
)

// run() константы

const (
	move = iota
	look
	get
	rotatedLeft
	rotatedRight
	recycling
	reproduction

	left  = 1
	right = -1

	maxFC    = 10 // Максимально возможный frame count
	minFC    = 1  // Минимально возможный frame count
	middleFC = 5  // Средний frame count
	smallFC  = 2  // Малый frame count
)

// Уровни отравления
const (
	pLevel1   = 1   // Минимальный уровень отравления
	pLevel2   = 5   // Малый уровень отравления
	pLevel3   = 25  // Средний уровень, ощутимый
	pLevel4   = 50  // Опасный уровень отравления
	pLevelDed = 75  // Смертельный уровень отравления
	PLevelMax = 100 // Максимальный уровень отравления
)
