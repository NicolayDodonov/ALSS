package main

import (
	"artificialLifeGo/internal/config"
	oTC "artificialLifeGo/internal/console/oldTextConsole"
	l "artificialLifeGo/internal/logger"
	bL "artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/model"
	"artificialLifeGo/internal/simulation"
)

func main() {
	//инициалируем программу
	MustInit()
	l.App.Info("Application is run")
	defer l.App.Info("Application exit")

	//Включаем консоль
	console := oTC.New()
	sim := simulation.New(console, 8)
	l.App.Info("Console init")

	//начинаем обучение
	l.App.Info("Simulation is run")
	sim.Train(10, 30, 1000, 10, 0, 50)
}

func MustInit() {
	conf := config.MustLoad("config/config.yaml")
	l.App = bL.MustNew("logs\\app.log", bL.Convert(conf.App))
	l.Ent = bL.MustNew("logs\\ent.log", bL.Convert(conf.Ent))
	l.Sim = bL.MustNew("logs\\sim.log", bL.Convert(conf.Sim))

	model.MaxGen = conf.Max
	model.LengthDNA = conf.Length
	model.EnergyPoint = conf.Energy
	model.TypeBrain = conf.Brain
}
