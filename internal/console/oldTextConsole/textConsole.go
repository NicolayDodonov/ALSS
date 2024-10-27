package oldTextConsole

import (
	"artificialLifeGo/internal/model"
	"atomicgo.dev/cursor"
	"fmt"
)

type TextConsole struct {
	Alphabet map[string]rune
}

var ASCIIAlphabet = map[string]rune{
	"empty":  ' ',
	"food":   'F',
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
	var canvas = make([][]rune, world.Xsize)

	//заполняем хост
	for x := 0; x < world.Xsize; x++ {
		canvas[x] = make([]rune, world.Ysize+1)
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
				if cell.Entity != nil && cell.Entity.Live {
					canvas[x][y] = tc.Alphabet["entity"]
				}
				canvas[x][y] = tc.Alphabet["empty"]
			case model.FoodCell:
				canvas[x][y] = tc.Alphabet["food"]
			case model.WallCell:
				canvas[x][y] = tc.Alphabet["wall"]
			default:
				canvas[x][y] = tc.Alphabet["nil"]
			}
		}
		//в конец добавляем перенос строки
		canvas[x][world.Ysize] = '\n'
	}
	//добавляем статичтическую информацию
	canvas[world.Xsize] = append(canvas[world.Xsize], []rune(world.GetPrettyStatistic())...)
	//рисуем холст
	for i := 0; i < len(canvas); i++ {
		fmt.Print(string(canvas[i]))
	}
	//вернуть каретку в начало для перерисовки кадра
	//todo: создать свою реализацию движения коре
	cursor.Up(world.Xsize + 1)
}
