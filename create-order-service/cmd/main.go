package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_status",
			Help: "HTTP request status codes",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(requestStatus)
}

func sendRequests(targetURL string, rate int, numOfProducts int) {
	client := &http.Client{}
	ticker := time.NewTicker(time.Second / time.Duration(rate))
	for range ticker.C {
		go func() {
			rand.Seed(time.Now().UnixNano())
			id := rand.Intn(numOfProducts) + 1
			resp, err := client.Get(targetURL + "?id=" + strconv.Itoa(id))
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				requestStatus.WithLabelValues("error").Inc()
				return
			}
			defer resp.Body.Close()
			requestStatus.WithLabelValues(resp.Status).Inc()
		}()
	}
}

func main() {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())

	requestsPerMoreThanSecond := 100 // Default rate
	targetURL := os.Getenv("TARGET_URL")
	numOfproducts, err := strconv.Atoi(os.Getenv("NUM_OF_PRODUCTS"))
	if err != nil {
		log.Fatal("Num of products should be integer")
	}
	go sendRequests(targetURL, requestsPerMoreThanSecond, numOfproducts)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":2112", nil))
}
