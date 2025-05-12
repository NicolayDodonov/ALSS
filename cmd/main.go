package main

import (
	"artificialLifeGo/internal/server"
	"log"
)

const (
	adr = "192.168.1.42:8080"
)

func main() {
	server := server.New(adr)
	if err := server.Start(); err != nil {
		log.Print(err)
	}
}
