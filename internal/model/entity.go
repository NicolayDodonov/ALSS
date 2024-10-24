package model

import (
	"math/rand"
	"time"
)

const (
	maxGen = 64
)

func NewEntity(ID, longDNA int) *Entity {
	return &Entity{
		ID,
		0,
		100,
		0,
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

func (e *Entity) Run() error {
	if !e.IsLive {

	}
}
