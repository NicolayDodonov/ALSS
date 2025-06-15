package ALSS

import (
	"io"
	"os"
	"strconv"
)

const path = "logs/stat.log"

type Statistic struct {
	AgentStat   `json:"agent"`
	CommandStat `json:"command"`
	DeathStat   `json:"death"`
	GenStat     `json:"gen"`
}

// обновляется командой Statistic.update()
type AgentStat struct {
	AvgAge    float64 `json:"age"`    //средний возраст агента
	AvgEnergy float64 `json:"energy"` //средняя энергия агента
}

// обновляется Controller.runAgents()
type CommandStat struct {
	Sun   int `json:"photosynthesis"` //получение энергии от солнца
	Hemo  int `json:"chemosynthesis"` //получение энергии от загрязнения
	Mine  int `json:"minersynthesis"` //получение энергии от минералов
	Hunt  int `json:"hunting"`        //получение энергии от охоты
	Other int `json:"other"`          //иное получение энергии (от чего????)
}

// обновляется Controller.runAgents()
type DeathStat struct {
	Age     int `json:"age"`     //количество смертей от возраста
	Hunt    int `json:"hunt"`    //количество смертей от охоты
	OverEn  int `json:"fat"`     //количество смертей от переизбытка энергией
	LowEn   int `json:"skinny"`  //количество смертей от нехватки энергии
	Mineral int `json:"mineral"` //количество смертей от минерального отравления
}

// обновляется командой Statistic.update
type GenStat struct {
	AvgCom float64
	AvgJmp float64
}

// update обновляет средние данные модели
func (s *Statistic) update(c *Controller) {
	// Обнуляем изменяемые параметры
	s.AvgAge = 0
	s.AvgEnergy = 0

	s.AvgCom = 0
	s.AvgJmp = 0

	// перебираем живых агентов
	var count float64 = 0
	for nod := c.agents.root; nod != nil; nod = nod.next {
		// на всякий пожарный проверяем точно ли жив агент
		// на деле не надо, но почему бы не проверить?
		if nod.value.Energy > 0 {
			count++
			// собираем инфу по энергии
			s.AvgAge += float64(nod.value.Energy)
			s.AvgEnergy += float64(nod.value.Energy)
			// собираем инфу по генам
			for _, gen := range nod.value.Genome.Array {
				if gen > maxGenCommand {
					s.AvgJmp++
				} else {
					s.AvgCom++
				}
			}
		}
	}
	//после сбора суммарных данных, определим средние данные
	s.AvgAge = s.AvgAge / count
	s.AvgEnergy = s.AvgEnergy / count

	s.AvgCom = s.AvgCom / count
	s.AvgJmp = s.AvgJmp / count
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

//AvgAge; AvgEnergy; AvgCom; AvgJmp; Sun; Hemo; Mine; Hunt; Other; Age; Hunt; Fat, Skinny; Mineral;

func (s Statistic) String() string {
	return strconv.FormatFloat(s.AgentStat.AvgAge, 'f', 3, 64) + ";" +
		strconv.FormatFloat(s.AgentStat.AvgEnergy, 'f', 3, 64) + "; = ;" +
		strconv.FormatFloat(s.GenStat.AvgCom, 'f', 3, 64) + ";" +
		strconv.FormatFloat(s.GenStat.AvgJmp, 'f', 3, 64) + "; = ;" +
		strconv.Itoa(s.CommandStat.Sun) + ";" +
		strconv.Itoa(s.CommandStat.Hemo) + ";" +
		strconv.Itoa(s.CommandStat.Mine) + ";" +
		strconv.Itoa(s.CommandStat.Hunt) + ";" +
		strconv.Itoa(s.CommandStat.Other) + "; = ;" +
		strconv.Itoa(s.DeathStat.Age) + ";" +
		strconv.Itoa(s.DeathStat.Hunt) + ";" +
		strconv.Itoa(s.DeathStat.OverEn) + ";" +
		strconv.Itoa(s.DeathStat.LowEn) + ";" +
		strconv.Itoa(s.DeathStat.Mineral) + ";" +
		";\n"

}
