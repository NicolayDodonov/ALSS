package main

import (
	"artificialLifeGo/internal/alModel"
	"artificialLifeGo/internal/config"
	oldConsole "artificialLifeGo/internal/console/oldTextConsole"
	loggers "artificialLifeGo/internal/logger"
	bLogger "artificialLifeGo/internal/logger/baseLogger"
	sim "artificialLifeGo/internal/simulation"
	"artificialLifeGo/internal/storage/file"
	"log"
)

func main() {
	//инициалируем программу
	conf := MustInit()
	loggers.App.Info("Application is run")
	defer loggers.App.Info("Application exit")

	//Включаем консоль
	console := oldConsole.New(conf.Console.TimeOut)
	storage := file.New(conf.Storage.PathAge, conf.PathTrain)
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

	alModel.MaxGen = conf.MaxGen
	alModel.MaxFoodPercent = conf.MaxFood
	alModel.LengthDNA = conf.Length
	alModel.EnergyPoint = conf.Energy
	alModel.TypeBrain = conf.Brain
	alModel.LoopX = conf.Loop.X
	alModel.LoopY = conf.Loop.Y
	alModel.PoisonEnable = conf.PoisonEnable
	alModel.BasePoisonLevel = conf.BaseLevel

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
