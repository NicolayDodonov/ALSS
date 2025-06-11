package ALSS

// Основные настройки модели.
type Parameters struct {
	typeGenome      string
	sizeGenome      int
	maxGen          int
	startPopulation int
	baseAgentEnergy int
	maxAgentAge     int
	maxAgentEnergy  int

	energyCost          int
	attackProfitPercent int
	madePollution       int

	minEnergyToBirth int
	countMutation    int
}
