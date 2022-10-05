// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rguinea/GoGRPCTutorial/internal/rocket (interfaces: Store)

// Package rocket is a generated GoMock package.
package rocket

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteRocket mocks base method.
func (m *MockStore) DeleteRocket(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRocket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRocket indicates an expected call of DeleteRocket.
func (mr *MockStoreMockRecorder) DeleteRocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRocket", reflect.TypeOf((*MockStore)(nil).DeleteRocket), arg0)
}

// GetRocketById mocks base method.
func (m *MockStore) GetRocketById(arg0 string) (Rocket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRocketById", arg0)
	ret0, _ := ret[0].(Rocket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRocketById indicates an expected call of GetRocketById.
func (mr *MockStoreMockRecorder) GetRocketById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRocketById", reflect.TypeOf((*MockStore)(nil).GetRocketById), arg0)
}

// InsertRocket mocks base method.
func (m *MockStore) InsertRocket(arg0 Rocket) (Rocket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRocket", arg0)
	ret0, _ := ret[0].(Rocket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRocket indicates an expected call of InsertRocket.
func (mr *MockStoreMockRecorder) InsertRocket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRocket", reflect.TypeOf((*MockStore)(nil).InsertRocket), arg0)
}
