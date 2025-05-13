package ALSS

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/logger/baseLogger"
	"container/list"
	"context"
)

// Controller основная структура пакета и единственная внешне доступная.
// Обеспечивает контроль над внутренней логикой и реализует интерфейс управления и передачи данных.
type Controller struct {
	Parameters Parameters

	world  *world
	agents *list.List
	l      logger.Logger
}

func NewController(conf *config.Config, l *baseLogger.Logger, count, sun, sea, age, energy int) *Controller {
	return &Controller{
		l: l,
		Parameters: Parameters{
			WorldParam{
				X:            conf.X,
				Y:            conf.Y,
				Illumination: sun,
				SeaLevel:     sea,
			},
			AgentParam{
				typeGenome:          conf.TypeGenome,
				sizeGenome:          conf.SizeGenome,
				startPopulation:     count,
				baseAgentEnergy:     conf.BaseEnergy,
				maxAgentAge:         age,
				maxAgentEnergy:      energy,
				energyCost:          conf.ActionCost,
				attackProfitPercent: conf.HuntingCoefficient,
				madePollution:       conf.PollutionCost,
				minEnergyToBirth:    conf.BirthCost,
				countMutation:       conf.MutationCount,
			},
		},
	}
}

func (c *Controller) Run(frame chan *Frame, ctx context.Context) {
	for {
		//model work here
		if err := c.runAgents(); err != nil {
			c.l.Error(err.Error())
			c.sync()
		}

		c.removeDeadAgents()

		//update mStat
		c.world.updateStat()

		if c.worldDead() {
			break
		}

		frame <- c.MakeFrame()
	}
}

// InitModel создаёт world, проводит по настройкам пользователя генерацию ландшафта и базовых ресурсов.
// Так же создаёт по настройкам пользователя двусвязный спиок agent
func (c *Controller) InitModel() {
	c.makeWorld()
	c.makeAgents()
	c.sync()
}
