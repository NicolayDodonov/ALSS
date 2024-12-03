package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Logger     `yaml:"logger"`
	Model      `yaml:"modelConst"`
	Simulation `yaml:"simulation"`
	Storage    `yaml:"storage"`
}

type Logger struct {
	Type    string `yaml:"type" env-required:"true"`
	App     string `yaml:"appLevel" env-default:"Off"`
	Ent     string `yaml:"entLevel" env-default:"Off"`
	Sim     string `yaml:"simLevel" env-default:"Off"`
	PathAge string `yaml:"appPath"`
	PathEnt string `yaml:"entPath"`
	PathSim string `yaml:"simPath"`
}

type Model struct {
	MaxGen  int    `yaml:"maxGen"`
	MaxFood int    `yaml:"foodPercent"`
	Length  int    `yaml:"lengthDNA"`
	Energy  int    `yaml:"energyPoint"`
	Brain   string `yaml:"brain"`
	Loop    `yaml:"loop"`
}

type Loop struct {
	X bool `yaml:"x" env-default:"false"`
	Y bool `yaml:"y" env-default:"false"`
}

type Simulation struct {
	Type              string `yaml:"type" env-required:"true"`
	WorldSizeX        int    `yaml:"X" env-default:"10"`
	WorldSizeY        int    `yaml:"Y" env-default:"10"`
	StartPopulation   int    `yaml:"startPop"`
	EndPopulation     int    `yaml:"endPop"`
	RecurseUpdateRate int    `yaml:"resourceUpdate"`
	FinalAgeTrain     int    `yaml:"ageExit"`
	MutationCount     int    `yaml:"mutation"`
	Console           `yaml:"console"`
}

type Console struct {
	Type    string `yaml:"type" env-required:"true"`
	TimeOut int    `yaml:"timeOut"`
}

type Storage struct {
	Type      string `yaml:"type" env-required:"true"`
	PathAge   string `yaml:"pathAge"`
	PathTrain string `yaml:"pathTrain"`
}

func MustLoad(path string) *Config {
	configPath := path
	if configPath == "" {
		log.Fatal("config is not found")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
