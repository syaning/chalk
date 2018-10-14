package chalk

func Format(str string, mode style) string {
	return mode.Open + str + mode.Close
}

// modifier

func Reset(str string) string {
	return Format(str, Modifier["reset"])
}

func Bold(str string) string {
	return Format(str, Modifier["bold"])
}

func Dim(str string) string {
	return Format(str, Modifier["dim"])
}

func Italic(str string) string {
	return Format(str, Modifier["italic"])
}

func Underline(str string) string {
	return Format(str, Modifier["underline"])
}

func Inverse(str string) string {
	return Format(str, Modifier["inverse"])
}

func Hidden(str string) string {
	return Format(str, Modifier["hidden"])
}

func Strikethrough(str string) string {
	return Format(str, Modifier["strikethrough"])
}

// color

func Black(str string) string {
	return Format(str, Color["black"])
}

func Red(str string) string {
	return Format(str, Color["red"])
}

func Green(str string) string {
	return Format(str, Color["green"])
}

func Yellow(str string) string {
	return Format(str, Color["yellow"])
}

func Blue(str string) string {
	return Format(str, Color["blue"])
}

func Magenta(str string) string {
	return Format(str, Color["magenta"])
}

func Cyan(str string) string {
	return Format(str, Color["cyan"])
}

func White(str string) string {
	return Format(str, Color["white"])
}

// bright color

func BlackBright(str string) string {
	return Format(str, Color["blackBright"])
}

func RedBright(str string) string {
	return Format(str, Color["redBright"])
}

func GreenBright(str string) string {
	return Format(str, Color["greenBright"])
}

func YellowBright(str string) string {
	return Format(str, Color["yellowBright"])
}

func BlueBright(str string) string {
	return Format(str, Color["blueBright"])
}

func MagentaBright(str string) string {
	return Format(str, Color["magentaBright"])
}

func CyanBright(str string) string {
	return Format(str, Color["cyanBright"])
}

func WhiteBright(str string) string {
	return Format(str, Color["whiteBright"])
}

func Gray(str string) string {
	return BlackBright(str)
}

func Grey(str string) string {
	return BlackBright(str)
}

// background color

func BgBlack(str string) string {
	return Format(str, BgColor["bgBlack"])
}

func BgRed(str string) string {
	return Format(str, BgColor["bgRed"])
}

func BgGreen(str string) string {
	return Format(str, BgColor["bgGreen"])
}

func BgYellow(str string) string {
	return Format(str, BgColor["bgYellow"])
}

func BgBlue(str string) string {
	return Format(str, BgColor["bgBlue"])
}

func BgMagenta(str string) string {
	return Format(str, BgColor["bgMagenta"])
}

func BgCyan(str string) string {
	return Format(str, BgColor["bgCyan"])
}

func BgWhite(str string) string {
	return Format(str, BgColor["bgWhite"])
}

// bright background color

func BgBlackBright(str string) string {
	return Format(str, BgColor["bgBlackBright"])
}

func BgRedBright(str string) string {
	return Format(str, BgColor["bgRedBright"])
}

func BgGreenBright(str string) string {
	return Format(str, BgColor["bgGreenBright"])
}

func BgYellowBright(str string) string {
	return Format(str, BgColor["bgYellowBright"])
}

func BgBlueBright(str string) string {
	return Format(str, BgColor["bgBlueBright"])
}

func BgMagentaBright(str string) string {
	return Format(str, BgColor["bgMagentaBright"])
}

func BgCyanBright(str string) string {
	return Format(str, BgColor["bgCyanBright"])
}

func BgWhiteBright(str string) string {
	return Format(str, BgColor["bgWhiteBright"])
}

func BgGray(str string) string {
	return BgBlackBright(str)
}

func BgGrey(str string) string {
	return BgBlackBright(str)
}
