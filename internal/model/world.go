package model

func NewWorld(x, y, population int) *World {
	return &World{
		Xsize:       x,
		Ysize:       y,
		Map:         make([][]Cell, population),
		ArrayEntity: make([]Entity, population),
		Statistic: Statistic{
			population,
			0,
			0,
			0,
			0,
		},
	}
}

func (world *World) NewGeneration() error {
	for i := 0; i < world.CountEntity; i++ {
		//создание нового поколения
	}
	return nil
}

func (w *World) SetGeneration() error {
	return nil
}

func (world *World) Clear() error {
	return nil
}

func (world *World) Update() error {
	return nil
}

func (world *World) Execute() error {
	return nil
}

func (world *World) GetCellData(cord Coordinates) (*Cell, error) {
	return nil, nil
}

func (world *World) SetCellFood(cord Coordinates, dFood int) error {
	return nil
}

func (world *World) SetCellPoison(cord Coordinates, dPoison int) error {
	return nil
}

func (world *World) SetCellEntity(cord Coordinates, entity *Entity) error {
	return nil
}

func (world *World) GetStatistic() (string, error) {
	return " ", nil
}
