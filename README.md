# SetZeros

It appears the naive implementation
```go
for idx := range dst {
	dst[idx] = 0
}
```

is the fastest one. Try also run the benchmark test of
[github.com/tmthrgd/go-memset](https://github.com/tmthrgd/go-memset).
It shows the same performance for Memset and GoZero on my laptop.

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/slice
BenchmarkSetZeros/BenchmarkSetZeros_words_0-8         	356671656	         3.36 ns/op
BenchmarkSetZeros/BenchmarkSetZeros_words_1-8         	380630080	         3.13 ns/op	 319.06 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_words_15-8        	275007468	         4.32 ns/op	3476.23 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_words_16-8        	231319143	         4.87 ns/op	3287.16 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_words_256-8       	43794810	        26.8 ns/op	9542.79 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_words_65536-8     	  457748	      2667 ns/op	24570.79 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_naive_0-8         	427999196	         2.83 ns/op
BenchmarkSetZeros/BenchmarkSetZeros_naive_1-8         	278331762	         4.54 ns/op	 220.44 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_naive_15-8        	247765914	         4.81 ns/op	3115.98 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_naive_16-8        	250207819	         6.30 ns/op	2541.02 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_naive_256-8       	121795636	         9.75 ns/op	26261.91 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_naive_65536-8     	  672644	      1783 ns/op	36755.97 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_copy_0-8          	285977827	         4.15 ns/op
BenchmarkSetZeros/BenchmarkSetZeros_copy_1-8          	219969361	         5.48 ns/op	 182.40 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_copy_15-8         	222189085	         5.40 ns/op	2780.06 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_copy_16-8         	205798772	         5.44 ns/op	2939.77 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_copy_256-8        	95995905	        12.3 ns/op	20816.85 MB/s
BenchmarkSetZeros/BenchmarkSetZeros_copy_65536-8      	  637640	      1775 ns/op	36930.88 MB/s
PASS
ok  	github.com/xaionaro-go/slice	28.301s
```

So `SetZeros` is just the naive implementation.