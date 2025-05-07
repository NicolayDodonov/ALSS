package ALSS

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
			Cells:    c.mapToJSON(),
		},
	}
	return frame
}

func (c *Controller) mapToJSON() *[][]CellJSON {
	Map := make([][]CellJSON, c.world.MaxY)
	for y, cells := range c.world.Map {
		Map = append(Map, make([]CellJSON, c.world.MaxX))
		for x, cell := range cells {
			Map[y][x] = CellJSON{
				Height:  cell.Height,
				Mineral: cell.localMinerals,
				Agent:   AgentJSON{},
			}
		}
	}
	return &Map
}
