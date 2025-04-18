package ALSS

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

type genome struct {
	ID      string
	Pointer int
	Array   []uint8
}

func newRandomGenome(size int) *genome {
	g := &genome{
		ID:      "",
		Pointer: rand.Int() % size,
		Array:   make([]uint8, size),
	}
	for i := range g.Array {
		g.Array[i] = uint8(rand.Int())
	}
	g.hashID()
	return g
}

func newZeroGenome(size int) *genome {
	g := &genome{
		ID:      "",
		Pointer: rand.Int() % size,
		Array:   make([]uint8, size),
	}
	for i := range g.Array {
		g.Array[i] = 0
	}
	g.hashID()
	return g
}

func newBaseGenome() *genome {
	g := &genome{
		ID:      "",
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
	g.hashID()
	return g
}

func equals(g1, g2 *genome) bool {
	if g1.ID == g2.ID {
		return true
	}
	return false
}

func (g *genome) mutation(countMutation int) {
	for i := 0; i < countMutation; i++ {
		g.Array[i] = g.Array[rand.Int()%len(g.Array)]
	}
	g.hashID()
}

func (g *genome) hashID() {
	hash := md5.Sum(g.Array)
	g.ID = hex.EncodeToString(hash[:])
}

func (g *genome) jumpPointer(jumpRange int) {
	g.Pointer = (jumpRange + g.Pointer) % len(g.Array)
}

func (g *genome) getGen() int {
	return int(g.Array[g.Pointer])
}
