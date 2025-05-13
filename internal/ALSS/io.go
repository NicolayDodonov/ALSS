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
			Cells:    &c.world.Map,
		},
	}
	return frame
}
