# YARF benchmarks

This are several benchmarks made to see how YARF performs under different conditions and compared against other similar frameworks. 


### Frameworks used for comparision 

- github.com/julienschmidt/httprouter
- github.com/zenazn/goji
- github.com/gorilla/mux
- github.com/go-martini/martini
- github.com/gin-gonic/gin


### Running the benchmarks

To run the benchmarks yourself simply clone this repository, step into the root directory and run:

```
 go test -benchmem -bench .
```
