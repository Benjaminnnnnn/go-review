package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range b[:n] {
		v := int(b[i])
		if v >= 65 && v <= 90 {
			b[i] = byte(65 + (v-52)%26)
		} else if v >= 97 && v <= 122 {
			b[i] = byte(97 + (v-84)%26)
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
