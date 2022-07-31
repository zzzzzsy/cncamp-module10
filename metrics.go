package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// totalRequests = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "httpclient_requests_total",
	// 		Help: "How many HTTP requests processed, partitioned by path.",
	// 	},
	// 	[]string{"path"},
	// )

	// responseStatus = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "httpclient_response_status",
	// 		Help: "Status of HTTP response.",
	// 	},
	// 	[]string{"status"},
	// )

	httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "httpclient_response_time_seconds",
		Buckets: prometheus.ExponentialBuckets(.1, 2, 6),
		Help:    "Duration of HTTP requests, partitioned by path.",
	}, []string{"path"})
)

// middleware is used to intercept incoming HTTP calls and apply general functions upon them.
// func prometheusMiddleware(n httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 		path := r.URL.Path
// 		log.Printf("HTTP request sent to %s from %s", path, r.RemoteAddr)

// 		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
// 		defer timer.ObserveDuration()

// 		rw := negroni.NewResponseWriter(w)

// 		// call registered handler
// 		n(rw, r, ps)
// 		statusCode := rw.Status()

// 		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
// 		totalRequests.WithLabelValues(path).Inc()
// 	}
// }

// func wrapper(h http.Handler) httprouter.Handle {
// 	return prometheusMiddleware(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 		h.ServeHTTP(w, r)
// 	})
// }
