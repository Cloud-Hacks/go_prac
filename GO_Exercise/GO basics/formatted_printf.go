package main

import (
	"fmt"
	"io"
	"os"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func main() {

	x, y := Fprintf(os.Stdout, "hey")
	fmt.Println(x, y)
}
