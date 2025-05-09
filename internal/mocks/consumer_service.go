package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"xyz_golang/internal/domain"
)

type ConsumerService struct {
	mock.Mock
}

func (m *ConsumerService) GetByID(id uuid.UUID) (*domain.Consumer, error) {
	ret := m.Called(id)
	var r0 *domain.Consumer
	if rf, ok := ret.Get(0).(func(id uuid.UUID) *domain.Consumer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Consumer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(id uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *ConsumerService) GetLimit(id uuid.UUID) (*[]domain.Limit, error) {
	ret := m.Called(id)
	var r0 *[]domain.Limit
	if rf, ok := ret.Get(0).(func(id uuid.UUID) *[]domain.Limit); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Limit)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(id uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (m *ConsumerService) Store(consumer *domain.Consumer) error {
	ret := m.Called(consumer)
	var r0 error
	if rf, ok := ret.Get(0).(func(consumer *domain.Consumer) error); ok {
		r0 = rf(consumer)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

func (m *ConsumerService) StoreLimit(limit *domain.Limit) error {
	ret := m.Called(limit)
	var r0 error
	if rf, ok := ret.Get(0).(func(limit *domain.Limit) error); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
