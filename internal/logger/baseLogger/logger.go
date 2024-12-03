package baseLogger

import (
	"io"
	"os"
	"time"
)

// Logger структура базового логгера, хранит в себе адрес хрангилища логгеров и уровень логгирования.
type Logger struct {
	file  *os.File
	level LoggerType
}

// LoggerType это уровень логгирования.
type LoggerType uint8

// New создаёт новый логгер и возвращает или указатель на него или ошибку.
func New(path string, level LoggerType) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{}, err
	}
	return &Logger{
		file:  file,
		level: level,
	}, nil
}

func (logger *Logger) Debug(msg string) {
	if logger.level <= DebugLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Debug+msg+"\n")
	}
}

func (logger *Logger) Info(msg string) {
	if logger.level <= InfoLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Info+msg+"\n")
	}
}

func (logger *Logger) Panic(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Panic+msg+"\n")
		panic(msg)
	}
}

func (logger *Logger) Error(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Error+msg+"\n")
	}
}

func (logger *Logger) Fatal(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+Fatal+msg+"\n")
		os.Exit(1)
	}
}

func Convert(s string) LoggerType {
	switch s {
	case "Debug":
		return DebugLevel
	case "Info":
		return InfoLevel
	case "Error":
		return ErrorLevel
	case "Off":
		return OffLevel
	default:
		return OffLevel
	}
}
