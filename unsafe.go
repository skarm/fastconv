//go:build go1.20
// +build go1.20

package fastconv

import "unsafe"

// S2B converts string to a byte slice without memory allocation.
func S2B(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// B2S converts byte slice to a string without memory allocation.
func B2S(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
