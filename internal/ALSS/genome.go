package ALSS

import (
	"math/rand"
)

type genome struct {
	ID      string
	Pointer int
	Array   []uint8
}

func newRandomGenome(size int) *genome {
	g := &genome{
		Pointer: rand.Int() % size,
		Array:   make([]uint8, size),
	}
	for i := range g.Array {
		g.Array[i] = uint8(rand.Int())
	}
	makeID(g)
	return g
}

func newZeroGenome(size int) *genome {
	g := &genome{
		Pointer: rand.Int() % size,
		Array:   make([]uint8, size),
	}
	for i := range g.Array {
		g.Array[i] = 0
	}
	makeID(g)
	return g
}

func newBaseGenome() *genome {
	g := &genome{
		Pointer: 0,
		Array: []uint8{
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
			25, 25, 25, 25, 25, 25, 25, 25,
		},
	}
	makeID(g)
	return g
}

func equals(g1, g2 *genome) bool {
	difference := 0
	for i := 0; i < len(g1.Array); i++ {
		if g1.Array[i] != g2.Array[i] {
			difference++
		}
		if difference > 1 {
			return false
		}
	}
	return true
}

func (g *genome) mutation(countMutation int) {
	for i := 0; i < countMutation; i++ {
		g.Array[i] = g.Array[rand.Int()%len(g.Array)]
	}
	makeID(g)
}

func (g *genome) jumpPointer(jumpRange int) {
	g.Pointer = (jumpRange + g.Pointer) % len(g.Array)
}

func (g *genome) getGen() int {
	return int(g.Array[g.Pointer])
}

func (g *genome) getGenIndex(i int) int {
	return int(g.Array[i%len(g.Array)])
}
