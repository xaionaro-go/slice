// +build amd64

package slice

import (
	"reflect"
	"unsafe"
)

func setZerosWords(dst []byte) {
	// slower than even setZerosNaive :(

	ptr := (*reflect.SliceHeader)(unsafe.Pointer(&dst)).Data
	eptr := ptr + uintptr(len(dst) & ^0x07)
	for ; ptr < eptr; {
		*(*uint64)(unsafe.Pointer(ptr)) = 0
		ptr += 8
	}
	restLen := len(dst) & 0x07
	if restLen & 0x04 != 0 {
		*(*uint32)(unsafe.Pointer(ptr)) = 0
		ptr += 4
	}
	if restLen & 0x02 != 0 {
		*(*uint16)(unsafe.Pointer(ptr)) = 0
		ptr += 2
	}
	if restLen & 0x01 != 0 {
		*(*uint8)(unsafe.Pointer(ptr)) = 0
		ptr += 1
	}
}

