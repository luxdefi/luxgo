// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/proposervm/proposer (interfaces: Windower)

// Package proposer is a generated GoMock package.
package proposer

import (
	context "context"
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	gomock "github.com/golang/mock/gomock"
)

// MockWindower is a mock of Windower interface.
type MockWindower struct {
	ctrl     *gomock.Controller
	recorder *MockWindowerMockRecorder
}

// MockWindowerMockRecorder is the mock recorder for MockWindower.
type MockWindowerMockRecorder struct {
	mock *MockWindower
}

// NewMockWindower creates a new mock instance.
func NewMockWindower(ctrl *gomock.Controller) *MockWindower {
	mock := &MockWindower{ctrl: ctrl}
	mock.recorder = &MockWindowerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWindower) EXPECT() *MockWindowerMockRecorder {
	return m.recorder
}

// Delay mocks base method.
func (m *MockWindower) Delay(arg0 context.Context, arg1, arg2 uint64, arg3 ids.NodeID) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delay indicates an expected call of Delay.
func (mr *MockWindowerMockRecorder) Delay(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockWindower)(nil).Delay), arg0, arg1, arg2, arg3)
}
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 6bf817bb8 (Add proposer list to proposer.Windower (#2366))

// Proposers mocks base method.
func (m *MockWindower) Proposers(arg0 context.Context, arg1, arg2 uint64) ([]ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proposers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Proposers indicates an expected call of Proposers.
func (mr *MockWindowerMockRecorder) Proposers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proposers", reflect.TypeOf((*MockWindower)(nil).Proposers), arg0, arg1, arg2)
}
<<<<<<< HEAD
=======
>>>>>>> 37ccd9a48 (Add BuildBlockWithContext as an optional VM method (#2210))
=======
>>>>>>> 6bf817bb8 (Add proposer list to proposer.Windower (#2366))
