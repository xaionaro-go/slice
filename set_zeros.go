package slice

func setZerosNaive(dst []byte) {
	for idx := range dst {
		dst[idx] = 0
	}
}

var zeros = make([]byte, 1<<20)
func setZerosViaCopy(dst []byte) {
	if len(dst) > len(zeros) {
		panic("should not happen")
	}
	copy(dst, zeros)
}

func SetZeros(dst []byte) {
	setZerosNaive(dst)
}