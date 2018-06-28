package inmem

import (
	"sync"

	"github.com/timcurless/wings/flightservice/model"
)

type tripRepository struct {
	mtx   sync.RWMutex
	trips map[string]*model.Trip
}

func NewTripRepository() model.Repository {
	return &tripRepository{
		trips: make(map[string]*model.Trip),
	}
}

func (r *tripRepository) Store(t *model.Trip) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.trips[t.ID] = t
	return nil
}

func (r *tripRepository) Find(id string) (*model.Trip, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.trips[id]; ok {
		return val, nil
	}
	return nil, model.ErrUnknown
}

func (r *tripRepository) FindAll() []*model.Trip {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	t := make([]*model.Trip, 0, len(r.trips))
	for _, val := range r.trips {
		t = append(t, val)
	}
	return t
}
