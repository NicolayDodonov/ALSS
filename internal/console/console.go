package console

import "artificialLifeGo/internal/model"

type Console interface {
	Print(*model.World)
}
