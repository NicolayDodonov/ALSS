package application

import "artificialLifeGo/internal/logger"

type Application struct {
	l logger.Logger
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run() {
	//основное тело программы
	//todo: server.Start()
}
