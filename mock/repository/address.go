// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/repository/address/address.go

// Package mock_address is a generated GoMock package.
package repository

import (
	context "context"
	entity "m1-article-service/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAddress is a mock of Address interface.
type MockAddress struct {
	ctrl     *gomock.Controller
	recorder *MockAddressMockRecorder
}

// MockAddressMockRecorder is the mock recorder for MockAddress.
type MockAddressMockRecorder struct {
	mock *MockAddress
}

// NewMockAddress creates a new mock instance.
func NewMockAddress(ctrl *gomock.Controller) *MockAddress {
	mock := &MockAddress{ctrl: ctrl}
	mock.recorder = &MockAddressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddress) EXPECT() *MockAddressMockRecorder {
	return m.recorder
}

// BatchCreate mocks base method.
func (m *MockAddress) BatchCreate(arg0 context.Context, arg1 []*entity.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreate indicates an expected call of BatchCreate.
func (mr *MockAddressMockRecorder) BatchCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreate", reflect.TypeOf((*MockAddress)(nil).BatchCreate), arg0, arg1)
}

// Create mocks base method.
func (m *MockAddress) Create(arg0 context.Context, arg1 *entity.Address) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAddressMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAddress)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockAddress) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAddressMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAddress)(nil).Delete), arg0, arg1)
}

// Detail mocks base method.
func (m *MockAddress) Detail(arg0 context.Context, arg1 int64) (*entity.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detail", arg0, arg1)
	ret0, _ := ret[0].(*entity.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail.
func (mr *MockAddressMockRecorder) Detail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockAddress)(nil).Detail), arg0, arg1)
}

// List mocks base method.
func (m *MockAddress) List(arg0 context.Context, arg1 uint16) ([]*entity.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAddressMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAddress)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockAddress) Update(arg0 context.Context, arg1 *entity.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAddressMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAddress)(nil).Update), arg0, arg1)
}
