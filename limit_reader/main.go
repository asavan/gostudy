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
	if 0 == g.limit {
		return 0, io.EOF
	}
	if len(bytes) > g.limit {
		bytes = bytes[0:g.limit]
	}

	m, err := g.r.Read(bytes)
	g.limit -= m
	if 0 == g.limit && err == nil {
		return m, io.EOF
	}

	return m, err
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
