// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/luxd/snow/engine/snowman/block (interfaces: StateSyncableVM)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	block "github.com/luxdefi/luxd/snow/engine/snowman/block"
	gomock "github.com/golang/mock/gomock"
)

// MockStateSyncableVM is a mock of StateSyncableVM interface.
type MockStateSyncableVM struct {
	ctrl     *gomock.Controller
	recorder *MockStateSyncableVMMockRecorder
}

// MockStateSyncableVMMockRecorder is the mock recorder for MockStateSyncableVM.
type MockStateSyncableVMMockRecorder struct {
	mock *MockStateSyncableVM
}

// NewMockStateSyncableVM creates a new mock instance.
func NewMockStateSyncableVM(ctrl *gomock.Controller) *MockStateSyncableVM {
	mock := &MockStateSyncableVM{ctrl: ctrl}
	mock.recorder = &MockStateSyncableVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateSyncableVM) EXPECT() *MockStateSyncableVMMockRecorder {
	return m.recorder
}

// GetLastStateSummary mocks base method.
func (m *MockStateSyncableVM) GetLastStateSummary() (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastStateSummary")
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastStateSummary indicates an expected call of GetLastStateSummary.
func (mr *MockStateSyncableVMMockRecorder) GetLastStateSummary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastStateSummary", reflect.TypeOf((*MockStateSyncableVM)(nil).GetLastStateSummary))
}

// GetOngoingSyncStateSummary mocks base method.
func (m *MockStateSyncableVM) GetOngoingSyncStateSummary() (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOngoingSyncStateSummary")
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOngoingSyncStateSummary indicates an expected call of GetOngoingSyncStateSummary.
func (mr *MockStateSyncableVMMockRecorder) GetOngoingSyncStateSummary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOngoingSyncStateSummary", reflect.TypeOf((*MockStateSyncableVM)(nil).GetOngoingSyncStateSummary))
}

// GetStateSummary mocks base method.
func (m *MockStateSyncableVM) GetStateSummary(arg0 uint64) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummary", arg0)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateSummary indicates an expected call of GetStateSummary.
func (mr *MockStateSyncableVMMockRecorder) GetStateSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummary", reflect.TypeOf((*MockStateSyncableVM)(nil).GetStateSummary), arg0)
}

// ParseStateSummary mocks base method.
func (m *MockStateSyncableVM) ParseStateSummary(arg0 []byte) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseStateSummary", arg0)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseStateSummary indicates an expected call of ParseStateSummary.
func (mr *MockStateSyncableVMMockRecorder) ParseStateSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseStateSummary", reflect.TypeOf((*MockStateSyncableVM)(nil).ParseStateSummary), arg0)
}

// StateSyncEnabled mocks base method.
func (m *MockStateSyncableVM) StateSyncEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSyncEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSyncEnabled indicates an expected call of StateSyncEnabled.
func (mr *MockStateSyncableVMMockRecorder) StateSyncEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSyncEnabled", reflect.TypeOf((*MockStateSyncableVM)(nil).StateSyncEnabled))
}
