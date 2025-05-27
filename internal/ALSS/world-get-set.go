package ALSS

import "fmt"

func (w *world) addLocalMinerals(c *coordinates, d int) error {
	if c.X >= 0 && c.X < w.MaxX &&
		c.Y >= 0 && c.Y < w.MaxY {
		w.Map[c.X][c.Y].LocalMinerals += d
		return nil
	}
	return fmt.Errorf("cant set minerals: out of bound!")
}

func (w *world) getCell(c *coordinates) *cell {
	if c.X >= 0 && c.X < w.MaxX &&
		c.Y >= 0 && c.Y < w.MaxY {
		return &w.Map[c.X][c.Y]
	}
	return nil
}
