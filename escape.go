package escape

const (
	Up = 'A' + iota
	Down
	Right
	Left
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

	return b
}

func Cursor(n int, dir byte) []byte {
	return AppendCursor(nil, n, dir)
}

func AppendCursor(b []byte, n int, dir byte) []byte {
	b = append(b, '\x1b', '[')

	w := 1
	for q := n; q >= 10; q /= 10 {
		w++
	}

	i := len(b)

	b = append(b, "            "[:w]...)

	for q, j := n, w-1; j >= 1; j-- {
		b[i+j] = byte(q%10) + '0'
		q /= 10
	}

	b = append(b, dir)

	return b
}

func AppendRaw(b, seq []byte) []byte {
	b = append(b, '\x1b', '[')
	b = append(b, seq...)

	return b
}

func AppendRawString(b []byte, seq string) []byte {
	b = append(b, '\x1b', '[')
	b = append(b, seq...)

	return b
}
