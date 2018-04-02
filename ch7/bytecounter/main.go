// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// exercise 7.1
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}

// ecercise 7.2
type countingWriter struct {
	w io.Writer
	c *int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	*(cw.c) += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w, c: new(int64)}
	*(cw.c) = int64(0)
	return cw, cw.c
}

// exercise 7.5

type limitReader struct {
	n int64
	r io.Reader
}

func (lr *limitReader) Read(p []byte) (int, error) {
	var err error
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
		err = io.EOF
	}
	n, err := lr.r.Read(p)
	lr.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, n: n}
}

func limitReaderTest() {
	testStr := "1234556"
	lr := LimitReader(strings.NewReader(testStr), 3)
	lr100 := LimitReader(strings.NewReader(testStr), 100)
	result := make([]byte, 3)
	result100 := make([]byte, 100)
	n, err := lr.Read(result)
	n100, err100 := lr100.Read(result100)
	fmt.Printf("3 ?= %d", n)
	fmt.Printf("100 ?= %d", n100)
	fmt.Printf("nil ?= %s", err100)
	fmt.Printf("EOF ?= %d", err)

}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	// exercise 7.1
	var wc WordCounter
	cw, n := CountingWriter(&wc)
	_, err := fmt.Fprint(cw, "1 2 3 44444\n55\n66")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("numberOfWords: %d\n", wc)
	fmt.Printf("cw.c: %d\n", *n)
	limitReaderTest()
}
