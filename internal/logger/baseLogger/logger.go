package baseLogger

import (
	log "artificialLifeGo/internal/logger"
	"fmt"
	"io"
	"os"
	"time"
)

// Logger структура базового логгера, хранит в себе адрес хрангилища логгеров и уровень логгирования.
type Logger struct {
	file  *os.File
	level log.LoggerType
}

// New создаёт новый логгер и возвращает или указатель на него или ошибку.
func New(path string, level log.LoggerType) (*Logger, error) {
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
	if logger.level <= log.DebugLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+log.Debug+msg+"\n")
	}
}

func (logger *Logger) Info(msg string) {
	if logger.level <= log.InfoLevel {
		t := time.Now()
		message := t.String() + " " + log.Info + msg + "\n"
		fmt.Print(message)
		_, _ = io.WriteString(logger.file, message)

	}
}

func (logger *Logger) Panic(msg string) {
	if logger.level <= log.ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+log.Panic+msg+"\n")
		panic(msg)
	}
}

func (logger *Logger) Error(msg string) {
	if logger.level <= log.ErrorLevel {
		t := time.Now()
		message := t.String() + " " + log.Error + msg + "\n"
		fmt.Print(message)
		_, _ = io.WriteString(logger.file, message)
	}
}

func (logger *Logger) Fatal(msg string) {
	if logger.level <= log.ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+" "+log.Fatal+msg+"\n")
		os.Exit(1)
	}
}
