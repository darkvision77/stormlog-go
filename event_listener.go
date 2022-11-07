package stormlog

type EventListener interface {
	Handle(e Event)
	Sync() error
	Close() error
}
