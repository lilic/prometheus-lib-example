package lib

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SomeDefaultValue = "blah"

type FancyLib struct {
	Blah    string
	metrics *metrics
}

type metrics struct {
	someTotal *prometheus.CounterVec
}

func newMetrics(registry prometheus.Registerer) *metrics {
	return &metrics{
		someTotal: promauto.With(registry).NewCounterVec(
			prometheus.CounterOpts{
				Name: "lib_total",
				Help: "Total number of HTTP requests by status code and method.",
			},
			[]string{"label"},
		),
	}
}

func New(registry prometheus.Registerer) *FancyLib {
	var m *metrics
	// This allows us to test this function
	// as in tests we don't necessary pass
	// a new metrics registry.
	if registry != nil {
		m = newMetrics(registry)
	}

	return &FancyLib{
		Blah:    SomeDefaultValue,
		metrics: m,
	}
}

func (f *FancyLib) OurFancyFunc() string {
	// Lets do some stuff.
	fmt.Println("you called our fancy lib func")
	time.Sleep(5 * time.Second)

	// Inc the metric.
	f.metrics.someTotal.WithLabelValues("foo").Inc()

	f.Blah = "yay"

	return fmt.Sprintf("we did not call NewMetrics as we are library: %s", f.Blah)
}
