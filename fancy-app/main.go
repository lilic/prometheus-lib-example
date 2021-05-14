package main

import (
	"fmt"
	"net/http"

	"github.com/lilic/prometheus-lib-example/lib"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ()

func main() {
	// Create a new prometheus registry.
	registry := prometheus.NewRegistry()
	// Create our own component metric.
	counter := promauto.With(registry).NewCounterVec(prometheus.CounterOpts{
		Name: "test_total",
		Help: "test help",
	},
		[]string{"result"},
	)

	l := lib.New(registry)

	// As we never called any metrics related things
	// in the SomeLibFunc function we are not inheriting any of the library metrics.
	fmt.Println(l.OurFancyFunc())

	// Increment our own local metrics.
	counter.WithLabelValues("blah").Inc()

	// To handle and serve the metrics we pass in our new local registry.
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	fmt.Println("serving metrics on :2112")
	http.ListenAndServe(":2112", nil)
}
