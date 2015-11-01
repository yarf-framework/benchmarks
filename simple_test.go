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

func (y *YarfHello) Get(c *yarf.Context) error {
	c.Render("Hello world!")

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

func generateSimpleRequests(b *testing.B) ([]*httptest.ResponseRecorder, []*http.Request) {
	responses := make([]*httptest.ResponseRecorder, b.N)
	requests := make([]*http.Request, b.N)

	for i := 0; i < b.N; i++ {
		responses[i] = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		if err != nil {
			b.Logf("got unexpected error: %q", err.Error())
			b.Fail()
		}
		requests[i] = req
	}

	return responses, requests
}

// Benchmarks
func BenchmarkSimpleYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/", new(YarfHello))

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimpleHttpRouter(b *testing.B) {
	router := httprouter.New()
	router.GET("/", HttpRouterHello)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimpleGoji(b *testing.B) {
	g := web.New()
	g.Get("/", GojiHello)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimpleGorilla(b *testing.B) {
	m := mux.NewRouter()
	m.HandleFunc("/", HttpHello)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimpleMartini(b *testing.B) {
	r := martini.NewRouter()
	m := martini.New()
	r.Get("/", HttpHello)
	m.Action(r.Handle)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimpleGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/", GinHello)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkSimplePat(b *testing.B) {
	m := pat.New()
	m.Get("/", http.HandlerFunc(HttpHello))

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}
