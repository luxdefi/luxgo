// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxfi/node/snow/networking/router (interfaces: Router)

// Package router is a generated GoMock package.
package router

import (
	context "context"
	reflect "reflect"
	time "time"

	ids "github.com/luxfi/node/ids"
	message "github.com/luxfi/node/message"
	p2p "github.com/luxfi/node/proto/pb/p2p"
	handler "github.com/luxfi/node/snow/networking/handler"
	timeout "github.com/luxfi/node/snow/networking/timeout"
	logging "github.com/luxfi/node/utils/logging"
	set "github.com/luxfi/node/utils/set"
	version "github.com/luxfi/node/version"
	gomock "go.uber.org/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// AddChain mocks base method.
func (m *MockRouter) AddChain(arg0 context.Context, arg1 handler.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChain", arg0, arg1)
}

// AddChain indicates an expected call of AddChain.
func (mr *MockRouterMockRecorder) AddChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChain", reflect.TypeOf((*MockRouter)(nil).AddChain), arg0, arg1)
}

// Benched mocks base method.
func (m *MockRouter) Benched(arg0 ids.ID, arg1 ids.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Benched", arg0, arg1)
}

// Benched indicates an expected call of Benched.
func (mr *MockRouterMockRecorder) Benched(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Benched", reflect.TypeOf((*MockRouter)(nil).Benched), arg0, arg1)
}

// Connected mocks base method.
func (m *MockRouter) Connected(arg0 ids.NodeID, arg1 *version.Application, arg2 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connected", arg0, arg1, arg2)
}

// Connected indicates an expected call of Connected.
func (mr *MockRouterMockRecorder) Connected(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockRouter)(nil).Connected), arg0, arg1, arg2)
}

// Disconnected mocks base method.
func (m *MockRouter) Disconnected(arg0 ids.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnected", arg0)
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockRouterMockRecorder) Disconnected(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockRouter)(nil).Disconnected), arg0)
}

// HandleInbound mocks base method.
func (m *MockRouter) HandleInbound(arg0 context.Context, arg1 message.InboundMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleInbound", arg0, arg1)
}

// HandleInbound indicates an expected call of HandleInbound.
func (mr *MockRouterMockRecorder) HandleInbound(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInbound", reflect.TypeOf((*MockRouter)(nil).HandleInbound), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockRouter) HealthCheck(arg0 context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockRouterMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockRouter)(nil).HealthCheck), arg0)
}

// Initialize mocks base method.
func (m *MockRouter) Initialize(arg0 ids.NodeID, arg1 logging.Logger, arg2 timeout.Manager, arg3 time.Duration, arg4 set.Set[ids.ID], arg5 bool, arg6 set.Set[ids.ID], arg7 func(int), arg8 HealthConfig, arg9 string, arg10 prometheus.Registerer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockRouterMockRecorder) Initialize(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockRouter)(nil).Initialize), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RegisterRequest mocks base method.
func (m *MockRouter) RegisterRequest(arg0 context.Context, arg1 ids.NodeID, arg2, arg3 ids.ID, arg4 uint32, arg5 message.Op, arg6 message.InboundMessage, arg7 p2p.EngineType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRequest", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// RegisterRequest indicates an expected call of RegisterRequest.
func (mr *MockRouterMockRecorder) RegisterRequest(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRequest", reflect.TypeOf((*MockRouter)(nil).RegisterRequest), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// Shutdown mocks base method.
func (m *MockRouter) Shutdown(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown", arg0)
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockRouterMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRouter)(nil).Shutdown), arg0)
}

// Unbenched mocks base method.
func (m *MockRouter) Unbenched(arg0 ids.ID, arg1 ids.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unbenched", arg0, arg1)
}

// Unbenched indicates an expected call of Unbenched.
func (mr *MockRouterMockRecorder) Unbenched(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unbenched", reflect.TypeOf((*MockRouter)(nil).Unbenched), arg0, arg1)
}
