package main

import (
	"api-account/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mylibs/pkg/observability/motel"
	"net/http"
	"os"
	"time"
)

var client http.Client
var appname = os.Getenv("APP_NAME")
var traceEndpoint = os.Getenv("TRACE_ENDPOINT")
var hostEndpoint = os.Getenv("ENDPOINT")

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "code"},
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

		responseDuration.WithLabelValues(r.Method).Observe(duration)
	}()
	cus := &[]model.Customer{}
	log.Info().Msg("called")
	err := callGet[[]model.Customer](context.Background(), fmt.Sprintf("http://%s/customers", hostEndpoint), cus, "tid")
	a := model.Account{
		Id:        "1",
		Customers: *cus,
	}
	marshal, err := json.Marshal(a)

	if err != nil {
		log.Error().Err(err).Msg("error on customerx")
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
	client = motel.NewClient()
}

func main() {
	zerolog.TimestampFieldName = "date"
	zerolog.ErrorFieldName = "message"
	log.Logger = log.With().Str("application", appname).Logger()

	handler := http.HandlerFunc(httpHandler)
	http.Handle("/accounts", handler)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":3001", nil)
	panic(err)
}

func callGet[T any](ctx context.Context, url string, t *T, tid string) error {
	log.Info().Msgf("call GET:%s", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	res, err := client.Do(req)
	requestsTotal.WithLabelValues(req.Method, res.Status).Inc()
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	if res == nil {
		return errors.New("404 customer")
	}
	if res.StatusCode > 400 {
		log.Error().Err(err).Msg("error on customer")
		return err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	err = json.Unmarshal(resBody, t)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	return nil
}
