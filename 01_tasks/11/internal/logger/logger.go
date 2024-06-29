package logger

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func New() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Println("INFO", msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Println("DEBUG", msg)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Println("ERROR", msg, err.Error())
}

func (l *Logger) Warn(msg string) {
	l.logger.Println("WARN", msg)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg, err.Error())
}
