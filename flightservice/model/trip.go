package model

import (
	"errors"

	"github.com/google/uuid"
)

type Trip struct {
	ID          string `json:"id"`
	User        string `json:"user"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// Repository provides access to a store of trips
type Repository interface {
	Store(trip *Trip) error
	Find(id string) (*Trip, error)
	FindAll() []*Trip
}

var ErrUnknown = errors.New("unknown trip")

func NewTripID() string {
	newUUID, _ := uuid.NewRandom()
	newID := newUUID.String()
	return newID
}
