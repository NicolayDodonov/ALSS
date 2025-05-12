package ALSS

import (
	"container/list"
	"fmt"
	"reflect"
)

func makeID(i interface{}) {
	v := reflect.ValueOf(i).Elem()
	switch v.Type().Name() {
	case "agent":
		v.Field(0).SetString("A")
	case "genome":
		v.Field(0).SetString("S")
	}
}

// makeAgents формирует массив начальных агентов модели.
func (c *Controller) makeAgents() {
	agents := list.New()
	for i := 0; i < c.Parameters.startPopulation; i++ {
		agents.PushBack(newAgent(c))
	}
	c.agents = agents
}

// makeWorld создаёт мир, генерирует карту с начальными условиями от пользователя.
func (c *Controller) makeWorld() {
	w := newWorld(c.Parameters.X, c.Parameters.Y)
	w.initMap()
	c.world = w
}

// sync - синхронизация агентов и мира.
//
// Исправление списка агентов (удаление мёртвых не удалённых агентов).
func (c *Controller) sync() {
	c.l.Info("start synchronization model")

	c.removeDeadAgents()

	//удаляем все ссылки живых, мёртвых и ошибочных агентов из мира
	for _, cells := range c.world.Map {
		for _, cell := range cells {
			cell.Agent = nil
		}
	}

	for a := c.agents.Front(); a != nil; a = a.Next() {
		cell, _ := c.world.getCell(&a.Value.(*agent).coordinates)
		cell.Agent = a.Value.(*agent)
	}

}

// runAgents проходится по списку каждого живого агента и запускает
func (c *Controller) runAgents() error {
	for element := c.agents.Front(); element != nil; element = element.Next() {
		if err := element.Value.(*agent).run(c, element); err != nil {
			// проверяем, распологается ли агент в своей клетке, или случился рассинхрон.
			if element.Value.(*agent).checkSelfPosition(c) {
				return fmt.Errorf("agent ID:%s position check failed", element.Value.(*agent).ID)
			}

			c.l.Error(err.Error())
		}
	}
	return nil
}

// removeDeadAgents проходит по всему списку агентов и очищает его от мёртвых агентов
func (c *Controller) removeDeadAgents() {
	for a := c.agents.Front(); a != nil; a = a.Next() {
		if a.Value.(*agent).Energy <= 0 {
			//удаляем ссылку на агента в клетке
			cell, _ := c.world.getCell(&a.Value.(*agent).coordinates)
			cell.Agent = nil
			c.agents.Remove(a) //todo: Возможное место появление бага!!!
		}
	}
}

// worldDead проверяет весь список агентов модели на жизнеспособность.
// если таких нет, если все агенты умертвы - возвращает True.
func (c *Controller) worldDead() bool {
	for a := c.agents.Front(); a != nil; a = a.Next() {
		if a.Value.(*agent).Energy > 0 {
			return false
		}
	}
	return true
}
