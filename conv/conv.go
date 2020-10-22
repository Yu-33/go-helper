package conv

import "unsafe"

// BytesToString do unsafe convert byte array to string
// any change to b will affect s
func BytesToString(b []byte) (s string) {
	s = *(*string)(unsafe.Pointer(&b))
	return
}

// StringToBytes do unsafe convert string to byte array
// any change to b will affect s
func StringToBytes(s string) (b []byte) {
	// string's header, see https://golang.org/pkg/reflect/#StringHeader
	sh := (*[2]uintptr)(unsafe.Pointer(&s))
	// slice's header, see https://golang.org/pkg/reflect/#SliceHeader
	bh := [3]uintptr{sh[0], sh[1], sh[1]}
	b = *(*[]byte)(unsafe.Pointer(&bh))
	return
}
