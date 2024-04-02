package main

import (
	"apppuc/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mylibs/pkg/observability/motel"
	"mylibs/pkg/util"
	"net/http"
	"os"
	"sync"
	"time"
)

var client http.Client

var appname = os.Getenv("APP_NAME")
var traceEndpoint = os.Getenv("TRACE_ENDPOINT")
var hostEndpoint = os.Getenv("ENDPOINT")
var productEndpoint = os.Getenv("PRODUCT_ENDPOINT")

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_error",
			Help: "Total number of error HTTP requests.",
		},
		[]string{"status_code"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	zerolog.TimestampFieldName = "date"
	zerolog.ErrorFieldName = "message"
	client = motel.NewClient()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the WaitGroup

	// Goroutine 1
	go func() {
		defer wg.Done() // Notify WaitGroup when this goroutine completes
		// Your logic for goroutine 1 here
		fmt.Println("Goroutine 1 started")
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":3002", nil)
		panic(err)
	}()

	// Goroutine 2
	go func() {
		defer wg.Done() // Notify WaitGroup when this goroutine completes
		// Your logic for goroutine 2 here
		//trace
		defer func() {

			log.Info().Msg("gateway init")
			for true {
				err := server(context.Background())
				if err != nil {
					requestsTotal.WithLabelValues(err.Code).Inc()
					log.Error().Err(err).Msg("")
				}
				time.Sleep(1 * time.Second)
				//time.Sleep((60 * time.Second) * 10)
			}
		}()
		fmt.Println("Goroutine 2 started")
	}()

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("Both goroutines have completed")

}

func server(ctx context.Context) *HttpError {
	tid := util.RandStringBytes(18)
	ps := &[]model.Product{}
	err = callGet[[]model.Product](ctx, fmt.Sprintf("http://%s/products/1", productEndpoint), ps, tid)
	if err != nil {
		return err
	}
	return nil
}

func callGet[T any](ctx context.Context, url string, t *T, tid string) *HttpError {
	log.Info().Msgf("call GET:%s", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("")
		return BuildError("500", err)
	}
	if res.StatusCode > 400 {
		log.Error().Err(err).Msg("")
		return BuildError(res.Status, err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("")
		return BuildError("400", err)
	}
	err = json.Unmarshal(resBody, t)
	if err != nil {
		log.Error().Err(err).Msg("")
		return BuildError("400", err)
	}
	log.Info().Msgf("success call on %s", url)
	return nil
}

type HttpError struct {
	Code string `json:"code"`
	E    error  `json:"e"`
}

func (h HttpError) Error() string {
	return h.Code
}

func BuildError(code string, err error) *HttpError {
	return &HttpError{code, err}
}
