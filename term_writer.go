package stormlog

import (
	"fmt"
	"os"

	"github.com/darkvision77/stormlog-go/colors"
)

type TermWriter struct {
	ColorsEnabled bool
	ColorMap map[Level]colors.Color
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
	}
}

func (w *TermWriter) Handle(e Event) {
	s := e.String()
	if w.ColorsEnabled {
		s = colors.Colorize(s, w.ColorMap[e.Level])
	}

	if e.Level >= WARNING {
		fmt.Fprintln(os.Stderr, s)
	} else {
		fmt.Fprintln(os.Stdout, s)
	}
}

func (w *TermWriter) Sync() error {
	return nil
}

func (w *TermWriter) Close() error {
	return nil
}
