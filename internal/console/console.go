package console

import "artificialLifeGo/internal/alModel"

type Console interface {
	Print(*alModel.World)
}
