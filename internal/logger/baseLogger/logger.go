package baseLogger

import (
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	file  *os.File
	level loggerType
}

type loggerType uint8

func MustNew(path string, level loggerType) *Logger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(Fatal+"opening file: %v", err)
	}
	return &Logger{
		file:  file,
		level: level,
	}
}

func (logger *Logger) Debug(msg string) {
	if logger.level <= DebugLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+Debug+msg+"\n")
	}
}

func (logger *Logger) Info(msg string) {
	if logger.level <= InfoLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+Info+msg+"\n")
	}
}

func (logger *Logger) Panic(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+Panic+msg+"\n")
		panic(msg)
	}
}

func (logger *Logger) Error(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+Error+msg+"\n")
	}
}

func (logger *Logger) Fatal(msg string) {
	if logger.level <= ErrorLevel {
		t := time.Now()
		_, _ = io.WriteString(logger.file, t.String()+Fatal+msg+"\n")
		os.Exit(1)
	}
}
