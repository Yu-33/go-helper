package gconv

import "unsafe"

// BytesToString converts bytes to string without memory copy.
//
// NOTICE: This operations is unsafe, any changed to b will affect s.
func BytesToString(b []byte) (s string) {
	s = *(*string)(unsafe.Pointer(&b))
	return
}

// StringToBytes converts string to bytes without memory copy
//
// NOTICE: This operations is unsafe, any changed to b will affect s.
func StringToBytes(s string) (b []byte) {
	// string's header, see https://golang.org/pkg/reflect/#StringHeader
	sh := (*[2]uintptr)(unsafe.Pointer(&s))
	// slice's header, see https://golang.org/pkg/reflect/#SliceHeader
	bh := [3]uintptr{sh[0], sh[1], sh[1]}
	b = *(*[]byte)(unsafe.Pointer(&bh))
	return
}
