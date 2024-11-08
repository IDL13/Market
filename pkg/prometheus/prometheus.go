package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Listen(addr string) error {
	fmt.Println("[PROMETHEUS STARTED]")
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, mux)
}
