package ALSS

type Frame struct {
	Stat *Statistic `json:"stat"`
	Map  *MapJSON   `json:"map"`
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
		Stat: &c.Stats,
		Map: &MapJSON{
			X:        c.world.MaxX,
			Y:        c.world.MaxY,
			SeaLevel: c.world.SeaLevel,
			Cells:    &c.world.Map,
		},
	}
	return frame
}
