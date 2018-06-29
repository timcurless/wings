package flightservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"github.com/timcurless/wings/flightservice/inmem"
	"github.com/timcurless/wings/flightservice/model"
)

func TestGetTrip(t *testing.T) {
	mockRepo := &inmem.MockInmemClient{}

	mockRepo.On("Find", "123").Return(&model.Trip{ID: "123", User: "tess", Origin: "ORD", Destination: "SFO"}, nil)
	mockRepo.On("Find", "456").Return(&model.Trip{}, fmt.Errorf("Mock error"))

	trips = mockRepo

	Convey("Given a HTTP GET request for /trips/123", t, func() {
		req := httptest.NewRequest("GET", "/trips/123", nil)
		res := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			NewRouter().ServeHTTP(res, req)

			Convey("Then the response should be a 200", func() {
				So(res.Code, ShouldEqual, 200)

				trip := model.Trip{}
				json.Unmarshal(res.Body.Bytes(), &trip)

				So(trip.ID, ShouldEqual, "123")
				So(trip.User, ShouldEqual, "tess")
				So(trip.Origin, ShouldEqual, "ORD")
				So(trip.Destination, ShouldEqual, "SFO")
			})
		})
	})

	Convey("Given a HTTP GET request for /trips/456", t, func() {
		req := httptest.NewRequest("GET", "/trips/456", nil)
		res := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			NewRouter().ServeHTTP(res, req)

			Convey("Then the response should be a 404", func() {
				So(res.Code, ShouldEqual, 404)
			})
		})
	})

}

func TestNewTrip(t *testing.T) {

	mockRepo := &inmem.MockInmemClient{}

	body := []byte(`{"user":"tess","origin":"ORD","destination":"SFO"}`)

	trip := model.Trip{}
	json.Unmarshal(body, &trip)

	mockRepo.On("Store", mock.MatchedBy(
		func(stored *model.Trip) bool {
			return (stored.User == trip.User) &&
				(stored.Origin == trip.Origin) &&
				(stored.Destination == trip.Destination)
		},
	)).Return(nil)

	trips = mockRepo

	Convey("Given a HTTP POST request for /trips", t, func() {
		req := httptest.NewRequest("POST", "/trips", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		Convey("When the POST request is handled by the router", func() {
			NewRouter().ServeHTTP(res, req)

			Convey("Then the response code should be 201", func() {
				So(res.Code, ShouldEqual, 201)

				trip := model.Trip{}
				json.Unmarshal(res.Body.Bytes(), &trip)

				So(trip.ID, ShouldNotBeNil)
				So(trip.User, ShouldEqual, "tess")
				So(trip.Origin, ShouldEqual, "ORD")
				So(trip.Destination, ShouldEqual, "SFO")
			})
		})
	})

}

func TestGetTrips(t *testing.T) {
	mockRepo := &inmem.MockInmemClient{}

	json1 := []byte(`{"id":"123","user":"tess","origin":"ORD","destination":"SFO"}`)
	json2 := []byte(`{"id":"456","user":"mike","origin":"DFW","destination":"MSN"}`)

	myTrips := make([]*model.Trip, 2)
	json.Unmarshal(json1, &myTrips[0])
	json.Unmarshal(json2, &myTrips[1])

	mockRepo.On("FindAll").Return(myTrips)

	trips = mockRepo

	Convey("Given a HTTP GET request for /trips", t, func() {
		req := httptest.NewRequest("GET", "/trips", nil)
		res := httptest.NewRecorder()

		Convey("When the request is handled by the router", func() {
			NewRouter().ServeHTTP(res, req)

			Convey("Then the response should be a 200", func() {
				So(res.Code, ShouldEqual, 200)

				resTrips := make([]*model.Trip, 2)

				json.Unmarshal(res.Body.Bytes(), &resTrips)
				So(resTrips[0].ID, ShouldEqual, "123")
				So(resTrips[1].ID, ShouldEqual, "456")
				So(resTrips[0].User, ShouldEqual, "tess")
				So(resTrips[1].User, ShouldEqual, "mike")
				So(resTrips[0].Origin, ShouldEqual, "ORD")
				So(resTrips[1].Origin, ShouldEqual, "DFW")
				So(resTrips[0].Destination, ShouldEqual, "SFO")
				So(resTrips[1].Destination, ShouldEqual, "MSN")
			})
		})
	})

}
