# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 
The results shows how using the route cache improves the performance beating any other framework in the test.

Note: Running the benchmarks will need at least 14GB of free memory.

### Frameworks used for comparison 

- [Gin](https://github.com/gin-gonic/gin)
- [Goji](https://github.com/zenazn/goji)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Martini](https://github.com/go-martini/martini)
- [Pat](https://github.com/bmizerany/pat)
- [Yarf](https://github.com/yarf-framework/yarf)


### Running the benchmarks

To run the benchmarks yourself simply clone this repository, step into the root directory and run:

```
 # CPU profiling
 go test -bench .

 # Memory profiling
 go test -benchmem -bench .
```


### Benchmark results (totally subjective to local hardware configuration)

```
BenchmarkMultiYarf-8              500000          3712 ns/op        1532 B/op         13 allocs/op
BenchmarkMultiYarfCached-8       2000000           751 ns/op         128 B/op          3 allocs/op
BenchmarkMultiHttpRouter-8       1000000          1811 ns/op        1001 B/op          7 allocs/op
BenchmarkMultiGoji-8              500000          2654 ns/op        1203 B/op          8 allocs/op
BenchmarkMultiGorilla-8           200000          6428 ns/op        1744 B/op         23 allocs/op
BenchmarkMultiMartini-8           200000          6144 ns/op        1424 B/op         17 allocs/op
BenchmarkMultiGin-8              1000000          1258 ns/op         814 B/op          5 allocs/op
BenchmarkMultiPat-8              1000000          1923 ns/op        1056 B/op         13 allocs/op
BenchmarkParamYarf-8             1000000          2405 ns/op        1312 B/op         13 allocs/op
BenchmarkParamYarfCached-8       3000000           518 ns/op         120 B/op          3 allocs/op
BenchmarkParamHttpRouter-8       1000000          1692 ns/op         832 B/op          7 allocs/op
BenchmarkParamGoji-8             1000000          2182 ns/op        1136 B/op          8 allocs/op
BenchmarkParamGorilla-8           500000          3788 ns/op        1592 B/op         17 allocs/op
BenchmarkParamMartini-8           200000          5308 ns/op        1280 B/op         17 allocs/op
BenchmarkParamGin-8              2000000          1392 ns/op         784 B/op          5 allocs/op
BenchmarkParamPat-8              1000000          2148 ns/op        1056 B/op         13 allocs/op
BenchmarkSimpleYarf-8            3000000           528 ns/op         176 B/op          5 allocs/op
BenchmarkSimpleYarfCached-8       500000          3086 ns/op        1568 B/op         14 allocs/op
BenchmarkSimpleHttpRouter-8      1000000          1678 ns/op         800 B/op          6 allocs/op
BenchmarkSimpleGoji-8            1000000          1771 ns/op         800 B/op          6 allocs/op
BenchmarkSimpleGorilla-8          500000          3128 ns/op        1264 B/op         14 allocs/op
BenchmarkSimpleMartini-8          300000          5161 ns/op         896 B/op         13 allocs/op
BenchmarkSimpleGin-8             2000000          1176 ns/op         768 B/op          4 allocs/op
BenchmarkSimplePat-8             1000000          1645 ns/op         848 B/op          7 allocs/op
```
