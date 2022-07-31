package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func ping(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defaultHandler(rw, r, "Pong")
}

func hello(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))
	defer timer.ObserveDuration()
	randomSleep()
	defaultHandler(rw, r, "Hello LiveRamp SRE")
}

func ip(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	randomSleep()
	defaultHandler(rw, r, getIP(r))
}

func defaultHandler(rw http.ResponseWriter, r *http.Request, msg string) {
	setResponseHeader(rw, r)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(msg))
	logs := HttpServerLog{
		Request:    *r,
		StatusCode: http.StatusOK,
	}
	httpServerLog(logs)
}

func setResponseHeader(rw http.ResponseWriter, r *http.Request) {
	if v, exist := os.LookupEnv("VERSION"); exist {
		rw.Header().Add("User-Version", v)
	} else {
		log.Infof("The env variable %s does not exist\n", v)
	}

	for name, values := range r.Header {
		for _, value := range values {
			rw.Header().Set(name, value)
		}
	}
}
