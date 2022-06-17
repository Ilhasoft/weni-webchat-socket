// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/history/service.go

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockService) Get(contactURN, channelUUID string, limit, page int) ([]MessagePayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", contactURN, channelUUID, limit, page)
	ret0, _ := ret[0].([]MessagePayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceMockRecorder) Get(contactURN, channelUUID, limit, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockService)(nil).Get), contactURN, channelUUID, limit, page)
}

// Save mocks base method.
func (m *MockService) Save(msg MessagePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockServiceMockRecorder) Save(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockService)(nil).Save), msg)
}
