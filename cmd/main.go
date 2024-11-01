package main

import (
	oTC "artificialLifeGo/internal/console/oldTextConsole"
	"artificialLifeGo/internal/logger"
	bL "artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/simulation"
)

func main() {

	MustLogger()

	console := oTC.New()
	sim := simulation.New(console, 8)

	sim.Train(10, 30, 1000, 10, 0, 50)
}

func MustLogger() {
	logger.App = bL.MustNew("logs\\app.log", bL.ErrorLevel)
	logger.Sim = bL.MustNew("logs\\sim.log", bL.InfoLevel)
}
