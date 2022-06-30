package stormlog

import (
	"fmt"
	"time"
)

type Logger interface {
	AddListener(listener EventListener)
	Event(event Event)

	Trace(v ...any)
	Tracef(format string, v ...any)

	Debug(v ...any)
	Debugf(format string, v ...any)

	Info(v ...any)
	Infof(format string, v ...any)

	Warning(v ...any)
	Warningf(format string, v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)

	Critical(v ...any)
	Criticalf(format string, v ...any)
}

type logger struct {
	Listeners []EventListener
}

func New() Logger {
	return &logger{
		Listeners: []EventListener{},
	}
}

func (log *logger) AddListener(listener EventListener) {
	log.Listeners = append(log.Listeners, listener)
}

func (log *logger) Event(event Event) {
	for _, i := range log.Listeners {
		i.Handle(event)
	}
}

func (log *logger) makeEvent(level Level, message string) {
	e := Event{
		Level: level,
		Timestamp: time.Now(),
		Message: message,
	}
	log.Event(e)
}

func (log *logger) Trace(v ...any) {
	log.makeEvent(TRACE, fmt.Sprint(v...))
}
func (log *logger) Tracef(format string, v ...any) {
	log.makeEvent(TRACE, fmt.Sprintf(format, v...))
}

func (log *logger) Debug(v ...any) {
	log.makeEvent(DEBUG, fmt.Sprint(v...))
}
func (log *logger) Debugf(format string, v ...any) {
	log.makeEvent(DEBUG, fmt.Sprintf(format, v...))
}

func (log *logger) Info(v ...any) {
	log.makeEvent(INFO, fmt.Sprint(v...))
}
func (log *logger) Infof(format string, v ...any) {
	log.makeEvent(INFO, fmt.Sprintf(format, v...))
}

func (log *logger) Warning(v ...any) {
	log.makeEvent(WARNING, fmt.Sprint(v...))
}
func (log *logger) Warningf(format string, v ...any) {
	log.makeEvent(WARNING, fmt.Sprintf(format, v...))
}

func (log *logger) Error(v ...any) {
	log.makeEvent(ERROR, fmt.Sprint(v...))
}
func (log *logger) Errorf(format string, v ...any) {
	log.makeEvent(ERROR, fmt.Sprintf(format, v...))
}

func (log *logger) Critical(v ...any) {
	log.makeEvent(CRITICAL, fmt.Sprint(v...))
}
func (log *logger) Criticalf(format string, v ...any) {
	log.makeEvent(CRITICAL, fmt.Sprintf(format, v...))
}
