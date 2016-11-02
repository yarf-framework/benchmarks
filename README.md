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
BenchmarkMultiYarf-8              500000          3949 ns/op        1564 B/op         14 allocs/op
BenchmarkMultiYarfCached-8       1000000          1552 ns/op         375 B/op          4 allocs/op
BenchmarkMultiHttpRouter-8       1000000          1968 ns/op        1001 B/op          7 allocs/op
BenchmarkMultiGoji-8              500000          2669 ns/op        1203 B/op          8 allocs/op
BenchmarkMultiGorilla-8           200000          6686 ns/op        1744 B/op         23 allocs/op
BenchmarkMultiMartini-8           200000          6947 ns/op        1424 B/op         17 allocs/op
BenchmarkMultiGin-8              1000000          1381 ns/op         814 B/op          5 allocs/op
BenchmarkMultiPat-8              1000000          2061 ns/op        1056 B/op         13 allocs/op
BenchmarkParamYarf-8              500000          2685 ns/op        1344 B/op         14 allocs/op
BenchmarkParamYarfCached-8       3000000           646 ns/op         144 B/op          3 allocs/op
BenchmarkParamHttpRouter-8       1000000          1835 ns/op         832 B/op          7 allocs/op
BenchmarkParamGoji-8             1000000          2206 ns/op        1136 B/op          8 allocs/op
BenchmarkParamGorilla-8           500000          3995 ns/op        1592 B/op         17 allocs/op
BenchmarkParamMartini-8           200000          5492 ns/op        1280 B/op         17 allocs/op
BenchmarkParamGin-8              2000000          1446 ns/op         784 B/op          5 allocs/op
BenchmarkParamPat-8              1000000          2132 ns/op        1056 B/op         13 allocs/op
BenchmarkSimpleYarf-8            1000000          2072 ns/op         992 B/op         11 allocs/op
BenchmarkSimpleYarfCached-8      3000000           453 ns/op         144 B/op          3 allocs/op
BenchmarkSimpleHttpRouter-8      1000000          1702 ns/op         800 B/op          6 allocs/op
BenchmarkSimpleGoji-8            1000000          1779 ns/op         800 B/op          6 allocs/op
BenchmarkSimpleGorilla-8          500000          3087 ns/op        1264 B/op         14 allocs/op
BenchmarkSimpleMartini-8          200000          5354 ns/op         896 B/op         13 allocs/op
BenchmarkSimpleGin-8             2000000          1205 ns/op         768 B/op          4 allocs/op
BenchmarkSimplePat-8             1000000          1732 ns/op         848 B/op          7 allocs/op
```
