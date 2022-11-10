package stormlog

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/darkvision77/stormlog-go/internal"
)

type Logger interface {
	AddListener(listener EventListener)
	Event(event Event)

	SetName(name string)
	Module(name string) Logger

	Sync() error
	Close() error

	HandlePanic()

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
	moduleName string
}

func New() Logger {
	return &logger{
		Listeners: []EventListener{},
		moduleName: "main",
	}
}

func (log *logger) HandlePanic() {
	err := recover()
	if err != nil {
		log.Debugf("Stacktrace:\n%s", string(debug.Stack()))
		log.Criticalf("Unhandled Exception: %v", err)
	}
}

func (log *logger) AddListener(listener EventListener) {
	log.Listeners = append(log.Listeners, listener)
}

func (log *logger) Event(event Event) {
	for _, i := range log.Listeners {
		i.Handle(event)
	}
	if event.Level >= CRITICAL {
		panic(event.Message)
	}
}

func (log *logger) SetName(name string) {
	log.moduleName = name
}

func (log *logger) Module(name string) Logger {
	return &logger{Listeners: log.Listeners, moduleName: name}
}

func (log *logger) Sync() error {
	for _, i := range log.Listeners {
		err := i.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func (log *logger) Close() error {
	for _, i := range log.Listeners {
		err := i.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (log *logger) makeEvent(level Level, message string) {
	e := Event{
		Level:     level,
		Module:    log.moduleName,
		Thread:    internal.GetGoroutineId(),
		Timestamp: time.Now(),
		Message:   message,
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
