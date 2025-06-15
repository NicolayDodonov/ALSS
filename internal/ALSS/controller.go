package ALSS

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/logger/baseLogger"
	"context"
	"log"
)

// Controller основная структура пакета и единственная внешне доступная.
// Обеспечивает контроль над внутренней логикой и реализует интерфейс управления и передачи данных.
type Controller struct {
	Parameters Parameters
	Stats      Statistic

	world  *world
	agents *list
	l      logger.Logger
}

// Создаёт новый контроллер модели.
func NewController(conf *config.Config, l *baseLogger.Logger, count, sun, sea, age, energy int) *Controller {
	return &Controller{
		l: l,
		Stats: Statistic{
			AgentStat{0, 0, 0},
			CommandStat{0, 0, 0, 0, 0},
			GenStat{0, 0},
			ResursesStat{0, 0, 0},
			0,
			conf.Stats,
		},
		world: &world{
			MaxX: conf.X,
			MaxY: conf.Y,
			userParameters: userParameters{
				Illumination: sun,
				SeaLevel:     sea,
			},
		},
		Parameters: Parameters{
			typeGenome:          conf.TypeGenome,
			sizeGenome:          conf.SizeGenome,
			maxGen:              conf.MaxGen,
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
	}
}

// Run запускает основной цикл модели, идущий до гибели всех агентов в моделе.
func (c *Controller) Run(frame chan *Frame, ctx context.Context) {
	defer func() {
		log.Printf("Model is shutdown")
	}()

	//Основной цикл
	for {
		select {
		case <-ctx.Done():
			break
		default:
			//модель работает здесь
			c.world.update()

			//для каждого агента вызываем такт жизни
			if err := c.runAgents(); err != nil {
				//Логируем ошибки
				c.l.Error(err.Error())
				//и производим синхронизацию
				if err := c.sync(); err != nil {
					return
				}
			}

			//обновляем статистику
			c.Stats.update(c)

			//отправляем в канал кадр мира
			frame <- c.MakeFrame()

			//проверяем количество живых агентов в мире.
			if c.worldDead() {
				return
			}
		}
	}
}

// InitModel создаёт world, проводит по настройкам пользователя генерацию ландшафта и базовых ресурсов.
// Так же создаёт по настройкам пользователя двусвязный спиок agent
func (c *Controller) InitModel() {
	c.world.init()
	c.makeAgents()
	c.sync()
}
