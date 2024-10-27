package model

const (
	maxGen      = 16
	lengthDNA   = 64
	energyPoint = 10
)
const (
	emptyCell cellType = iota
	foodCell
	wallCell
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
