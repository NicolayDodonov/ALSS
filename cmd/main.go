package main

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/server"
	"log"
)

const (
	adr = "ip adress"
)

func main() {
	conf := config.MustLoad("config/config.yaml")
	l, err := baseLogger.New(conf.Logger.Path, baseLogger.Convert(conf.Logger.Level))
	if err != nil {
		log.Fatalf(err.Error())
	}
	s := server.New(conf, l)
	if err := s.Start(); err != nil {
		log.Print(err)
	}
}
