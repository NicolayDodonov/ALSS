package simulation

import (
	"artificialLifeGo/internal/storage"
)

// Simulation структура описывающая симуляцию
type Simulation struct {
	storage.Storage
}

func New(storage storage.Storage) (s *Simulation) {
	return &Simulation{
		storage,
	}
}
