package ALSS

import (
	"artificialLifeGo/internal/config"
)

// Controller основная структура пакета и единственная внешне доступная.
// Обеспечивает контроль над внутренней логикой и реализует интерфейс управления и передачи данных.
type Controller struct {
	world  *world
	agents *[]*agent //todo: change type array to linked-list
	mStat
	Parameters mParam
}

type mStat struct {
	//todo: add model work info
}

type mParam struct {
	//todo: add model parameters
	baseHP        int
	maxAge        int
	typeGenome    string
	sizeGenome    int
	countMutation int
}

func NewController(conf config.Config) *Controller {
	return &Controller{}
}

func (c *Controller) Run() {
	c.world = newWorld()
	c.agents = makeAgents()

	for {
		//model work here
		c.runAgents()
		//update mStat
		c.updateStat()
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
