package main

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/server"
	"log"
)

const (
	adr = "ip adress"
)

func main() {
	conf := config.MustLoad("config/config.yaml")
	s := server.New(conf)
	if err := s.Start(); err != nil {
		log.Print(err)
	}
}
