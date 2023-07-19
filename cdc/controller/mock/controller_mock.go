// Code generated by MockGen. DO NOT EDIT.
// Source: cdc/controller/controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	orchestrator "github.com/pingcap/tiflow/pkg/orchestrator"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// AsyncStop mocks base method.
func (m *MockController) AsyncStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AsyncStop")
}

// AsyncStop indicates an expected call of AsyncStop.
func (mr *MockControllerMockRecorder) AsyncStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncStop", reflect.TypeOf((*MockController)(nil).AsyncStop))
}

// Tick mocks base method.
func (m *MockController) Tick(ctx context.Context, state orchestrator.ReactorState) (orchestrator.ReactorState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick", ctx, state)
	ret0, _ := ret[0].(orchestrator.ReactorState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tick indicates an expected call of Tick.
func (mr *MockControllerMockRecorder) Tick(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockController)(nil).Tick), ctx, state)
}