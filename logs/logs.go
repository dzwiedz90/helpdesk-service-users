package logs

type LoggerInterface interface {
	InfoLogger(message string)
	ErrorLogger(message string)
}
