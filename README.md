# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 


### Frameworks used for comparision 

- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Goji](https://github.com/zenazn/goji)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Martini](https://github.com/go-martini/martini)
- [Gin](https://github.com/gin-gonic/gin)


### Running the benchmarks

To run the benchmarks yourself simply clone this repository, step into the root directory and run:

```
 go test -benchmem -bench .
```
