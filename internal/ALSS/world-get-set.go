package ALSS

import "fmt"

func (w *world) addMinerals(c *coordinates, d int) {
	if check(c, w.MaxX, w.MaxY) {
		cell, err := w.getCell(c)
		if err != nil {
			return
		}
		cell.LocalMinerals = cell.LocalMinerals + d
		if cell.LocalMinerals > maxMineral {
			cell.LocalMinerals = maxMineral
		}
		return
	}
}

func (w *world) addPollution(d int) {
	w.Pollution = (w.Pollution + d) % maxPollution
}

func (w *world) getCell(c *coordinates) (*cell, error) {
	if check(c, w.MaxX, w.MaxY) {
		return &w.Map[c.X][c.Y], nil
	}
	return nil, fmt.Errorf("cant get cell: out of bound!")
}
