# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 


### Frameworks used for comparision 

- [HttpRouter](github.com/julienschmidt/httprouter)
- [Goji](github.com/zenazn/goji)
- [Gorilla Mux](github.com/gorilla/mux)
- [Martini](github.com/go-martini/martini)
- [Gin](github.com/gin-gonic/gin)


### Running the benchmarks

To run the benchmarks yourself simply clone this repository, step into the root directory and run:

```
 go test -benchmem -bench .
```
