package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Server
	ALSS
}

type Server struct {
	IP   string `env:"ip" envDefault:"127.0.0.1"`
	Port string `env:"port" envDefault:"8080"`
}

type ALSS struct {
	World
	Agent
}

type World struct {
	X int `env:"x" envDefault:"10"`
	Y int `env:"y" envDefault:"10"`
}

type Agent struct {
	TypeGenome         string `env:"typeGenome"`
	SizeGenome         int    `env:"sizeGenome"`
	BaseEnergy         int    `env:"baseEnergy"`
	ActionCost         int    `env:"actionCost"`
	PollutionCost      int    `env:"pollutionCost"`
	HuntingCoefficient int    `env:"huntingCoefficient"`
	BirthCost          int    `env:"birthCost"`
	MutationCount      int    `env:"mutationCount"`
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
