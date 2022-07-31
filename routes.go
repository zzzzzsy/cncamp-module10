package main

import (
	"net/http"

	"github.com/gorilla/context"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"ping", "GET", "/ping", ping},
		Route{"hello", "GET", "/hello", hello},
		Route{"ip", "GET", "/ip", ip},
		Route{"metrics", "GET", "/metrics", wrapHandler(promhttp.Handler())},
	}
	return routes
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		context.Set(req, "params", ps)
		h.ServeHTTP(rw, req)
	}
}
