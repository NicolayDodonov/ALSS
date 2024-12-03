package alModel

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// newDNA принимает длинну генокода и возвращает структуру генокода с укикальный id.
func newDNA(longDNA int) *DNA {
	var Array []int
	for i := 0; i < longDNA; i++ {
		Array = append(Array, rand.Intn(MaxGen))
	}
	return &DNA{
		time.Now().Nanosecond() + rand.Intn(1000),
		0,
		Array,
	}
}

// set устанавливает все поля DNA в значении input.
func (d *DNA) set(input *DNA) {
	d.ID = input.ID
	d.Pointer = input.Pointer
	for i := 0; i < len(d.Array); i++ {
		d.Array[i] = input.Array[i]
	}
}

// mutation случайно изменяет значение одного гена в DNA.Array.
func (d *DNA) mutation(index int) {
	d.Array[index] = rand.Intn(MaxGen)
	d.ID = time.Now().Nanosecond() + rand.Intn(1000)
}

// jump обеспечивает зацикленный прыжок по DNA.Array.
//
// В качесте аргумента принимает байт в DNA.Array на который
// указывает DNA.Pointer.
func (d *DNA) jump() {
	d.Pointer += (d.Pointer + d.Array[d.Pointer]) % LengthDNA
}

// String создаёт на основе DNA.Array строку содержащую
// информацию об id генокода и его битовую составляющую.
// Реализация стандартного интерфейса Stringer.
func (d *DNA) String() string {
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
