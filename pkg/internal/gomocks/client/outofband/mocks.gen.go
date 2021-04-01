// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/client/outofband (interfaces: Provider,OobService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	outofband "github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/outofband"
	kms "github.com/hyperledger/aries-framework-go/pkg/kms"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// KMS mocks base method.
func (m *MockProvider) KMS() kms.KeyManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KMS")
	ret0, _ := ret[0].(kms.KeyManager)
	return ret0
}

// KMS indicates an expected call of KMS.
func (mr *MockProviderMockRecorder) KMS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KMS", reflect.TypeOf((*MockProvider)(nil).KMS))
}

// Service mocks base method.
func (m *MockProvider) Service(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Service indicates an expected call of Service.
func (mr *MockProviderMockRecorder) Service(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockProvider)(nil).Service), arg0)
}

// ServiceEndpoint mocks base method.
func (m *MockProvider) ServiceEndpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceEndpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceEndpoint indicates an expected call of ServiceEndpoint.
func (mr *MockProviderMockRecorder) ServiceEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceEndpoint", reflect.TypeOf((*MockProvider)(nil).ServiceEndpoint))
}

// MockOobService is a mock of OobService interface.
type MockOobService struct {
	ctrl     *gomock.Controller
	recorder *MockOobServiceMockRecorder
}

// MockOobServiceMockRecorder is the mock recorder for MockOobService.
type MockOobServiceMockRecorder struct {
	mock *MockOobService
}

// NewMockOobService creates a new mock instance.
func NewMockOobService(ctrl *gomock.Controller) *MockOobService {
	mock := &MockOobService{ctrl: ctrl}
	mock.recorder = &MockOobServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOobService) EXPECT() *MockOobServiceMockRecorder {
	return m.recorder
}

// AcceptInvitation mocks base method.
func (m *MockOobService) AcceptInvitation(arg0 *outofband.Invitation, arg1 string, arg2 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptInvitation", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptInvitation indicates an expected call of AcceptInvitation.
func (mr *MockOobServiceMockRecorder) AcceptInvitation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptInvitation", reflect.TypeOf((*MockOobService)(nil).AcceptInvitation), arg0, arg1, arg2)
}

// ActionContinue mocks base method.
func (m *MockOobService) ActionContinue(arg0 string, arg1 outofband.Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionContinue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionContinue indicates an expected call of ActionContinue.
func (mr *MockOobServiceMockRecorder) ActionContinue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionContinue", reflect.TypeOf((*MockOobService)(nil).ActionContinue), arg0, arg1)
}

// ActionStop mocks base method.
func (m *MockOobService) ActionStop(arg0 string, arg1 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionStop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionStop indicates an expected call of ActionStop.
func (mr *MockOobServiceMockRecorder) ActionStop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionStop", reflect.TypeOf((*MockOobService)(nil).ActionStop), arg0, arg1)
}

// Actions mocks base method.
func (m *MockOobService) Actions() ([]outofband.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].([]outofband.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Actions indicates an expected call of Actions.
func (mr *MockOobServiceMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockOobService)(nil).Actions))
}

// RegisterActionEvent mocks base method.
func (m *MockOobService) RegisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterActionEvent indicates an expected call of RegisterActionEvent.
func (mr *MockOobServiceMockRecorder) RegisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterActionEvent", reflect.TypeOf((*MockOobService)(nil).RegisterActionEvent), arg0)
}

// RegisterMsgEvent mocks base method.
func (m *MockOobService) RegisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMsgEvent indicates an expected call of RegisterMsgEvent.
func (mr *MockOobServiceMockRecorder) RegisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMsgEvent", reflect.TypeOf((*MockOobService)(nil).RegisterMsgEvent), arg0)
}

// SaveInvitation mocks base method.
func (m *MockOobService) SaveInvitation(arg0 *outofband.Invitation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInvitation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveInvitation indicates an expected call of SaveInvitation.
func (mr *MockOobServiceMockRecorder) SaveInvitation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInvitation", reflect.TypeOf((*MockOobService)(nil).SaveInvitation), arg0)
}

// UnregisterActionEvent mocks base method.
func (m *MockOobService) UnregisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterActionEvent indicates an expected call of UnregisterActionEvent.
func (mr *MockOobServiceMockRecorder) UnregisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterActionEvent", reflect.TypeOf((*MockOobService)(nil).UnregisterActionEvent), arg0)
}

// UnregisterMsgEvent mocks base method.
func (m *MockOobService) UnregisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterMsgEvent indicates an expected call of UnregisterMsgEvent.
func (mr *MockOobServiceMockRecorder) UnregisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterMsgEvent", reflect.TypeOf((*MockOobService)(nil).UnregisterMsgEvent), arg0)
}
