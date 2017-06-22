package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bellavista/router"
	"github.com/bmizerany/pat"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/yarf-framework/yarf"
	"github.com/zenazn/goji/web"
)

// Yarf
type YarfParam struct {
	yarf.Resource
}

func (y *YarfParam) Get(c *yarf.Context) error {
	c.Render("Hello, " + c.Param("name"))

	return nil
}

// Bella Vista
func BVParam(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + router.GetString(r, "name1")))
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
	c.String(200, "Hello, "+c.Param("name"))
}

// Pat
func PatParam(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, " + req.URL.Query().Get(":name")))
}

func generateParamRequests(b *testing.B) ([]*httptest.ResponseRecorder, []*http.Request) {
	responses := make([]*httptest.ResponseRecorder, b.N)
	requests := make([]*http.Request, b.N)

	for i := 0; i < b.N; i++ {
		responses[i] = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "http://localhost:8080/hello/Joe", nil)
		if err != nil {
			b.Logf("got unexpected error: %q", err.Error())
			b.Fail()
		}
		requests[i] = req
	}

	return responses, requests
}

// Benchmarks

func BenchmarkParamBV(b *testing.B) {
	r := router.New("/")
	r.Add("/hello/:name", http.HandlerFunc(BVHello))
	d := router.Route(r)

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamYarfCached(b *testing.B) {
	y := yarf.New()
	y.Add("/hello/:name", new(YarfParam))

	responses, requests := generateParamRequests(b)

	// Warmup
	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/hello/:name", GinParam)

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamHttpRouter(b *testing.B) {
	router := httprouter.New()
	router.GET("/hello/:name", HttpRouterParam)

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamPat(b *testing.B) {
	m := pat.New()
	m.Get("/", http.HandlerFunc(PatParam))

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamGoji(b *testing.B) {
	g := web.New()
	g.Get("/hello/:name", GojiParam)

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamYarf(b *testing.B) {
	y := yarf.New()
	y.UseCache = false
	y.Add("/hello/:name", new(YarfParam))

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamMartini(b *testing.B) {
	r := martini.NewRouter()
	m := martini.New()
	r.Get("/hello/:name", MartiniParam)
	m.Action(r.Handle)

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkParamGorilla(b *testing.B) {
	m := mux.NewRouter()
	m.HandleFunc("/hello/{name}", GorillaParam)

	responses, requests := generateParamRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}
