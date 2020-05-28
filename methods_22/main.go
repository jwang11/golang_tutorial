/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import (
	"golang.org/x/tour/reader"
)
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {
	blen := len(b)
	for i := 0; i < blen; i++ {
		b[i] = 'A'
	}
	return blen, nil
}

func main() {
	reader.Validate(MyReader{})
}
