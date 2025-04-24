package ALSS

type Parameters struct {
	ControllerParam
	WorldParam
	AgentParam
}

type WorldParam struct {
	X, Y              int
	baseSunCost       int
	baseMineralCost   int
	baseGrassCost     int
	basePollutionPart int
}

type AgentParam struct {
	typeGenome      string
	sizeGenome      int
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

type ControllerParam struct {
}
