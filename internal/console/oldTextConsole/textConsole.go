package oldTextConsole

import (
	"artificialLifeGo/internal/model"
	"atomicgo.dev/cursor"
	"fmt"
	"time"
)

type TextConsole struct {
	Alphabet map[string]byte
}

var ASCIIAlphabet = map[string]byte{
	"empty":  ' ',
	"food":   '+',
	"wall":   '#',
	"entity": '0',
	"nil":    '?',
}

func New() *TextConsole {
	return &TextConsole{
		ASCIIAlphabet,
	}
}

// Print выводит на экран кадр мира + статистическую информацию
func (tc *TextConsole) Print(world *model.World) {
	//создаём холст
	var canvas = make([][]byte, world.Xsize)
	//заполняем хост
	for x := 0; x < world.Xsize; x++ {
		canvas[x] = make([]byte, world.Ysize)
		//заполняем строку холста
		for y := 0; y < world.Ysize; y++ {
			//получаем клетку мира
			cell, err := world.GetCellData(model.Coordinates{X: x, Y: y})
			if err != nil {
				//если почему то не можем получить - пропускаем её
				continue
			}
			//смотрим что в ней и соотвественно доавляем на холст

			switch cell.Types {
			case model.EmptyCell:
				canvas[x][y] = tc.Alphabet["empty"]
			case model.FoodCell:
				canvas[x][y] = tc.Alphabet["food"]
			case model.WallCell:
				canvas[x][y] = tc.Alphabet["wall"]
			default:
				canvas[x][y] = tc.Alphabet["nil"]
			}
			if cell.Entity != nil {
				canvas[x][y] = '1'
			}
		}
		//в конец добавляем перенос строки
	}

	//рисуем холст
	for i := 0; i < len(canvas); i++ {
		fmt.Print("|" + string(canvas[i]) + "|\n")
	}
	fmt.Print(world.GetPrettyStatistic() + "\n")
	//вернуть каретку в начало для перерисовки кадра
	//todo: создать свою реализацию движения коретки
	cursor.Up(world.Xsize + 5)
	time.Sleep(100 * time.Millisecond)
}
