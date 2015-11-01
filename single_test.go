package benchmarks

import (
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

	responses, requests := generateSimpleRequests(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.ServeHTTP(responses[i], requests[i])
	}
}
