package chalk

import "fmt"

type style struct {
	Open  string
	Close string
}

func newStyle(a, b int) style {
	return style{
		Open:  fmt.Sprintf("\u001B[%dm", a),
		Close: fmt.Sprintf("\u001B[%dm", b),
	}
}

var Modifier = map[string]style{
	"reset":         newStyle(0, 0),
	"bold":          newStyle(1, 22),
	"dim":           newStyle(2, 22),
	"italic":        newStyle(3, 23),
	"underline":     newStyle(4, 24),
	"inverse":       newStyle(7, 27),
	"hidden":        newStyle(8, 28),
	"strikethrough": newStyle(9, 29),
}

var Color = map[string]style{
	"black":   newStyle(30, 39),
	"red":     newStyle(31, 39),
	"green":   newStyle(32, 39),
	"yellow":  newStyle(33, 39),
	"blue":    newStyle(34, 39),
	"magenta": newStyle(35, 39),
	"cyan":    newStyle(36, 39),
	"white":   newStyle(37, 39),

	// bright color
	"blackBright":   newStyle(90, 39),
	"redBright":     newStyle(91, 39),
	"greenBright":   newStyle(92, 39),
	"yellowBright":  newStyle(93, 39),
	"blueBright":    newStyle(94, 39),
	"magentaBright": newStyle(95, 39),
	"cyanBright":    newStyle(96, 39),
	"whiteBright":   newStyle(97, 39),
}

var BgColor = map[string]style{
	"bgBlack":   newStyle(40, 49),
	"bgRed":     newStyle(41, 49),
	"bgGreen":   newStyle(42, 49),
	"bgYellow":  newStyle(43, 49),
	"bgBlue":    newStyle(44, 49),
	"bgMagenta": newStyle(45, 49),
	"bgCyan":    newStyle(46, 49),
	"bgWhite":   newStyle(47, 49),

	// bright color
	"bgBlackBright":   newStyle(100, 49),
	"bgRedBright":     newStyle(101, 49),
	"bgGreenBright":   newStyle(102, 49),
	"bgYellowBright":  newStyle(103, 49),
	"bgBlueBright":    newStyle(104, 49),
	"bgMagentaBright": newStyle(105, 49),
	"bgCyanBright":    newStyle(106, 49),
	"bgWhiteBright":   newStyle(107, 49),
}
