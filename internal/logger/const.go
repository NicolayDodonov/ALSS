package logger

const (
	Debug = "[DEB] "
	Info  = "[INF] "
	Panic = "[PNC] "
	Error = "[ERR] "
	Fatal = "[FAT] "
)

const (
	DebugLevel LoggerType = 1
	InfoLevel  LoggerType = 2
	ErrorLevel LoggerType = 3
	OffLevel   LoggerType = 10
)

// LoggerType это уровень логгирования.
type LoggerType uint8

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
