package main

import (
	"api-customer/internal/service"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

var appname = os.Getenv("APP_NAME")
var traceEndpoint = os.Getenv("TRACE_ENDPOINT")
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method"},
	)
	responseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Histogram of response latencies for HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

var httpHandler = func(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		requestsTotal.WithLabelValues(r.Method).Inc()
		responseDuration.WithLabelValues(r.Method).Observe(duration)
	}()
	tid := r.Header.Get("tid")
	lc := service.GetCustomersAll()
	log.Info().Str("tid", tid).Msg("finding all customers")
	marshal, err := json.Marshal(lc)
	if err != nil {
		log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		log.Info().Msg(string(marshal))
		w.WriteHeader(http.StatusOK)
		w.Write(marshal)
	}
}

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseDuration)
}

func main() {

	//logs
	zerolog.TimestampFieldName = "date"
	zerolog.ErrorFieldName = "message"
	log.Logger = log.With().Str("application", appname).Logger()

	handler := http.HandlerFunc(httpHandler)
	http.Handle("/customers", handler)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":3000", nil)
	panic(err)
}
