package flightservice

import (
	"net/http"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"NewTrip",
		"/trips",
		"POST",
		NewTrip,
	},
	Route{
		"GetTrips",
		"/trips",
		"GET",
		GetTrips,
	},
	Route{
		"GetTrip",
		"/trips/{tripId}",
		"GET",
		GetTrip,
	},
}
