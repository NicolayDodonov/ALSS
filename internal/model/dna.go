package model

import (
	"math/rand"
	"time"
)

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

func (d *DNA) Set(d2 DNA) {
	*d = d2
}

// Mutation случайно изменяет значение одного гена в DNA.Array.
func (d *DNA) Mutation(index int) {
	d.Array[index] += rand.Intn(maxGen)
}
