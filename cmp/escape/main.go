package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nikandfor/cli"
	"github.com/nikandfor/errors"

	"github.com/nikandfor/escape"
	"github.com/nikandfor/escape/color"
)

func main() {
	cli.App = cli.Command{
		Name:   "escape sequence generator",
		Action: run,
		Args:   cli.Args{},
		Flags: []*cli.Flag{
			cli.NewFlag("n", false, "do not output the trailing newline"),
		},
	}

	cli.RunAndExit(os.Args)
}

func run(c *cli.Command) (err error) {
	var b []byte

	for i, a := range c.Args {
		if !strings.HasPrefix(a, "-") {
			fmt.Printf("%s%s", b, a)

			b = b[:0]

			continue
		}

		for len(a) != 0 && a[0] == '-' {
			a = a[1:]
		}

		if a == "reset" || a == "r" {
			b = color.Append(b[:0], color.Reset)

			continue
		}

		if strings.HasPrefix(a, "raw=") {
			a = strings.TrimPrefix(a, "raw=")

			b = escape.AppendRawString(b, a)

			continue
		}

		var col int
		if strings.HasPrefix(a, "bg-") {
			col += color.Background

			a = strings.TrimPrefix(a, "bg-")
		}

		var col2 int
		if strings.HasPrefix(a, "bright-") {
			col2 = color.Bright

			a = strings.TrimPrefix(a, "bright-")
		}

		switch a {
		case "bold":
			col += color.Bold
		case "underline":
			col += color.Underline
		case "reversed":
			col += color.Reversed

		case "black":
			col += color.Black
		case "red":
			col += color.Red
		case "green":
			col += color.Green
		case "yellow":
			col += color.Yellow
		case "blue":
			col += color.Blue
		case "magenta":
			col += color.Magenta
		case "cyan":
			col += color.Cyan
		case "white":
			col += color.White
		default:
			return errors.New("unknown color: %v", c.Args[i])
		}

		if col2 != 0 {
			b = color.Append(b, col, col2)
		} else {
			b = color.Append(b, col)
		}
	}

	if len(b) != 0 {
		fmt.Printf("%s", b)
	}

	if !c.Bool("n") {
		fmt.Println()
	}

	return
}
