package model

func Sum(a, b Coordinates) Coordinates {
	SumCord := Coordinates{
		a.X + b.X,
		a.Y + b.Y,
	}
	return SumCord
}

func Del(a, b Coordinates) Coordinates {
	DelCord := Coordinates{
		a.X - b.X,
		a.Y - b.Y,
	}
	return DelCord
}

func makeTurn(turn turns) Coordinates {
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
