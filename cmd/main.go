package main

import (
	"artificialLifeGo/internal/config"
	oldTextConsole "artificialLifeGo/internal/console/oldTextConsole"
	loggers "artificialLifeGo/internal/logger"
	bLogger "artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/model"
	sim "artificialLifeGo/internal/simulation"
	"artificialLifeGo/internal/storage/fileSt"
	"log"
)

func main() {
	//инициалируем программу
	conf := MustInit()
	loggers.App.Info("Application is run")
	defer loggers.App.Info("Application exit")

	//Включаем консоль
	console := oldTextConsole.New()
	storage := fileSt.New(conf.Storage.PathAge, conf.PathTrain)
	simulation := sim.New(console, storage)
	loggers.App.Info("Console init")

	//начинаем обучение
	loggers.App.Info("Simulation is run")
	_ = simulation.Train()
}

func MustInit() *config.Config {
	var err error
	conf := config.MustLoad("config/config.yaml")
	loggers.App, err = bLogger.New(conf.Logger.PathAge, bLogger.Convert(conf.App))
	if err != nil {
		log.Fatal(err)
	}
	loggers.Ent, err = bLogger.New(conf.PathEnt, bLogger.Convert(conf.Ent))
	if err != nil {
		log.Fatal(err)
	}
	loggers.Sim, err = bLogger.New(conf.PathSim, bLogger.Convert(conf.Sim))
	if err != nil {
		log.Fatal(err)
	}

	model.MaxGen = conf.Max
	model.LengthDNA = conf.Length
	model.EnergyPoint = conf.Energy
	model.TypeBrain = conf.Brain
	model.LoopX = conf.Loop.X
	model.LoopY = conf.Loop.Y

	sim.TypeSimulation = conf.Simulation.Type
	sim.WorldSizeX = conf.Simulation.WorldSizeX
	sim.WorldSizeY = conf.Simulation.WorldSizeY
	sim.StartPopulation = conf.Simulation.StartPopulation
	sim.EndPopulation = conf.Simulation.EndPopulation
	sim.RecurseUpdateRate = conf.Simulation.RecurseUpdateRate
	sim.FinalAgeTrain = conf.Simulation.FinalAgeTrain
	sim.MutationCount = conf.Simulation.MutationCount
	return conf
}
