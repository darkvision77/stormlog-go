package colors

type Color string

const (
	nc     Color = "\033[0m"
	RED    Color = "\033[0;31m"
	YELLOW Color = "\033[0;33m"
	GRAY   Color = "\033[0;90m"
	BLUE   Color = "\033[0;34m"
	GREEN  Color = "\033[0;32m"
	WHITE  Color = "\033[0;37m"
)

func Colorize(s string, color Color) string {
	return string(color)+s+string(nc)
}
