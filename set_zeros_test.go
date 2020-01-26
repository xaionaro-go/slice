package slice

import (
	"fmt"
	"testing"
)

var (
	implementations = map[string]func([]byte){
		"naive": setZerosNaive,
		"copy": setZerosViaCopy,
	}
)

func TestSetZeros(t *testing.T) {
	sliceLens := []int{
		0, 1, 2, 7, 8, 15, 16, 17, 65535, 65536, 65538,
	}

	for implName, implFn := range implementations {
		for _, sliceLen := range sliceLens {
			s := make([]byte, sliceLen+1)
			for idx := range s {
				s[idx] = 1
			}
			fn := func(t *testing.T) {
				implFn(s[:sliceLen])
				for idx := 0; idx<sliceLen; idx++ {
					if s[idx] != 0 {
						t.Errorf("s[%d] == %d\n", idx, s[idx])
					}
				}
				if s[sliceLen] != 1 {
					t.Errorf("s[%d] == %d", sliceLen, s[sliceLen])
				}
			}
			t.Run(fmt.Sprintf("TestSetZeros_%s_%d", implName, len(s)), fn)
		}
	}
}

func BenchmarkSetZeros(b *testing.B) {
	slices := [][]byte{
		make([]byte, 0),
		make([]byte, 1),
		make([]byte, 15),
		make([]byte, 16),
		make([]byte, 256),
		make([]byte, 65536),
	}

	for implName, implFn := range implementations {
		for _, s := range slices {
			fn := func(b *testing.B) {
				b.SetBytes(int64(len(s)))
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					implFn(s)
				}
			}
			b.Run(fmt.Sprintf("BenchmarkSetZeros_%s_%d", implName, len(s)), fn)
		}
	}
}
