package flightservice

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/timcurless/wings/flightservice/inmem"
	"github.com/timcurless/wings/flightservice/model"
)

var trips = inmem.NewTripRepository()

func NewTrip(w http.ResponseWriter, r *http.Request) {
	var body struct {
		User        string `json:"user"`
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("There was a problem with the new trip request")
	}

	tripId := model.NewTripID()

	trip := model.Trip{
		ID:          tripId,
		User:        body.User,
		Origin:      body.Origin,
		Destination: body.Destination,
	}

	if err := trips.Store(&trip); err != nil {
		log.Println("There was a problem creating the new trip")
	}

	data, _ := json.Marshal(trip)

	sendJSONResponse(w, 201, data)
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
	var tripId = mux.Vars(r)["tripId"]
	trip, err := trips.Find(tripId)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	}
	data, _ := json.Marshal(trip)

	sendJSONResponse(w, 200, data)
}

func GetTrips(w http.ResponseWriter, r *http.Request) {
	allTrips := trips.FindAll()

	data, _ := json.Marshal(allTrips)

	sendJSONResponse(w, 200, data)
}

func sendJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
