package inmem

import (
	"github.com/stretchr/testify/mock"
	"github.com/timcurless/wings/flightservice/model"
)

type MockInmemClient struct {
	mock.Mock
}

func (m *MockInmemClient) Store(t *model.Trip) error {
	args := m.Mock.Called(t)
	return args.Error(0)
}

func (m *MockInmemClient) Find(id string) (*model.Trip, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(*model.Trip), args.Error(1)
}

func (m *MockInmemClient) FindAll() []*model.Trip {
	args := m.Mock.Called()
	return args.Get(0).([]*model.Trip)
}
