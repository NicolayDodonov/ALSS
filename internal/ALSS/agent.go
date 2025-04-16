package ALSS

type agent struct {
	angle  angle
	genome dna
}

type dna struct {
	pointer int
	array   []uint8
}

func (a agent) run() error {
	//todo: live?
	//todo: run dna
	//todo: pollution
	//todo: dead control
	return nil
}
