package ALSS

import (
	"fmt"
)

func newID() string {
	//todo: сделать id генератор
	return "random id"
}

// makeAgents формирует массив начальных агентов модели.
func (c *Controller) makeAgents() {
	agents := newList()
	for i := 0; i < c.Parameters.startPopulation; i++ {
		agents.add(newAgent(c))
	}
	c.agents = agents
}

// sync - синхронизация агентов и мира.
//
// Исправление списка агентов (удаление мёртвых не удалённых агентов).
func (c *Controller) sync() error {
	c.l.Info("start synchronization model")

	//критическое место очистки агентов. Если возникает агент вне списка - завершить модель
	if err := c.removeDeadAgents(); err != nil {
		return err
	}

	//удаляем все ссылки живых, мёртвых и ошибочных агентов из мира
	for _, cells := range c.world.Map {
		for _, cell := range cells {
			cell.Agent = nil
		}
	}

	//записываем всех агентов из списка повторно.
	for nod := c.agents.root; nod != nil; nod = nod.next {
		cell, _ := c.world.getCell(&nod.value.coordinates)
		cell.Agent = nod.value
	}
	return nil
}

// runAgents проходится по списку каждого живого агента и запускает
func (c *Controller) runAgents() error {
	for nod := c.agents.root; nod != nil; nod = nod.next {
		if err := nod.value.run(c); err != nil {
			if nod.value.checkSelfPosition(c) {
				return fmt.Errorf("agent ID:%s position check failed", nod.value.ID)
			}
			c.l.Error(err.Error())
		}

	}
	return nil
}

// removeDeadAgents проходит по всему списку агентов и очищает его от мёртвых агентов.
// может вызвать ошибку. В таком случае лучше всего завершить работу модели.
func (c *Controller) removeDeadAgents() error {
	for nod := c.agents.root; nod != nil; nod = nod.next {
		if nod.value.Energy <= 0 {
			cell, _ := c.world.getCell(&nod.value.coordinates)
			cell.Agent = nil
			if err := c.agents.del(nod.value); err != nil {
				//критическое место. Если тут возникает ошибка, можно прекращать симуляцию
				return err
			}
		}
	}
	return nil
}

// worldDead проверяет весь список агентов модели на жизнеспособность.
// если таких нет, если все агенты умертвы - возвращает True.
func (c *Controller) worldDead() bool {
	for nod := c.agents.root; nod != nil; nod = nod.next {
		if nod.value.Energy > 0 {
			return false
		}
	}
	return true
}
