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
BenchmarkMultiYarf-8             100      15034192 ns/op     3829325 B/op      90606 allocs/op
BenchmarkMultiHttpRouter-8       100      12624134 ns/op     4249046 B/op      80000 allocs/op
BenchmarkMultiGoji-8             100      17996501 ns/op     6330484 B/op      90004 allocs/op
BenchmarkMultiGorilla-8           30      43971258 ns/op    13142188 B/op     240015 allocs/op
BenchmarkMultiMartini-8           20      55866361 ns/op    16202111 B/op     220022 allocs/op
BenchmarkMultiGin-8              100      13654760 ns/op     3129934 B/op      80002 allocs/op
BenchmarkMultiPat-8              100      19181758 ns/op     5556737 B/op     140005 allocs/op
BenchmarkParamYarf-8         5000000           295 ns/op         127 B/op          3 allocs/op
BenchmarkParamHttpRouter-8  10000000           212 ns/op          79 B/op          2 allocs/op
BenchmarkParamGoji-8         2000000           625 ns/op         390 B/op          3 allocs/op
BenchmarkParamGorilla-8      1000000          2210 ns/op         886 B/op         12 allocs/op
BenchmarkParamMartini-8       300000          4076 ns/op        1265 B/op         16 allocs/op
BenchmarkParamGin-8          5000000           337 ns/op          63 B/op          2 allocs/op
BenchmarkParamPat-8          1000000          1079 ns/op         315 B/op          8 allocs/op
BenchmarkSimpleYarf-8        5000000           234 ns/op         127 B/op          3 allocs/op
BenchmarkSimpleHttpRouter-8 20000000            89.6 ns/op        47 B/op          1 allocs/op
BenchmarkSimpleGoji-8        5000000           230 ns/op          47 B/op          1 allocs/op
BenchmarkSimpleGorilla-8     1000000          1395 ns/op         535 B/op          9 allocs/op
BenchmarkSimpleMartini-8      500000          3375 ns/op         888 B/op         12 allocs/op
BenchmarkSimpleGin-8         5000000           240 ns/op          47 B/op          1 allocs/op
BenchmarkSimplePat-8        10000000           177 ns/op          95 B/op          2 allocs/op

```
