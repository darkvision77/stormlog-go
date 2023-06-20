package stormlog

import (
	"fmt"
	"io"
	"os"

	"github.com/darkvision77/stormlog-go/colors"
)

type TermWriter struct {
	ColorsEnabled bool
	ColorMap map[Level]colors.Color
	Output io.Writer
}

func NewTermWriter(colorsEnabled bool) *TermWriter {
	return &TermWriter{
		ColorsEnabled: colorsEnabled,
		ColorMap: map[Level]colors.Color{
			TRACE:    colors.GRAY,
			DEBUG:    colors.WHITE,
			INFO:     colors.GREEN,
			WARNING:  colors.YELLOW,
			ERROR:    colors.RED,
			CRITICAL: colors.RED,
		},
		Output: nil,
	}
}

func (w *TermWriter) WithOutput(out io.Writer) *TermWriter {
	w.Output = out
	return w
}

func (w *TermWriter) Handle(e Event) {
	s := e.String()
	if w.ColorsEnabled {
		s = colors.Colorize(s, w.ColorMap[e.Level])
	}

	var out io.Writer = os.Stdout
	if w.Output != nil {
		out = w.Output
	} else if e.Level >= WARNING {
		out = os.Stderr
	}
	fmt.Fprintln(out, s)
}

func (w *TermWriter) Sync() error {
	return nil
}

func (w *TermWriter) Close() error {
	return nil
}
