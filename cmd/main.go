package main

import (
	oTC "artificialLifeGo/internal/console/oldTextConsole"
	"artificialLifeGo/internal/simulation"
)

func main() {
	console := oTC.New()
	sim := simulation.New(console, 8)
}
