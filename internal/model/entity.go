package model

import (
	"math/rand"
	"time"
)

const (
	maxGen    = 64
	lengthDNA = 64
	left      = 1
	right     = -1
)
const (
	move = iota
	look
	get
	rotatedLeft
	rotatedRight
	recycling
	reproduction
)

func NewEntity(ID, x, y, longDNA int) *Entity {
	return &Entity{
		ID,
		0,
		100,
		true,
		0,
		Coordinates{
			x,
			y,
		},
		*NewDNA(longDNA),
	}
}

func NewDNA(longDNA int) *DNA {
	var Array []int
	for i := 0; i < longDNA; i++ {
		Array = append(Array, rand.Intn(maxGen))
	}
	return &DNA{
		time.Now().Nanosecond(),
		0,
		Array,
	}
}

func (e *Entity) Run(w *World) (err error) {
	if !e.Live {
		//todo: добавить логгирование
		return nil
	}

	for frameCount := 0; frameCount < 100; {
		switch e.Array[e.Pointer] {
		case move:
			err = e.move(w)
			frameCount += 5
		case look:
			err = e.look(w)
			frameCount += 2
		case get:
			err = e.get(w)
			frameCount += 5
		case rotatedLeft:
			e.rotation(left)
			frameCount++
		case rotatedRight:
			e.rotation(right)
			frameCount++
		case recycling:
			err = e.recycling(w)
			frameCount += 5
		case reproduction:
			err = e.reproduction()
			frameCount += 12
		default:
			e.jump()
			frameCount++
		}

		e.Pointer++
		e.loopPointer()

		if err != nil {
			//todo: добавить логгирование
			return err
		}
	}
	return nil
}

func (e *Entity) move(w *World) error {
	return nil
}

func (e *Entity) rotation(turnCount turns) {
	e.turn += turnCount
	if e.turn > 7 {
		e.turn = 0
	}
	if e.turn < 0 {
		e.turn = 7
	}
}

func (e *Entity) look(w *World) error {
	return nil
}

func (e *Entity) get(w *World) error {
	return nil
}

func (e *Entity) recycling(w *World) error {
	return nil
}

func (e *Entity) reproduction() error {
	return nil
}

func (e *Entity) jump() {
	e.Pointer += (e.Pointer + e.Array[e.Pointer]) % lengthDNA
}

func (e *Entity) loopPointer() {
	e.Pointer = e.Pointer % lengthDNA
}
