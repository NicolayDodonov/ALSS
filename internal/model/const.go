package model

var (
	MaxGen         = 16
	MaxFoodPercent = 50
	LengthDNA      = 64
	EnergyPoint    = 10
	TypeBrain      = "brain16"
	LoopX          = false
	LoopY          = false
)

const (
	EmptyCell CellTypes = iota
	FoodCell
	WallCell
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

	maxFC    = 10
	middleFC = 5
	smallFC  = 2
	minFC    = 1
)

// Poison level
const (
	pLevel1   = 1
	pLevel2   = 5
	pLevel3   = 25
	pLevel4   = 50
	pLevelDed = 75
	PLevelMax = 100
)
