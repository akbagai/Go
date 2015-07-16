package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"MSsql",
		"GET",
		"/sql",
		MSsql,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"FilterQuery",
		"GET",
		"/{key}/{value}",
		FilterQuery,
	},
	Route{
		"Subset",
		"GET",
		"/{howMany}",
		Subset,
	},
}
