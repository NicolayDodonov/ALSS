package ALSS

import "reflect"

func makeID(i interface{}) {
	v := reflect.ValueOf(i).Elem()
	switch v.Type().Name() {
	case "agent":
		v.Field(0).SetString("A")
	case "genome":
		v.Field(0).SetString("S")
	}
}

func makeAgents() *[]*agent {
	//todo: make agents linked-list former
	return nil
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

func (c *Controller) runAgents() {
	for _, a := range *c.agents {
		if err := a.run(c); err != nil {
			c.l.Error(err.Error())
		}
	}
}

func (c *Controller) removeDeadAgents() {

}

func (c *Controller) updateStat() {

}

func (c *Controller) worldDead() bool {
	return true
}
