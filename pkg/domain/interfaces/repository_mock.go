// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/interfaces/repository.go

// Package interfaces is a generated GoMock package.
package interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/scyanh/crawler/pkg/domain/entities"
)

// MockIMemoryLinkRepository is a mock of IMemoryLinkRepository interface.
type MockIMemoryLinkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMemoryLinkRepositoryMockRecorder
}

// MockIMemoryLinkRepositoryMockRecorder is the mock recorder for MockIMemoryLinkRepository.
type MockIMemoryLinkRepositoryMockRecorder struct {
	mock *MockIMemoryLinkRepository
}

// NewMockIMemoryLinkRepository creates a new mock instance.
func NewMockIMemoryLinkRepository(ctrl *gomock.Controller) *MockIMemoryLinkRepository {
	mock := &MockIMemoryLinkRepository{ctrl: ctrl}
	mock.recorder = &MockIMemoryLinkRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMemoryLinkRepository) EXPECT() *MockIMemoryLinkRepositoryMockRecorder {
	return m.recorder
}

// HasBeenVisited mocks base method.
func (m *MockIMemoryLinkRepository) HasBeenVisited(link entities.Link) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBeenVisited", link)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBeenVisited indicates an expected call of HasBeenVisited.
func (mr *MockIMemoryLinkRepositoryMockRecorder) HasBeenVisited(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBeenVisited", reflect.TypeOf((*MockIMemoryLinkRepository)(nil).HasBeenVisited), link)
}

// MarkAsVisited mocks base method.
func (m *MockIMemoryLinkRepository) MarkAsVisited(link entities.Link) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkAsVisited", link)
}

// MarkAsVisited indicates an expected call of MarkAsVisited.
func (mr *MockIMemoryLinkRepositoryMockRecorder) MarkAsVisited(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsVisited", reflect.TypeOf((*MockIMemoryLinkRepository)(nil).MarkAsVisited), link)
}
