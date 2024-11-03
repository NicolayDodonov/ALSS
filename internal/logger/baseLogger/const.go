package baseLogger

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
