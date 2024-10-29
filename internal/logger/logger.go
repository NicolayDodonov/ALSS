package logger

type Logger interface {
	Debug(string)
	Info(string)
	Panic(string)
	Error(string)
	Fatal(string)
}

var App Logger
var Sim Logger
