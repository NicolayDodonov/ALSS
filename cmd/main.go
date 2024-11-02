package main

import (
	"artificialLifeGo/internal/config"
	oTC "artificialLifeGo/internal/console/oldTextConsole"
	"artificialLifeGo/internal/logger"
	bL "artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/model"
	"artificialLifeGo/internal/simulation"
)

func main() {
	conf := config.MustLoad()
	MustLogger(&conf.Logger)
	SetModel(&conf.Model)

	console := oTC.New()
	sim := simulation.New(console, 8)

	sim.Train(1000, 10)
}

func MustLogger(conf *config.Logger) {
	logger.App = bL.MustNew("logs\\app.log", bL.Convert(conf.App))
	logger.Sim = bL.MustNew("logs\\sim.log", bL.Convert(conf.Sim))
}

func SetModel(conf *config.Model) {
	model.MaxGen = conf.Max
	model.LengthDNA = conf.Length
	model.EnergyPoint = conf.Energy
}
