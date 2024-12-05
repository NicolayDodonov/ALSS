package alModel

// глобальные константы модели
var (
	MaxGen          int // Максимальное значение числового гена
	MaxFoodPercent  int // Процент содержания в мире еды
	LengthDNA       int // Длинна генокода
	EnergyPoint     int // Базовый шаг увеличения энергии
	BasePoisonLevel int

	TypeBrain    string // Тип мозга бота используемый в моделе
	PoisonEnable bool
	LoopX        bool // зацикленность мира по X координате
	LoopY        bool // зацикленность мира по Y координате
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
