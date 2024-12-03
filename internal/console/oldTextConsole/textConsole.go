package oldTextConsole

import (
	"artificialLifeGo/internal/alModel"
	"atomicgo.dev/cursor"
	"fmt"
	"time"
)

type TextConsole struct {
	Alphabet map[string]rune
}

var ASCIIAlphabet = map[string]rune{
	"empty":  ' ',
	"food":   '▞',
	"wall":   '▓',
	"entity": '0',
	"nil":    '?',
	"poison": '░',
}

func New() *TextConsole {
	return &TextConsole{
		ASCIIAlphabet,
	}
}

// Print выводит на экран кадр мира + статистическую информацию
func (tc *TextConsole) Print(world *alModel.World) {
	//создаём холст
	var canvas = make([][]rune, world.Xsize)
	//заполняем хост
	for x := 0; x < world.Xsize; x++ {
		canvas[x] = make([]rune, world.Ysize)
		//заполняем строку холста
		for y := 0; y < world.Ysize; y++ {
			//получаем клетку мира
			cell, err := world.GetCellData(alModel.Coordinates{X: x, Y: y})
			if err != nil {
				//если почему то не можем получить - пропускаем её
				continue
			}
			//смотрим что в ней и соотвественно доавляем на холст

			switch cell.Types {
			case alModel.EmptyCell:
				canvas[x][y] = tc.Alphabet["empty"]
			case alModel.FoodCell:
				canvas[x][y] = tc.Alphabet["food"]
			case alModel.WallCell:
				canvas[x][y] = tc.Alphabet["wall"]
			default:
				canvas[x][y] = tc.Alphabet["nil"]
			}
			if cell.Poison > alModel.PLevelMax/2 {
				canvas[x][y] = tc.Alphabet["poison"]
			}
			//if cell.Entity != nil {
			//	canvas[x][y] = tc.Alphabet["entity"]
			//}
			if cell.Entity != nil {
				canvas[x][y] = tc.Alphabet["entity"]
			}
		}
		//в конец добавляем перенос строки
	}

	//рисуем холст
	for i := 0; i < len(canvas); i++ {
		fmt.Print("▓" + string(canvas[i]) + "▓    \n")
	}
	fmt.Print(world.GetPrettyStatistic() + "\n")
	//вернуть каретку в начало для перерисовки кадра
	//todo: создать свою реализацию движения коретки
	cursor.Up(world.Xsize + 5)
	time.Sleep(100 * time.Millisecond)
}
