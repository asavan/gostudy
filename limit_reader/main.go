package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitReader struct {
	r     io.Reader
	limit int
}

func (g *limitReader) Read(bytes []byte) (n int, err error) { // error — это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
	if len(bytes) <= g.limit {
		m, err := g.r.Read(bytes)
		g.limit -= m
		return m, err
	}
	buf := make([]byte, g.limit)

	m, err := g.r.Read(buf)
	if err != nil {
		return m, err
	}
	for i := 0; i < m; i++ {
		bytes[i] = buf[i]
	}
	g.limit -= m
	if 0 == g.limit {
		return m, io.EOF
	}

	return m, nil
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{
		r:     r,
		limit: n,
	}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
