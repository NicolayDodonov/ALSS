package ALSS

func (c *Controller) runAgents() {
	for _, a := range *c.agents {
		if err := a.run(c); err != nil {
			//todo: save error in log
		}
	}
}

func makeID(s string) string {
	switch s {
	case typeOfGenome:
		return ""
	case typeOfAgent:
		return ""
	default:
		return ""
	}
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

func (c *Controller) updateStat() {

}

func makeAgents() *[]*agent {
	//todo: make agents linked-list former
	return nil
}
