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
BenchmarkMultiYarf-8          500000          3595 ns/op        1260 B/op         15 allocs/op
BenchmarkMultiYarfCached-8   1000000          1702 ns/op         375 B/op          4 allocs/op
BenchmarkMultiHttpRouter-8   1000000          1595 ns/op         585 B/op          4 allocs/op
BenchmarkMultiGoji-8          500000          2233 ns/op         787 B/op          5 allocs/op
BenchmarkMultiGorilla-8       200000          6524 ns/op        1328 B/op         20 allocs/op
BenchmarkMultiMartini-8       200000          5846 ns/op        1376 B/op         16 allocs/op
BenchmarkMultiGin-8          2000000           987 ns/op         399 B/op          2 allocs/op
BenchmarkMultiPat-8          1000000          1478 ns/op         624 B/op          9 allocs/op
BenchmarkParamYarf-8         1000000          1583 ns/op         512 B/op          5 allocs/op
BenchmarkParamHttpRouter-8   1000000          1261 ns/op         416 B/op          4 allocs/op
BenchmarkParamGoji-8         1000000          1831 ns/op         720 B/op          5 allocs/op
BenchmarkParamGorilla-8       500000          3486 ns/op        1176 B/op         14 allocs/op
BenchmarkParamMartini-8       300000          4605 ns/op        1232 B/op         16 allocs/op
BenchmarkParamGin-8          2000000           571 ns/op         368 B/op          2 allocs/op
BenchmarkParamPat-8          1000000          1444 ns/op         624 B/op          9 allocs/op
BenchmarkSimpleYarf-8        1000000          1476 ns/op         512 B/op          5 allocs/op
BenchmarkSimpleHttpRouter-8  1000000          1143 ns/op         384 B/op          3 allocs/op
BenchmarkSimpleGoji-8        1000000          1311 ns/op         384 B/op          3 allocs/op
BenchmarkSimpleGorilla-8      500000          2650 ns/op         848 B/op         11 allocs/op
BenchmarkSimpleMartini-8      300000          3858 ns/op         848 B/op         12 allocs/op
BenchmarkSimpleGin-8         3000000           633 ns/op         352 B/op          1 allocs/op
BenchmarkSimplePat-8         1000000          1196 ns/op         432 B/op          4 allocs/op
```
