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

	world  *world
	agents *list.List
	l      logger.Logger
}

func NewController(conf config.Config) *Controller {
	return &Controller{}
}

func (c *Controller) Run() {
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
	}
}

// initModel создаёт world, проводит по настройкам пользователя генерацию ландшафта и базовых ресурсов.
// Так же создаёт по настройкам пользователя двусвязный спиок agent
func (c *Controller) initModel() {
	c.makeWorld()
	c.makeAgents()
	c.sync()
}

// Load загружает состояние модели из внешнего источника.
func (c *Controller) Load(data *[]byte) {

}

// Save выгружает состояние модели внешнему потребителю.
func (c *Controller) Save() {
}

// Frame передаёт кадр модели внешнему потребителю. Использует метод io.makeFrame().
func (c *Controller) Frame() *FrameJSON {
	return &FrameJSON{}
}
