# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 


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
 go test -benchmem -bench .
```


### Benchmark results (totally subjective to local hardware configuration)

```
BenchmarkMultiYarf-8             100      15852647 ns/op     4149288 B/op      90605 allocs/op
BenchmarkMultiHttpRouter-8       100      12663773 ns/op     4249047 B/op      80000 allocs/op
BenchmarkMultiGoji-8             100      18221827 ns/op     6330477 B/op      90004 allocs/op
BenchmarkMultiGorilla-8           30      43548624 ns/op    13142176 B/op     240015 allocs/op
BenchmarkMultiMartini-8           20      56270002 ns/op    16202178 B/op     220022 allocs/op
BenchmarkMultiGin-8              100      13578911 ns/op     3129951 B/op      80002 allocs/op
BenchmarkMultiPat-8              100      19447168 ns/op     5556713 B/op     140004 allocs/op
BenchmarkParamYarf-8         3000000           419 ns/op         153 B/op          3 allocs/op
BenchmarkParamHttpRouter-8  10000000           211 ns/op          79 B/op          2 allocs/op
BenchmarkParamGoji-8         2000000           624 ns/op         390 B/op          3 allocs/op
BenchmarkParamGorilla-8      1000000          2187 ns/op         886 B/op         12 allocs/op
BenchmarkParamMartini-8       300000          4096 ns/op        1265 B/op         16 allocs/op
BenchmarkParamGin-8          5000000           330 ns/op          63 B/op          2 allocs/op
BenchmarkParamPat-8          1000000          1059 ns/op         315 B/op          8 allocs/op
BenchmarkSimpleYarf-8        5000000           358 ns/op         159 B/op          3 allocs/op
BenchmarkSimpleHttpRouter-8 20000000            87.5 ns/op        47 B/op          1 allocs/op
BenchmarkSimpleGoji-8       10000000           224 ns/op          47 B/op          1 allocs/op
BenchmarkSimpleGorilla-8     1000000          1385 ns/op         535 B/op          9 allocs/op
BenchmarkSimpleMartini-8      500000          3393 ns/op         888 B/op         12 allocs/op
BenchmarkSimpleGin-8         5000000           237 ns/op          47 B/op          1 allocs/op
BenchmarkSimplePat-8        10000000           178 ns/op          95 B/op          2 allocs/op

```
