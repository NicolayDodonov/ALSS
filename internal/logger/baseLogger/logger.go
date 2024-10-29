package baseLogger

type Logger struct {
	path  string
	level loggerType
}

type loggerType uint8

func New(path string, level loggerType) *Logger {
	return &Logger{
		path:  path,
		level: level,
	}
}

func (logger *Logger) Debug(msg string) {

}

func (logger *Logger) Info(msg string) {

}

func (logger *Logger) Panic(msg string) {

}

func (logger *Logger) Error(msg string) {

}

func (logger *Logger) Fatal(msg string) {

}
