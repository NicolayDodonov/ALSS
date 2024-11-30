package model

var (
	MaxGen      = 16
	LengthDNA   = 64
	EnergyPoint = 10
	TypeBrain   = "brain16"
	LoopX       = false
	LoopY       = false
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
	pLevel1   = 5
	pLevel2   = 25
	pLevel3   = 50
	pLevel4   = 75
	pLevelDed = 100
)
