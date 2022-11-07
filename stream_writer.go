package stormlog

import (
	"fmt"
	"io"
)

type StreamWriter struct {
	Stream io.Writer
	closer io.Closer
}

func NewStreamWriter(w io.Writer) *StreamWriter {
	closer, _ := w.(io.Closer)
	return &StreamWriter{Stream: w, closer: closer}
}

func (w *StreamWriter) Handle(e Event) {
	fmt.Fprintln(w.Stream, e.String())
}

func (w *StreamWriter) SetStream(stream io.Writer) {
	w.Stream = stream
}

func (w *StreamWriter) Sync() error {
	return nil
}

func (w *StreamWriter) Close() error {
	if w.closer == nil {
		return nil
	}
	return w.closer.Close()
}
