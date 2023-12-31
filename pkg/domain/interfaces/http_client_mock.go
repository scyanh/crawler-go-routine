// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/interfaces/http_client.go

// Package interfaces is a generated GoMock package.
package interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHTTPClient is a mock of IHTTPClient interface.
type MockIHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockIHTTPClientMockRecorder
}

// MockIHTTPClientMockRecorder is the mock recorder for MockIHTTPClient.
type MockIHTTPClientMockRecorder struct {
	mock *MockIHTTPClient
}

// NewMockIHTTPClient creates a new mock instance.
func NewMockIHTTPClient(ctrl *gomock.Controller) *MockIHTTPClient {
	mock := &MockIHTTPClient{ctrl: ctrl}
	mock.recorder = &MockIHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHTTPClient) EXPECT() *MockIHTTPClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIHTTPClient) Get(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIHTTPClientMockRecorder) Get(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIHTTPClient)(nil).Get), url)
}
