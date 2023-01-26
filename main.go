package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/clockworksoul/mediawiki"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metrics *mediawiki.SiteinfoResponse

func getSiteinfo(url string) (*mediawiki.SiteinfoResponse, error) {
	c, err := mediawiki.New(url, "mediawiki-exporter")
	if err != nil {
		return nil, fmt.Errorf("failed to build mediawiki client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	r, err := c.Siteinfo().Prop("statistics").Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve mediawiki statistics")
	}

	return &r, nil
}

func health(w http.ResponseWriter, _ *http.Request) {
	if metrics == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write([]byte("OK"))
	}
}

func main() {
	var err error

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Must include target wiki API URL")
		os.Exit(1)
	}

	url := os.Args[1]

	// Initialize the data and fail fast if you can't
	if metrics, err = getSiteinfo(url); err != nil {
		log.Fatal("Failed to start:", err)
	} else {
		log.Println("Mediawiki exporter started")
	}

	go func() {
		t := time.NewTicker(time.Minute)
		for range t.C {
			if metrics, err = getSiteinfo(url); err != nil {
				log.Println("Error:", err)
			}
		}
	}()

	// Create a non-global registry.
	reg := prometheus.NewRegistry()
	counters = NewCounters(reg)

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.HandleFunc("/health", health)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
