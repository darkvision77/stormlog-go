package stormlog

type EventListener interface {
	Handle(e Event)
}
