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
}

type Logger struct {
	Type    string `yaml:"typeLogger" env-required:"true"`
	App     string `yaml:"appLevel" env-default:"Error"`
	Ent     string `yaml:"entLevel" env-default:"Error"`
	Sim     string `yaml:"simLevel" env-default:"Error"`
	AppPath string `yaml:"appPath" env-default:"log/all.log"`
	EntPath string `yaml:"entPath" env-default:"log/all.log"`
	SimPath string `yaml:"simPath" env-default:"log/all.log"`
}

type Model struct {
	Max    int    `yaml:"maxGen" env-default:"16"`
	Length int    `yaml:"lengthDNA" env-default:"64"`
	Energy int    `yaml:"energyPoint" env-default:"10"`
	Brain  string `yaml:"brain" env-default:"brain16"`
}

type Simulation struct {
	Type              string `yaml:"typeSimulation" env-required:"true"`
	WorldSizeX        int    `yaml:"X" env-default:"10"`
	WorldSizeY        int    `yaml:"Y" env-default:"10"`
	EndPopulation     int    `yaml:"endPop" env-default:""`
	recurseUpdateRate int    `yaml:"resurceUpdate" env-default:""`
	FinalAgeTrain     int    `yaml:"ageExit" env-default:""`
	MutationCount     int    `yaml:"mutation" env-default:""`
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
