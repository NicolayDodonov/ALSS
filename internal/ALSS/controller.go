package ALSS

import "artificialLifeGo/internal/config"

// Controller основная структура пакета и единственная внешне доступная.
// Обеспечивает контроль над внутренней логикой и реализует интерфейс управления и передачи данных.
type Controller struct {
	world  world
	agents []agent //todo: change type array to linked-list
	mStat
}

type mStat struct {
	//todo: add model work info
}

func NewController(conf config.Config) *Controller {
	return &Controller{}
}

func (c *Controller) Run() {

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
func (c *Controller) SaveModel() *WorldJson {
	return &WorldJson{}
}

// sync - синхронизация агентов и мира.
//
// Исправление списка агентов (удаление мёртвых не удалённых агентов).
func (c *Controller) sync() {
	//удаляем все ссылки живых, мёртвых и ошибочных агентов из мира
	for _, cells := range c.world.Map {
		for _, cell := range cells {
			cell.Agent = nil
		}
	}
	//удаление мёртвых агентов их списка
	//todo: реализовать при создании linked-list
}
