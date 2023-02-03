package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/clockworksoul/mediawiki"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getSiteinfo(url string) (*mediawiki.SiteinfoResponse, error) {
	c, err := mediawiki.New(url, "mediawiki-exporter")
	if err != nil {
		return nil, fmt.Errorf("failed to build mediawiki client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	password := os.Getenv("MEDIAWIKI_PASSWORD")
	username := os.Getenv("MEDIAWIKI_USERNAME")

	if len(username) > 0 {
		if _, err := c.BotLogin(ctx, username, password); err != nil {
			return nil, err
		}
	}

	r, err := c.Siteinfo().Prop("statistics").Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve statistics: %w", err)
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

	url := os.Getenv("MEDIAWIKI_API_URL")

	if url == "" && len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Must include target wiki API URL or use the MEDIAWIKI_API_URL envvar")
		os.Exit(3)
	}

	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	// Initialize the data and fail fast if you can't
	if metrics, err = getSiteinfo(url); err != nil {
		if strings.Contains(err.Error(), "login failure: ") {
			fmt.Fprintln(os.Stderr, "Invalid credentials:", err)
			os.Exit(4)
		} else if strings.Contains(err.Error(), "readapidenied: ") {
			fmt.Fprintln(os.Stderr, "Access denied:", err)
			os.Exit(5)
		}

		fmt.Fprintln(os.Stderr, "Failed to start:", err)
		os.Exit(2)
	} else {
		log.Println("Mediawiki exporter started")
	}

	go func() {
		t := time.NewTicker(time.Minute)
		for range t.C {
			m.Lock()
			if metrics, err = getSiteinfo(url); err != nil {
				log.Println("Error:", err)
			}
			m.Unlock()
		}
	}()

	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Set up the counters
	NewCounters(reg)

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.HandleFunc("/health", health)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
