package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Server `yaml:"Server"`
	ALSS   `yaml:"ALSS"`
	Logger `yaml:"Logger"`
}

type Server struct {
	IP   string `yaml:"ip" envDefault:"127.0.0.1"`
	Port string `yaml:"port" envDefault:"8080"`
}

type ALSS struct {
	World `yaml:"World"`
	Agent `yaml:"Agent"`
}

type World struct {
	X int `yaml:"x" envDefault:"10"`
	Y int `yaml:"y" envDefault:"10"`
}

type Agent struct {
	TypeGenome         string `yaml:"typeGenome"`
	SizeGenome         int    `yaml:"sizeGenome"`
	BaseEnergy         int    `yaml:"baseEnergy"`
	ActionCost         int    `yaml:"actionCost"`
	PollutionCost      int    `yaml:"pollutionCost"`
	HuntingCoefficient int    `yaml:"huntingCoefficient"`
	BirthCost          int    `yaml:"birthCost"`
	MutationCount      int    `yaml:"mutationCount"`
}

type Logger struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
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
