package main

import (
	"errors"
	"fmt"
	"io"
)

type MyWriter struct {
	data string
	size int
}

func (mw *MyWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, io.EOF
	}

	n := mw.size
	var err error = nil
	if len(p) < mw.size {
		n = len(p)
	} else {
		err = errors.New("p is larger than size")
	}

	mw.data = mw.data + string(p[0:n])

	return n, err

}

func main() {
	msg := []byte("the world with Go!!!")
	mw := MyWriter{"Save ", 7}
	i := 0
	var err error

	for err == nil && i < len(msg) {
		n, err := mw.Write(msg[i:])
		fmt.Printf("Written %d error %v —> %s\n", n, err, mw.data)
		i = i + n
	}
}
