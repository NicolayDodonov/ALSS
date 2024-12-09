package alModel

// sum складывает коордианты a и b, возвращая результа
func sum(a, b Coordinates) Coordinates {
	SumCord := Coordinates{
		a.X + b.X,
		a.Y + b.Y,
	}
	return SumCord
}

// del вычитает из координаты a - b, возвращая результат
func del(a, b Coordinates) Coordinates {
	DelCord := Coordinates{
		a.X - b.X,
		a.Y - b.Y,
	}
	return DelCord
}

// viewCell принимает угол взора и возвращает координату взгляда.
// Взгляд расчитывается по следующей схеме:
//
// 123
//
// 0*4
//
// 765
//
// Где * - бот, смотрящий в указанную сторону.
func viewCell(turn turns) Coordinates {
	cordTurn := Coordinates{
		0,
		0,
	}
	switch turn {
	case 0:
		cordTurn.Y--
	case 1:
		cordTurn.X++
		cordTurn.Y--
	case 2:
		cordTurn.X++
	case 3:
		cordTurn.X++
		cordTurn.Y++
	case 4:
		cordTurn.Y++
	case 5:
		cordTurn.X--
		cordTurn.Y++
	case 6:
		cordTurn.X--
	case 7:
		cordTurn.X--
		cordTurn.Y--
	}
	return cordTurn
}

// checkLimit проверяет, входит ли значение в лимиты [0,limit).
// Если входит, возвращает true, иначе false.
func checkLimit(value, limit Coordinates) bool {
	if value.X >= 0 &&
		value.Y >= 0 &&
		value.X < limit.X &&
		value.Y < limit.Y {
		return true
	} else {
		return false
	}
}
