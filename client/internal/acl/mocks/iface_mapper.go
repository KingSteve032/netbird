// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netbirdio/netbird/client/internal/acl (interfaces: IFaceMapper)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	iface "github.com/netbirdio/netbird/iface"
)

// MockIFaceMapper is a mock of IFaceMapper interface.
type MockIFaceMapper struct {
	ctrl     *gomock.Controller
	recorder *MockIFaceMapperMockRecorder
}

// MockIFaceMapperMockRecorder is the mock recorder for MockIFaceMapper.
type MockIFaceMapperMockRecorder struct {
	mock *MockIFaceMapper
}

// NewMockIFaceMapper creates a new mock instance.
func NewMockIFaceMapper(ctrl *gomock.Controller) *MockIFaceMapper {
	mock := &MockIFaceMapper{ctrl: ctrl}
	mock.recorder = &MockIFaceMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFaceMapper) EXPECT() *MockIFaceMapperMockRecorder {
	return m.recorder
}

// IsUserspaceBind mocks base method.
func (m *MockIFaceMapper) IsUserspaceBind() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserspaceBind")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserspaceBind indicates an expected call of IsUserspaceBind.
func (mr *MockIFaceMapperMockRecorder) IsUserspaceBind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserspaceBind", reflect.TypeOf((*MockIFaceMapper)(nil).IsUserspaceBind))
}

// Name mocks base method.
func (m *MockIFaceMapper) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockIFaceMapperMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockIFaceMapper)(nil).Name))
}

// SetFiltering mocks base method.
func (m *MockIFaceMapper) SetFiltering(arg0 iface.PacketFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFiltering", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFiltering indicates an expected call of SetFiltering.
func (mr *MockIFaceMapperMockRecorder) SetFiltering(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFiltering", reflect.TypeOf((*MockIFaceMapper)(nil).SetFiltering), arg0)
}