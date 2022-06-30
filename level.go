package stormlog

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	CRITICAL
)

func (level Level) String() string {
	switch level {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	}
	return "UNKNOWN"
}
