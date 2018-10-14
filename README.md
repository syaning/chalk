# chalk

Terminal string styling. Inspired by [@chalk/chalk](https://github.com/chalk/chalk).

## Installation

```bash
$ go get github.com/syaning/chalk
```

## Usage

```go
package main

import (
	"chalk"
	"fmt"
)

func main() {
	fmt.Println(
		chalk.Reset("reset"),
		chalk.Bold("bold"),
		chalk.Dim("dim"),
		chalk.Italic("italic"),
		chalk.Underline("underline"),
		chalk.Inverse("inverse"),
		chalk.Hidden("hidden"),
		chalk.Strikethrough("strikethrough"),
	)

	fmt.Println(
		chalk.Black("black"),
		chalk.Red("red"),
		chalk.Green("green"),
		chalk.Yellow("yellow"),
		chalk.Blue("blue"),
		chalk.Magenta("magenta"),
		chalk.Cyan("cyan"),
		chalk.White("white"),
		chalk.Gray("gray"),
		chalk.Grey("grey"),
	)

	fmt.Println(
		chalk.BlackBright("blackBright"),
		chalk.RedBright("redBright"),
		chalk.GreenBright("greenBright"),
		chalk.YellowBright("yellowBright"),
		chalk.BlueBright("blueBright"),
		chalk.MagentaBright("magentaBright"),
		chalk.CyanBright("cyanBright"),
		chalk.WhiteBright("whiteBright"),
	)

	fmt.Println(
		chalk.BgBlack(chalk.White("bgBlack")),
		chalk.BgRed(chalk.White("bgRed")),
		chalk.BgGreen("bgGreen"),
		chalk.BgYellow("bgYellow"),
		chalk.BgBlue("bgBlue"),
		chalk.BgMagenta("bgMagenta"),
		chalk.BgCyan("bgCyan"),
		chalk.BgWhite("bgWhite"),
		chalk.BgGray("bgGray"),
		chalk.BgGrey("bgGrey"),
	)

	fmt.Println(
		chalk.BgBlackBright("bgBlackBright"),
		chalk.BgRedBright("bgRedBright"),
		chalk.BgGreenBright("bgGreenBright"),
		chalk.BgYellowBright("bgYellowBright"),
		chalk.BgBlueBright("bgBlueBright"),
		chalk.BgMagentaBright("bgMagentaBright"),
		chalk.BgCyanBright("bgCyanBright"),
		chalk.BgWhiteBright("bgWhiteBright"),
	)
}
```