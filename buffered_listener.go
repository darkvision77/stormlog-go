package stormlog

import "sync"

type BufferedListener struct {
	Listener EventListener
	channel  chan Event
	wg       sync.WaitGroup
}

func Buffered(listener EventListener, bufferSize ...int) *BufferedListener {
	bs := 256
	if len(bufferSize) != 0 {
		bs = bufferSize[0]
	}
	w := &BufferedListener{
		Listener: listener,
		channel:  make(chan Event, bs),
		wg:       sync.WaitGroup{},
	}
	go w.runWorker()
	return w
}

func (w *BufferedListener) Handle(e Event) {
	w.wg.Add(1)
	w.channel <- e
}

// TODO
func (w *BufferedListener) runWorker() {
	for {
		e := <-w.channel
		w.Listener.Handle(e)
		w.wg.Done()
	}
}

func (w *BufferedListener) Sync() {
	w.wg.Wait()
}
