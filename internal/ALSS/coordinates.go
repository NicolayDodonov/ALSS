package ALSS

// coordinates структура, реализующая логику сложения и вычитания координат,
// изменения угла и связи координаты и угла
type coordinates struct {
	X int
	Y int
}

type angle uint8

// sum создаёт новые координаты как сумму a + b
func sum(a, b *coordinates) *coordinates {
	return &coordinates{X: a.X + b.X, Y: a.Y + b.Y}
}

// del создаёт новые координаты как разность a + b
func del(a, b *coordinates) *coordinates {
	return &coordinates{X: a.X - b.X, Y: a.Y - b.Y}
}

func check(c *coordinates, MaxX, MaxY int) bool {
	if (c.X >= 0 && c.X < MaxX) &&
		(c.Y >= 0 && c.Y < MaxY) {
		return true
	}
	return false
}

// offset создаёт новые координаты как смещение по углу a от c
func offset(c *coordinates, a angle) *coordinates {
	newC := coordinates{c.X, c.Y}
	/*--------y
	|	701
	|	6*2
	|	543
	x
	*/
	switch a {
	case 0:
		newC.X--
	case 1:
		newC.X--
		newC.Y++
	case 2:
		newC.Y++
	case 3:
		newC.X++
		newC.Y--
	case 4:
		newC.X++
	case 5:
		newC.X++
		newC.Y--
	case 6:
		newC.Y--
	case 7:
		newC.X--
		newC.Y--
	}
	return &newC
}

func (a angle) plus() {
	a++
	if a > 7 {
		a = 0
	}
}

func (a angle) minus() {
	a--
	if a > 7 {
		a = 7
	}
}
