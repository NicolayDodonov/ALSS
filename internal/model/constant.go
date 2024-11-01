package model

const (
	maxGen      = 16
	lengthDNA   = 64
	energyPoint = 10
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

// Poison level
const (
	pLevel1 = 5
	pLevel2 = 25
	pLevel3 = 50
	pLevel4 = 75
)
