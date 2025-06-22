package main

import (
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger"
	"artificialLifeGo/internal/logger/baseLogger"
	"artificialLifeGo/internal/server"
	"log"
)

// main это точка входа в программу.
func main() {
	//читаем файл конфигуратор
	c, l := mustInit()

	l.Info("application start")

	//запускаем сервер
	s := server.New(c, l)
	if err := s.Start(); err != nil {
		l.Error(err.Error())
	}
}

// mustInit это функция инициализации. Возвращает объекты типа *config.Config и *baseLogger.Logger.
func mustInit() (*config.Config, *baseLogger.Logger) {
	//читаем файл конфигуратор
	с := config.MustLoad("config/config.yaml")
	//создаём логгер
	l, err := baseLogger.New(с.Logger.Path, logger.Convert(с.Logger.Level))
	if err != nil {
		log.Panicf(err.Error())
	}
	return с, l
}
