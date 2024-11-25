package model

type brain interface {
	run(*Entity, *World) error
}

type brain16 struct {
}

func (brain16) run(e *Entity, w *World) error {
	return nil
}

type brain64 struct {
}

func (brain64) run(e *Entity, w *World) error {
	return nil
}
