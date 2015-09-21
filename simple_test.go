package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/pat"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/yarf-framework/yarf"
	"github.com/zenazn/goji/web"
)

// Yarf
type YarfHello struct {
	yarf.Resource
}

func (y *YarfHello) Get() error {
	y.Render("Hello world!")

	return nil
}

// HttpRouter
func HttpRouterHello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello world!"))
}

// Goji
func GojiHello(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// Standard net/http handler
func HttpHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// Gin
func GinHello(c *gin.Context) {
	c.String(200, "Hello world!")
}

// Benchmarks
func BenchmarkSimpleYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/", new(YarfHello))

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(res, req)
	}
}

func BenchmarkSimpleHttpRouter(b *testing.B) {
	router := httprouter.New()
	router.GET("/", HttpRouterHello)

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(res, req)
	}
}

func BenchmarkSimpleGoji(b *testing.B) {
	g := web.New()
	g.Get("/", GojiHello)

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		g.ServeHTTP(res, req)
	}
}

func BenchmarkSimpleGorilla(b *testing.B) {
	m := mux.NewRouter()
	m.HandleFunc("/", HttpHello)

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		m.ServeHTTP(res, req)
	}
}

func BenchmarkSimpleMartini(b *testing.B) {
	r := martini.NewRouter()
	m := martini.New()
	r.Get("/", HttpHello)
	m.Action(r.Handle)

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		m.ServeHTTP(res, req)
	}
}

func BenchmarkSimpleGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/", GinHello)

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		r.ServeHTTP(res, req)
	}
}

func BenchmarkSimplePat(b *testing.B) {
	m := pat.New()
	m.Get("/", http.HandlerFunc(HttpHello))

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		m.ServeHTTP(res, req)
	}
}
