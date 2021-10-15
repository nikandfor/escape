package color

// Colors
const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	Extended256

	Bright = 1

	Background = 10
)

const (
	Reset     = 0
	Bold      = 1
	Underline = 4
	Reversed  = 7
)

func New(c ...int) []byte {
	return Append(nil, c...)
}

func Append(b []byte, c ...int) []byte {
	if len(c) == 0 {
		return b
	}

	b = append(b, '\x1b', '[')

	for i, c := range c {
		if i != 0 {
			b = append(b, ';')
		}

		switch {
		case c < 10:
			b = append(b, '0'+byte(c%10))
		case c < 100:
			b = append(b, '0'+byte(c/10), '0'+byte(c%10))
		default:
			b = append(b, '0'+byte(c/100), '0'+byte(c/10%10), '0'+byte(c%10))
		}
	}

	b = append(b, 'm')

	return b
}

func New256(c int, bg bool) []byte {
	return Append256(nil, c, bg)
}

func Append256(b []byte, c int, bg bool) []byte {
	base := Extended256
	if bg {
		base += Background
	}

	return Append(b, base, 5, c)
}
