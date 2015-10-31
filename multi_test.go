package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"strconv"
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
type YarfMulti struct {
	yarf.Resource
}

func (y *YarfMulti) Get(c *yarf.Context) error {
	c.Render("Hello " + c.Param("name1") + "," + c.Param("name2") + "," + c.Param("name3") + "," + c.Param("name4"))

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
	c.String(200, "Hello, "+c.Param("name1")+","+c.Param("name2")+","+c.Param("name3")+","+c.Param("name4"))
}

// Pat
func PatMulti(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, " + req.URL.Query().Get(":name1") + req.URL.Query().Get(":name2") + req.URL.Query().Get(":name3") + req.URL.Query().Get(":name4")))
}

func generateMultiRequests(b *testing.B) ([]*httptest.ResponseRecorder, []*http.Request) {
	responses := make([]*httptest.ResponseRecorder, b.N)
	requests := make([]*http.Request, b.N)

	for i := 0; i < b.N; i++ {
		responses[i] = httptest.NewRecorder()

		req, err := http.NewRequest("GET", "http://localhost:8080/hello/"+strconv.Itoa(i)+"/"+strconv.Itoa(i)+"/"+strconv.Itoa(i)+"/"+strconv.Itoa(i), nil)
		if err != nil {
			b.Logf("got unexpected error: %q", err.Error())
			b.Fail()
		}
		requests[i] = req
	}

	return responses, requests
}

// Benchmarks
func BenchmarkMultiYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/hello/:name1/:name2/:name3/:name4", new(YarfMulti))

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiHttpRouter(b *testing.B) {
	router := httprouter.New()
	router.GET("/hello/:name1/:name2/:name3/:name4", HttpRouterMulti)

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiGoji(b *testing.B) {
	g := web.New()
	g.Get("/hello/:name1/:name2/:name3/:name4", GojiMulti)

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiGorilla(b *testing.B) {
	m := mux.NewRouter()
	m.HandleFunc("/hello/{name1}/{name2}/{name3}/{name4}", GorillaMulti)

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiMartini(b *testing.B) {
	r := martini.NewRouter()
	m := martini.New()
	r.Get("/hello/:name1/:name2/:name3/:name4", MartiniMulti)
	m.Action(r.Handle)

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/hello/:name1/:name2/:name3/:name4", GinMulti)

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.ServeHTTP(responses[i], requests[i])
	}
}

func BenchmarkMultiPat(b *testing.B) {
	m := pat.New()
	m.Get("/", http.HandlerFunc(PatMulti))

	responses, requests := generateMultiRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.ServeHTTP(responses[i], requests[i])
	}
}
