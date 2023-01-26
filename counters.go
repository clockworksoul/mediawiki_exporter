package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type counterset struct {
	Activeusers prometheus.CounterFunc
	Admins      prometheus.CounterFunc
	Articles    prometheus.CounterFunc
	Edits       prometheus.CounterFunc
	Images      prometheus.CounterFunc
	Jobs        prometheus.CounterFunc
	Pages       prometheus.CounterFunc
	Users       prometheus.CounterFunc
}

var counters counterset

func NewCounters(reg prometheus.Registerer) counterset {
	c := counterset{
		Activeusers: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_activeusers",
				Help: "Number of active users.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Activeusers) },
		),
		Admins: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_admins",
				Help: "Number of administrators.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Admins) },
		),
		Articles: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_articles",
				Help: "Number of articles.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Articles) },
		),
		Edits: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_edits",
				Help: "Number of edits.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Edits) },
		),
		Images: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_images",
				Help: "Number of images.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Images) },
		),
		Jobs: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_jobs",
				Help: "Number of jobs.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Jobs) },
		),
		Pages: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_pages",
				Help: "Number of pages.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Pages) },
		),
		Users: prometheus.NewCounterFunc(
			prometheus.CounterOpts{
				Name: "mediawiki_statistics_users",
				Help: "Number of users.",
			},
			func() float64 { return float64(metrics.Query.Statistics.Users) },
		),
	}

	reg.MustRegister(c.Activeusers)
	reg.MustRegister(c.Admins)
	reg.MustRegister(c.Articles)
	reg.MustRegister(c.Edits)
	reg.MustRegister(c.Images)
	reg.MustRegister(c.Jobs)
	reg.MustRegister(c.Pages)
	reg.MustRegister(c.Users)

	return c
}
