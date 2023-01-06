package helpers

import (
	"log"
	"os"
)

const (
	INFO  = ": ℹ INFO\t"
	DEBUG = ": ⚠ DEBUG\t"
	ERROR = ": ✖ ERROR\t"
)

type Loggers struct {
	service  string
	infoLog  *log.Logger
	errorLog *log.Logger
	debugLog *log.Logger
}

func NewLoggers(service string) *Loggers {
	logger := &Loggers{
		service:  service,
		infoLog:  log.New(os.Stdout, service+INFO, log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, service+ERROR, log.Ldate|log.Ltime),
		debugLog: log.New(os.Stdout, service+DEBUG, log.Ldate|log.Ltime),
	}
	return logger
}

func (l *Loggers) Info(msg string, args ...any) {
	if len(args) > 0 {
		l.infoLog.Println(msg, args)
	} else {
		l.infoLog.Println(msg)
	}
}

func (l *Loggers) Error(msg string, args ...any) {
	if len(args) > 0 {
		l.errorLog.Println(msg, args)
	} else {
		l.errorLog.Println(msg)
	}
}

func (l *Loggers) Debug(msg string, args ...any) {
	if len(args) > 0 {
		l.debugLog.Println(msg, args)
	} else {
		l.debugLog.Println(msg)
	}
}

func (l *Loggers) Fatal(msg string, args ...any) {
	if len(args) > 0 {
		l.errorLog.Fatalln(msg, args)
	} else {
		l.errorLog.Fatalln(msg)
	}
}

func (l *Loggers) Panic(msg string, args ...any) {
	if len(args) > 0 {
		l.errorLog.Panicln(msg, args)
	} else {
		l.errorLog.Panicln(msg)
	}
}
