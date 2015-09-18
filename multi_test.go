package benchmarks

import (
    "testing"
    "net/url"
	"net/http"
	"net/http/httptest"
	"strconv"
	
	"github.com/yarf-framework/yarf"
	"github.com/julienschmidt/httprouter"
    "github.com/zenazn/goji/web"
    "github.com/gorilla/mux"
    "github.com/go-martini/martini"
    "github.com/gin-gonic/gin"
)

// Yarf
type YarfMulti struct {
	yarf.Resource
}

func (y *YarfMulti) Get() error {
	y.Render("Hello " + y.Param("name1") + "," + y.Param("name2") + "," + y.Param("name3") + "," + y.Param("name4"))

	return nil
}

// HttpRouter
func HttpRouterMulti(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Write([]byte("Hello, " + ps.ByName("name1") + "," + ps.ByName("name2") + "," + ps.ByName("name3") + "," + ps.ByName("name4")))
}

// Goji
func GojiMulti(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, " + c.URLParams["name1"] + "," + c.URLParams["name2"] + "," + c.URLParams["name3"] + "," + c.URLParams["name4"]))
}

// Gorilla
func GorillaMulti(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, " + mux.Vars(r)["name1"] + "," + mux.Vars(r)["name2"] + "," + mux.Vars(r)["name3"] + "," + mux.Vars(r)["name4"]))
}

// Martini
func MartiniMulti(params martini.Params) string {
    return "Hello, " + params["name1"] + "," + params["name2"] + "," + params["name3"] + "," + params["name4"]
}

// Gin
func GinMulti(c *gin.Context) {
    c.String(200, "Hello, " + c.Param("name1") + "," + c.Param("name2") + "," + c.Param("name3") + "," + c.Param("name4"))
}

// Benchmarks
func BenchmarkMultiYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/hello/:name1/:name2/:name3/:name4", new(YarfMulti))

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/", nil)
	res := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
        		        y.ServeHTTP(res, req)
		            }
		        }
		    }
		}
	}
}

func BenchmarkMultiHttpRouter(b *testing.B) {
    router := httprouter.New()
    router.GET("/hello/:name1/:name2/:name3/:name4", HttpRouterMulti)

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
		                router.ServeHTTP(res, req)
	                }
	            }
		    }
		}
	}
}

func BenchmarkMultiGoji(b *testing.B) {
    g := web.New()
    g.Get("/hello/:name1/:name2/:name3/:name4", GojiMulti)

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
		                g.ServeHTTP(res, req)
	                }
	            }
		    }
		}
	}
}

func BenchmarkMultiGorilla(b *testing.B) {
    m := mux.NewRouter()
	m.HandleFunc("/hello/{name1}/{name2}/{name3}/{name4}", GorillaMulti)
	
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
		                m.ServeHTTP(res, req)
	                }
	            }
		    }
		}
	}
}

func BenchmarkMultiMartini(b *testing.B) {
    r := martini.NewRouter()
    m := martini.New()
	r.Get("/hello/:name1/:name2/:name3/:name4", MartiniMulti)
	m.Action(r.Handle)
	
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
		                m.ServeHTTP(res, req)
	                }
	            }
		    }
		}
	}
}

func BenchmarkMultiGin(b *testing.B) {
    gin.SetMode(gin.ReleaseMode)
    
	r := gin.New()
    r.GET("/hello/:name1/:name2/:name3/:name4", GinMulti)
    
    req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res := httptest.NewRecorder()
	
	// Run benchmark
	for i := 0; i < b.N; i++ {
		for a := 0; a < 10; a++ {
		    for b := 0; b < 10; b++ {
		        for c := 0; c < 10; c++ {
		            for d := 0; d < 10; d++ {
		                req.URL, _ = url.Parse("http://localhost:8080/hello/" + strconv.Itoa(a) + "/" + strconv.Itoa(b) + "/" + strconv.Itoa(c) + "/" + strconv.Itoa(d))
		                r.ServeHTTP(res, req)
	                }
	            }
		    }
		}
	}
}
