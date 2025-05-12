package main

import (
	"artificialLifeGo/internal/server"
	"log"
)

const (
	adr = "ip adress"
)

func main() {
	s := server.New(adr)
	if err := s.Start(); err != nil {
		log.Print(err)
	}
}
