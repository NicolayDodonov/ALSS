package main

import (
	"artificialLifeGo/internal/config"
	oTC "artificialLifeGo/internal/console/oldTextConsole"
	l "artificialLifeGo/internal/logger"
	bL "artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/model"
	sim "artificialLifeGo/internal/simulation"
	"log"
)

func main() {
	//инициалируем программу
	MustInit()
	l.App.Info("Application is run")
	defer l.App.Info("Application exit")

	//Включаем консоль
	console := oTC.New()
	simulation := sim.New(console)
	l.App.Info("Console init")

	//начинаем обучение
	l.App.Info("Simulation is run")
	_ = simulation.Train()
}

func MustInit() {
	var err error
	conf := config.MustLoad("config/config.yaml")
	l.App, err = bL.New("logs/app.log", bL.Convert(conf.App))
	if err != nil {
		log.Fatal(err)
	}
	l.Ent, err = bL.New("logs/ent.log", bL.Convert(conf.Ent))
	if err != nil {
		log.Fatal(err)
	}
	l.Sim, err = bL.New("logs/sim.log", bL.Convert(conf.Sim))
	if err != nil {
		log.Fatal(err)
	}

	model.MaxGen = conf.Max
	model.LengthDNA = conf.Length
	model.EnergyPoint = conf.Energy
	model.TypeBrain = conf.Brain

	sim.TypeSimulation = conf.Simulation.Type
	sim.WorldSizeX = conf.Simulation.WorldSizeX
	sim.WorldSizeY = conf.Simulation.WorldSizeY
	sim.StartPopulation = conf.Simulation.StartPopulation
	sim.EndPopulation = conf.Simulation.EndPopulation
	sim.RecurseUpdateRate = conf.Simulation.RecurseUpdateRate
	sim.FinalAgeTrain = conf.Simulation.FinalAgeTrain
	sim.MutationCount = conf.Simulation.MutationCount
}
