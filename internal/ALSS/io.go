package ALSS

type Frame struct {
	Stat StatisticJSON `json:"stat"`
	Map  MapJSON       `json:"map"`
}

type StatisticJSON struct {
	Year      int `json:"world_year"`
	Pollution int `json:"world_pollution"`
	TotalM    int `json:"total_minerals"`
	AvgM      int `json:"avg_minerals"`
	Live      int `json:"live_agent"`
	Death     int `json:"death_agent"`
}

type MapJSON struct {
	X        int  `json:"x_size"`
	Y        int  `json:"y_size"`
	SeaLevel int  `json:"sea_level"`
	Cells    *Map `json:"cells"`
}

type CellJSON struct {
	Height  int       `json:"height"`
	Mineral int       `json:"mineral"`
	Agent   AgentJSON `json:"agent"`
}

type AgentJSON struct {
	ID     string `json:"id"`
	Energy int    `json:"energy"`
	DNA    genome `json:"dna"`
}

type Message struct {
	Count  int `json:"count"`
	Sun    int `json:"sun"`
	Sea    int `json:"sea"`
	Age    int `json:"age"`
	Energy int `json:"energy"`
}

func (c *Controller) MakeFrame() *Frame {
	frame := &Frame{
		Stat: StatisticJSON{
			Year:      c.world.Year,
			Pollution: c.world.Pollution,
			TotalM:    c.world.TotMinerals,
			AvgM:      c.world.AvgMinerals,
			Live:      c.world.LiveAgent,
			Death:     c.world.DeathAgent,
		},
		Map: MapJSON{
			X:        c.world.MaxX,
			Y:        c.world.MaxY,
			SeaLevel: c.world.SeaLevel,
			Cells:    &c.world.Map,
		},
	}
	return frame
}
