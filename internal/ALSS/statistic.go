package ALSS

type Statistic struct {
	Resources
	Command
	Life
	Year int
}
type Resources struct {
	AvgMineral int
	Poison     int
}

type Command struct {
	AvgCommand int
	AvgJump    int
}

type Life struct {
	AvgEnergy  int
	CountAgent int
}

// update проверяет ряд параметров модели и сохраняет их в себе.
func (s Statistic) update(c Controller) {
	//Resources
	s.AvgMineral = 0
	for _, cells := range c.world.Map {
		for _, cell := range cells {
			s.AvgMineral += cell.LocalMinerals
		}
	}
	s.AvgMineral /= c.world.countCell
	//Command and energy
	s.AvgEnergy = 0
	s.AvgCommand = 0
	s.AvgJump = 0
	for nod := c.agents.root; nod != nil; nod = nod.next {
		for _, gen := range nod.value.Genome.Array {
			s.AvgEnergy = 0
			if gen > maxGenCommand {
				s.AvgJump++
			} else {
				s.AvgJump++
			}

		}
	}
	s.AvgEnergy /= c.agents.len
	s.AvgCommand /= c.agents.len
	s.AvgJump /= c.agents.len
	s.CountAgent = c.agents.len
}

func (s Statistic) save() {

}
