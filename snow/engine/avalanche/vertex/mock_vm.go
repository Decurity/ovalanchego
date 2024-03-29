// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex (interfaces: DAGVM)

// Package vertex is a generated GoMock package.
package vertex

import (
	context "context"
	reflect "reflect"
	time "time"

	manager "github.com/ava-labs/avalanchego/database/manager"
	ids "github.com/ava-labs/avalanchego/ids"
	snow "github.com/ava-labs/avalanchego/snow"
	snowman "github.com/ava-labs/avalanchego/snow/consensus/snowman"
	snowstorm "github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	common "github.com/ava-labs/avalanchego/snow/engine/common"
	version "github.com/ava-labs/avalanchego/version"
	gomock "github.com/golang/mock/gomock"
)

// MockDAGVM is a mock of DAGVM interface.
type MockDAGVM struct {
	ctrl     *gomock.Controller
	recorder *MockDAGVMMockRecorder
}

// MockDAGVMMockRecorder is the mock recorder for MockDAGVM.
type MockDAGVMMockRecorder struct {
	mock *MockDAGVM
}

// NewMockDAGVM creates a new mock instance.
func NewMockDAGVM(ctrl *gomock.Controller) *MockDAGVM {
	mock := &MockDAGVM{ctrl: ctrl}
	mock.recorder = &MockDAGVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAGVM) EXPECT() *MockDAGVMMockRecorder {
	return m.recorder
}

// AppGossip mocks base method.
func (m *MockDAGVM) AppGossip(arg0 context.Context, arg1 ids.NodeID, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppGossip indicates an expected call of AppGossip.
func (mr *MockDAGVMMockRecorder) AppGossip(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*MockDAGVM)(nil).AppGossip), arg0, arg1, arg2)
}

// AppRequest mocks base method.
func (m *MockDAGVM) AppRequest(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequest indicates an expected call of AppRequest.
func (mr *MockDAGVMMockRecorder) AppRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*MockDAGVM)(nil).AppRequest), arg0, arg1, arg2, arg3, arg4)
}

// AppRequestFailed mocks base method.
func (m *MockDAGVM) AppRequestFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequestFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequestFailed indicates an expected call of AppRequestFailed.
func (mr *MockDAGVMMockRecorder) AppRequestFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequestFailed", reflect.TypeOf((*MockDAGVM)(nil).AppRequestFailed), arg0, arg1, arg2)
}

// AppResponse mocks base method.
func (m *MockDAGVM) AppResponse(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockDAGVMMockRecorder) AppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockDAGVM)(nil).AppResponse), arg0, arg1, arg2, arg3)
}

// BuildBlock mocks base method.
func (m *MockDAGVM) BuildBlock(arg0 context.Context) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildBlock indicates an expected call of BuildBlock.
func (mr *MockDAGVMMockRecorder) BuildBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildBlock", reflect.TypeOf((*MockDAGVM)(nil).BuildBlock), arg0)
}

// Connected mocks base method.
func (m *MockDAGVM) Connected(arg0 context.Context, arg1 ids.NodeID, arg2 *version.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connected indicates an expected call of Connected.
func (mr *MockDAGVMMockRecorder) Connected(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockDAGVM)(nil).Connected), arg0, arg1, arg2)
}

// CreateHandlers mocks base method.
func (m *MockDAGVM) CreateHandlers(arg0 context.Context) (map[string]*common.HTTPHandler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHandlers", arg0)
	ret0, _ := ret[0].(map[string]*common.HTTPHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHandlers indicates an expected call of CreateHandlers.
func (mr *MockDAGVMMockRecorder) CreateHandlers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHandlers", reflect.TypeOf((*MockDAGVM)(nil).CreateHandlers), arg0)
}

// CreateStaticHandlers mocks base method.
func (m *MockDAGVM) CreateStaticHandlers(arg0 context.Context) (map[string]*common.HTTPHandler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStaticHandlers", arg0)
	ret0, _ := ret[0].(map[string]*common.HTTPHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStaticHandlers indicates an expected call of CreateStaticHandlers.
func (mr *MockDAGVMMockRecorder) CreateStaticHandlers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStaticHandlers", reflect.TypeOf((*MockDAGVM)(nil).CreateStaticHandlers), arg0)
}

// CrossChainAppRequest mocks base method.
func (m *MockDAGVM) CrossChainAppRequest(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequest indicates an expected call of CrossChainAppRequest.
func (mr *MockDAGVMMockRecorder) CrossChainAppRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequest", reflect.TypeOf((*MockDAGVM)(nil).CrossChainAppRequest), arg0, arg1, arg2, arg3, arg4)
}

// CrossChainAppRequestFailed mocks base method.
func (m *MockDAGVM) CrossChainAppRequestFailed(arg0 context.Context, arg1 ids.ID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequestFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequestFailed indicates an expected call of CrossChainAppRequestFailed.
func (mr *MockDAGVMMockRecorder) CrossChainAppRequestFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequestFailed", reflect.TypeOf((*MockDAGVM)(nil).CrossChainAppRequestFailed), arg0, arg1, arg2)
}

// CrossChainAppResponse mocks base method.
func (m *MockDAGVM) CrossChainAppResponse(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppResponse indicates an expected call of CrossChainAppResponse.
func (mr *MockDAGVMMockRecorder) CrossChainAppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppResponse", reflect.TypeOf((*MockDAGVM)(nil).CrossChainAppResponse), arg0, arg1, arg2, arg3)
}

// Disconnected mocks base method.
func (m *MockDAGVM) Disconnected(arg0 context.Context, arg1 ids.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnected", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockDAGVMMockRecorder) Disconnected(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockDAGVM)(nil).Disconnected), arg0, arg1)
}

// GetBlock mocks base method.
func (m *MockDAGVM) GetBlock(arg0 context.Context, arg1 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockDAGVMMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockDAGVM)(nil).GetBlock), arg0, arg1)
}

