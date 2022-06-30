package stormlog

import (
	"fmt"
	"io"
)

type StreamWriter struct {
	Stream io.Writer
}

func NewStreamWriter(w io.Writer) *StreamWriter {
	return &StreamWriter{Stream: w}
}

func (w *StreamWriter) Handle(e Event) {
	fmt.Fprintln(w.Stream, e.String())
}

func (w *StreamWriter) SetStream(stream io.Writer) {
	w.Stream = stream
}
