package container

import "bytes"

type Comparer interface {
	// Compare comparing with the target Comparator.
	// return -1 if current < target, 0 if current == target, 1 if current > target.
	Compare(target Comparer) int
}

type String string

func (k1 String) Compare(target Comparer) int {
	k2 := target.(String)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Byte byte

func (k1 Byte) Compare(target Comparer) int {
	k2 := target.(Byte)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Bytes []byte

func (k1 Bytes) Compare(target Comparer) int {
	k2 := target.(Bytes)
	return bytes.Compare(k1, k2)
}

type Int int

func (k1 Int) Compare(target Comparer) int {
	k2 := target.(Int)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Int8 int8

func (k1 Int8) Compare(target Comparer) int {
	k2 := target.(Int8)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Int16 int16

func (k1 Int16) Compare(target Comparer) int {
	k2 := target.(Int16)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Int32 int32

func (k1 Int32) Compare(target Comparer) int {
	k2 := target.(Int32)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Int64 int64

func (k1 Int64) Compare(target Comparer) int {
	k2 := target.(Int64)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Uint uint

func (k1 Uint) Compare(target Comparer) int {
	k2 := target.(Uint)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Uint8 uint8

func (k1 Uint8) Compare(target Comparer) int {
	k2 := target.(Uint8)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Uint16 uint16

func (k1 Uint16) Compare(target Comparer) int {
	k2 := target.(Uint16)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Uint32 uint32

func (k1 Uint32) Compare(target Comparer) int {
	k2 := target.(Uint32)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}

type Uint64 uint64

func (k1 Uint64) Compare(target Comparer) int {
	k2 := target.(Uint64)
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}
