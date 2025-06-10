package ALSS

import (
	"io"
	"os"
	"strconv"
)

const path = "logs/stat.log"

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
func (s *Statistic) update(c *Controller) {
	s.Year = c.world.Year
	s.Poison = c.world.Pollution

	//Resources
	s.AvgMineral = 0
	for _, cells := range c.world.Map {
		for _, cell := range cells {
			s.AvgMineral += cell.LocalMinerals
		}
	}
	s.AvgMineral /= c.world.CountCell
	//Command and energy
	s.AvgEnergy = 0
	s.AvgCommand = 0
	s.AvgJump = 0
	for nod := c.agents.root; nod != nil; nod = nod.next {
		s.AvgEnergy = nod.value.Energy
		for _, gen := range nod.value.Genome.Array {
			if gen > maxGenCommand {
				s.AvgJump++
			} else {
				s.AvgCommand++
			}
		}
	}
	if c.agents.len > 0 {
		s.AvgEnergy /= c.agents.len
		s.AvgCommand /= c.agents.len
		s.AvgJump /= c.agents.len
		s.CountAgent = c.agents.len
	}
}

func (s *Statistic) save() error {
	//открыть/создать новый файл
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//добавить новый слой
	_, err = io.WriteString(file, s.String())
	if err != nil {
		return err
	}

	return nil
}

//year; avgMin; poison; avgCom; avgJump; countAgent; avgEn;

func (s Statistic) String() string {
	return strconv.Itoa(s.Year) + "; " +
		strconv.Itoa(s.AvgMineral) + "; " +
		strconv.Itoa(s.Poison) + "; " +
		strconv.Itoa(s.AvgCommand) + "; " +
		strconv.Itoa(s.AvgJump) + "; " +
		strconv.Itoa(s.CountAgent) + "; " +
		strconv.Itoa(s.AvgEnergy) + ";\n"
}
