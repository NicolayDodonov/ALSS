package ALSS

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger"
	"container/list"
)

// Controller основная структура пакета и единственная внешне доступная.
// Обеспечивает контроль над внутренней логикой и реализует интерфейс управления и передачи данных.
type Controller struct {
	Parameters Parameters
	Statistics Statistics

	world  *world
	agents *list.List
	l      logger.Logger
}

func NewController(conf config.Config) *Controller {
	return &Controller{}
}

func (c *Controller) Run() {
	c.world = newWorld()
	c.makeAgents()

	c.world.init()
	for {
		//model work here
		if err := c.runAgents(); err != nil {
			c.l.Error(err.Error())
			c.sync()
		}

		c.removeDeadAgents()

		//update mStat
		c.updateStat()

		if c.worldDead() {
			break
		}
	}
}

// ResetModel обнуляет состояние мира, списка агентов, всей статистики
// всех геномов и тому подобное...
func (c *Controller) ResetModel() {

}

// ResetWorld обнуляет состояние всех клеток мира, обнуляет мировую статистику
// и иные параметры структуры world.
func (c *Controller) ResetWorld() {

}

// ResetAgents обнуляет состояние всех агентов в модели, очищает их геномы к стандартному
// удаляет мёртвых или иных ошибочных агентов из списка, пересобирает список агентов.
func (c *Controller) ResetAgents() {

}

// LoadModel загружает состояние модели из внешнего источника.
func (c *Controller) LoadModel(data *[]byte) {

}

// SaveModel выгружает состояние модели внешнему потребителю.
func (c *Controller) SaveModel() {
}

// GetFrame передаёт кадр модели внешнему потребителю. Использует метод io.makeFrame().
func (c *Controller) GetFrame() *FrameJSON {
	return &FrameJSON{}
}
