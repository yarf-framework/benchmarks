# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 
The results shows how using the route cache improves the performance beating most frameworks in the test.

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
BenchmarkMultiYarfCached-8           1000000          1171 ns/op         375 B/op          4 allocs/op
BenchmarkMultiGin-8                  1000000          1274 ns/op         814 B/op          5 allocs/op
BenchmarkMultiHttpRouter-8           1000000          1814 ns/op        1001 B/op          7 allocs/op
BenchmarkMultiPat-8                  1000000          1799 ns/op        1056 B/op         13 allocs/op
BenchmarkMultiGoji-8                 1000000          2276 ns/op        1209 B/op          8 allocs/op
BenchmarkMultiYarf-8                  500000          2977 ns/op        1564 B/op         14 allocs/op
BenchmarkMultiGorilla-8               300000          6270 ns/op        1834 B/op         24 allocs/op
BenchmarkMultiMartini-8               200000          6428 ns/op        1424 B/op         17 allocs/op
BenchmarkParamYarfCached-8           3000000           412 ns/op         144 B/op          3 allocs/op
BenchmarkParamGin-8                  1000000          1309 ns/op         784 B/op          5 allocs/op
BenchmarkParamHttpRouter-8           1000000          1547 ns/op         832 B/op          7 allocs/op
BenchmarkParamPat-8                  1000000          1921 ns/op        1056 B/op         13 allocs/op
BenchmarkParamGoji-8                 1000000          1884 ns/op        1136 B/op          8 allocs/op
BenchmarkParamYarf-8                  500000          2546 ns/op        1328 B/op         13 allocs/op
BenchmarkParamMartini-8               300000          5092 ns/op        1248 B/op         16 allocs/op
BenchmarkParamGorilla-8               500000          3457 ns/op        1592 B/op         17 allocs/op
BenchmarkSimpleYarfCached-8          3000000           405 ns/op         144 B/op          3 allocs/op
BenchmarkSimpleGin-8                 1000000          1069 ns/op         768 B/op          4 allocs/op
BenchmarkSimpleHttpRouter-8          1000000          1694 ns/op         800 B/op          6 allocs/op
BenchmarkSimplePat-8                 1000000          1815 ns/op         848 B/op          7 allocs/op
BenchmarkSimpleGoji-8                1000000          1819 ns/op         800 B/op          6 allocs/op
BenchmarkSimpleYarf-8                1000000          2073 ns/op         992 B/op         11 allocs/op
BenchmarkSimpleMartini-8              300000          5466 ns/op         880 B/op         12 allocs/op
BenchmarkSimpleGorilla-8              500000          3087 ns/op        1264 B/op         14 allocs/op
```
