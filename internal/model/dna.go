package model

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func NewDNA(longDNA int) *DNA {
	var Array []int
	for i := 0; i < longDNA; i++ {
		Array = append(Array, rand.Intn(maxGen))
	}
	return &DNA{
		time.Now().Nanosecond() + rand.Intn(1000),
		0,
		Array,
	}
}

func (d *DNA) Set(d2 DNA) {
	*d = d2
}

// GetDNAString создаёт на основе DNA.Array строку содержащую
// информацию об id генокода и его битовую составляющую.
func (d *DNA) GetDNAString() string {
	var s strings.Builder
	s.WriteString("DNA id: " + strconv.Itoa(d.ID) + "\n")
	for i, gen := range d.Array {
		if i%10 == 0 {
			s.WriteString("\n")
		}
		s.WriteString(strconv.Itoa(gen) + " ")
	}
	s.WriteString("\n")
	return s.String()
}

// Mutation случайно изменяет значение одного гена в DNA.Array.
func (d *DNA) Mutation(index int) {
	d.Array[index] += rand.Intn(maxGen)
	d.ID = time.Now().Nanosecond() + rand.Intn(1000)
}
