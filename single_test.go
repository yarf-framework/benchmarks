package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yarf-framework/yarf"
)

// Yarf
type YarfSingle struct {
	yarf.Resource
}

func (y *YarfSingle) Get(c *yarf.Context) error {
	c.Render("Hello world!")

	return nil
}

func BenchmarkSingleYarf(b *testing.B) {
	y := yarf.New()
	y.Add("/", new(YarfSingle))

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(res, req)
	}
}
