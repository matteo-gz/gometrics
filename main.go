package main

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	requests.With("get", "/", "200", "OK").Inc()
	w.Write([]byte("Hello World!"))
}

func main() {
	prometheus.MustRegister(_metricRequests)
	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", hello)
	log.Println("listen:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
	}
}

var _metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: "server",
	Subsystem: "requests",
	Name:      "code_total",
	Help:      "The total number of processed requests",
}, []string{"kind", "operation", "code", "reason"})
var requests = prom.NewCounter(_metricRequests)
