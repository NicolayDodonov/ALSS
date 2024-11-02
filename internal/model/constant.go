package model

var (
	MaxGen      = 16
	LengthDNA   = 64
	EnergyPoint = 10
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
)
