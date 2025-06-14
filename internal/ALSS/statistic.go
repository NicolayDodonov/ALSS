package ALSS

import (
	"io"
	"os"
	"strconv"
)

const path = "logs/stat.log"

type Statistic struct {
	AgentStat
	CommandStat
	DeathStat
	GenStat
	year int
}
type AgentStat struct {
	AvgAge    float64
	AvgEnergy float64
}
type CommandStat struct {
	Sun   int
	Hemo  int
	Mine  int
	Hunt  int
	Other int
}
type DeathStat struct {
	Age    int
	Hunt   int
	Energy int
}
type GenStat struct {
	AvgCom float64
	AvgJmp float64
}

// update проверяет ряд параметров модели и сохраняет их в себе.
func (s *Statistic) update(c *Controller) {

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
	return strconv.FormatFloat(s.AgentStat.AvgAge, 'f', 3, 64) + ";" +
		strconv.FormatFloat(s.AgentStat.AvgEnergy, 'f', 3, 64) + ";" +
		strconv.Itoa(s.CommandStat.Sun) + ";" +
		strconv.Itoa(s.CommandStat.Hemo) + ";" +
		strconv.Itoa(s.CommandStat.Mine) + ";" +
		strconv.Itoa(s.CommandStat.Hunt) + ";" +
		strconv.Itoa(s.CommandStat.Other) + ";" +
		strconv.Itoa(s.DeathStat.Age) + ";" +
		strconv.Itoa(s.DeathStat.Hunt) + ";" +
		strconv.Itoa(s.DeathStat.Energy) + ";" +
		strconv.FormatFloat(s.GenStat.AvgCom, 'f', 3, 64) + ";" +
		strconv.FormatFloat(s.GenStat.AvgJmp, 'f', 3, 64)

}
