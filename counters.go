package main

import (
	"sync"

	"github.com/clockworksoul/mediawiki"
	"github.com/prometheus/client_golang/prometheus"
)

var metrics *mediawiki.SiteinfoResponse
var m = &sync.RWMutex{}

func NewCounters(reg prometheus.Registerer) {
	activeusers := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_activeusers",
			Help: "Number of active users.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Activeusers)
		},
	)
	admins := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_admins",
			Help: "Number of administrators.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Admins)
		},
	)
	articles := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_articles",
			Help: "Number of articles.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Articles)
		},
	)
	edits := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_edits",
			Help: "Number of edits.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Edits)
		},
	)
	images := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_images",
			Help: "Number of images.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Images)
		},
	)
	jobs := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_jobs",
			Help: "Number of jobs.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Jobs)
		},
	)
	pages := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_pages",
			Help: "Number of pages.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Pages)
		},
	)
	users := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mediawiki_statistics_users",
			Help: "Number of users.",
		},
		func() float64 {
			m.RLock()
			defer m.RUnlock()
			return float64(metrics.Query.Statistics.Users)
		},
	)

	reg.MustRegister(activeusers)
	reg.MustRegister(admins)
	reg.MustRegister(articles)
	reg.MustRegister(edits)
	reg.MustRegister(images)
	reg.MustRegister(jobs)
	reg.MustRegister(pages)
	reg.MustRegister(users)
}
