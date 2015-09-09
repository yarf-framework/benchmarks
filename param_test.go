package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/yarf-framework/yarf"
	"github.com/julienschmidt/httprouter"
    "github.com/zenazn/goji/web"
    "github.com/gorilla/mux"
    "github.com/go-martini/martini"
    "github.com/gin-gonic/gin"
)

// Yarf
type YarfParam struct {
	yarf.Resource
}

func (y *YarfParam) Get() error {
	y.Render("Hello, " + y.Param("name"))

	return nil
}

// HttpRouter
func HttpRouterParam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Write([]byte("Hello, " + ps.ByName("name")))
}

// Goji
func GojiParam(c web.C, w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, " + c.URLParams["name"]))
}

// Gorilla
func GorillaParam(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, " + mux.Vars(r)["name"]))
}

// Martini
func MartiniParam(params martini.Params) string {
    return "Hello, " + params["name"]
}

// Gin
func GinParam(c *gin.Context) {
    c.String(200, "Hello, " + c.Param("name"))
}

// Benchmarks
func BenchmarkParamYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/hello/:name", new(YarfParam))

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(res, req)
	}
}

func BenchmarkParamHttpRouter(b *testing.B) {
    router := httprouter.New()
    router.GET("/hello/:name", HttpRouterParam)

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
	    router.ServeHTTP(res, req)
	}
}

func BenchmarkParamGoji(b *testing.B) {
    g := web.New()
    g.Get("/hello/:name", GojiParam)

	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
	    g.ServeHTTP(res, req)
	}
}

func BenchmarkParamGorilla(b *testing.B) {
    m := mux.NewRouter()
	m.HandleFunc("/hello/{name}", GorillaParam)
	
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
	    m.ServeHTTP(res, req)
	}
}

func BenchmarkParamMartini(b *testing.B) {
    r := martini.NewRouter()
    m := martini.New()
	r.Get("/hello/:name", MartiniParam)
	m.Action(r.Handle)
	
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
	    m.ServeHTTP(res, req)
	}
}

func BenchmarkParamGin(b *testing.B) {
    gin.SetMode(gin.ReleaseMode)
    
	r := gin.New()
    r.GET("/hello/:name", GinParam)
    
    req, _ := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
	res := httptest.NewRecorder()
	
	// Run benchmark
	for i := 0; i < b.N; i++ {
	    r.ServeHTTP(res, req)
	}
}
