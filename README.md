# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 

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
BenchmarkMultiYarf-8          500000          2539 ns/op         892 B/op         13 allocs/op
BenchmarkMultiYarfCached-8   1000000          1479 ns/op         375 B/op          4 allocs/op
BenchmarkMultiHttpRouter-8   3000000           873 ns/op         221 B/op          2 allocs/op
BenchmarkMultiGoji-8         1000000          1413 ns/op         425 B/op          3 allocs/op
BenchmarkMultiGorilla-8       200000          5984 ns/op        1088 B/op         18 allocs/op
BenchmarkMultiMartini-8       200000          6370 ns/op        1376 B/op         16 allocs/op
BenchmarkMultiGin-8          2000000          1039 ns/op         446 B/op          3 allocs/op
BenchmarkMultiPat-8          1000000          1504 ns/op         624 B/op          9 allocs/op
BenchmarkParamYarf-8         3000000           537 ns/op         144 B/op          3 allocs/op
BenchmarkParamHttpRouter-8   5000000           305 ns/op          48 B/op          2 allocs/op
BenchmarkParamGoji-8         2000000           793 ns/op         352 B/op          3 allocs/op
BenchmarkParamGorilla-8       500000          2733 ns/op         848 B/op         12 allocs/op
BenchmarkParamMartini-8       300000          5295 ns/op        1232 B/op         16 allocs/op
BenchmarkParamGin-8          2000000           671 ns/op         384 B/op          3 allocs/op
BenchmarkParamPat-8          1000000          1657 ns/op         624 B/op          9 allocs/op
BenchmarkSimpleYarf-8        3000000           464 ns/op         144 B/op          3 allocs/op
BenchmarkSimpleHttpRouter-8 10000000           168 ns/op          16 B/op          1 allocs/op
BenchmarkSimpleGoji-8        5000000           316 ns/op          16 B/op          1 allocs/op
BenchmarkSimpleGorilla-8     1000000          1644 ns/op         496 B/op          9 allocs/op
BenchmarkSimpleMartini-8      300000          4231 ns/op         848 B/op         12 allocs/op
BenchmarkSimpleGin-8         2000000           599 ns/op         368 B/op          2 allocs/op
BenchmarkSimplePat-8         5000000           245 ns/op          64 B/op          2 allocs/op
```