// GetTx mocks base method.
func (m *MockDAGVM) GetTx(arg0 context.Context, arg1 ids.ID) (snowstorm.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", arg0, arg1)
	ret0, _ := ret[0].(snowstorm.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx.
func (mr *MockDAGVMMockRecorder) GetTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockDAGVM)(nil).GetTx), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockDAGVM) HealthCheck(arg0 context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockDAGVMMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDAGVM)(nil).HealthCheck), arg0)
}

// Initialize mocks base method.
func (m *MockDAGVM) Initialize(arg0 context.Context, arg1 *snow.Context, arg2 manager.Manager, arg3, arg4, arg5 []byte, arg6 chan<- common.Message, arg7 []*common.Fx, arg8 common.AppSender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockDAGVMMockRecorder) Initialize(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockDAGVM)(nil).Initialize), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// LastAccepted mocks base method.
func (m *MockDAGVM) LastAccepted(arg0 context.Context) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastAccepted", arg0)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastAccepted indicates an expected call of LastAccepted.
func (mr *MockDAGVMMockRecorder) LastAccepted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastAccepted", reflect.TypeOf((*MockDAGVM)(nil).LastAccepted), arg0)
}

// Linearize mocks base method.
func (m *MockDAGVM) Linearize(arg0 context.Context, arg1 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Linearize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Linearize indicates an expected call of Linearize.
func (mr *MockDAGVMMockRecorder) Linearize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Linearize", reflect.TypeOf((*MockDAGVM)(nil).Linearize), arg0, arg1)
}

// ParseBlock mocks base method.
func (m *MockDAGVM) ParseBlock(arg0 context.Context, arg1 []byte) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBlock", arg0, arg1)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseBlock indicates an expected call of ParseBlock.
func (mr *MockDAGVMMockRecorder) ParseBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBlock", reflect.TypeOf((*MockDAGVM)(nil).ParseBlock), arg0, arg1)
}

// ParseTx mocks base method.
func (m *MockDAGVM) ParseTx(arg0 context.Context, arg1 []byte) (snowstorm.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTx", arg0, arg1)
	ret0, _ := ret[0].(snowstorm.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTx indicates an expected call of ParseTx.
func (mr *MockDAGVMMockRecorder) ParseTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTx", reflect.TypeOf((*MockDAGVM)(nil).ParseTx), arg0, arg1)
}

// PendingTxs mocks base method.
func (m *MockDAGVM) PendingTxs(arg0 context.Context) []snowstorm.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingTxs", arg0)
	ret0, _ := ret[0].([]snowstorm.Tx)
	return ret0
}

// PendingTxs indicates an expected call of PendingTxs.
func (mr *MockDAGVMMockRecorder) PendingTxs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingTxs", reflect.TypeOf((*MockDAGVM)(nil).PendingTxs), arg0)
}

// SetPreference mocks base method.
func (m *MockDAGVM) SetPreference(arg0 context.Context, arg1 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPreference", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPreference indicates an expected call of SetPreference.
func (mr *MockDAGVMMockRecorder) SetPreference(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPreference", reflect.TypeOf((*MockDAGVM)(nil).SetPreference), arg0, arg1)
}

// SetState mocks base method.
func (m *MockDAGVM) SetState(arg0 context.Context, arg1 snow.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState.
func (mr *MockDAGVMMockRecorder) SetState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockDAGVM)(nil).SetState), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockDAGVM) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockDAGVMMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockDAGVM)(nil).Shutdown), arg0)
}

// Version mocks base method.
func (m *MockDAGVM) Version(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockDAGVMMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockDAGVM)(nil).Version), arg0)
}
