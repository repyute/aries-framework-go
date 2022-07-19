// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/kms (interfaces: Provider,KeyManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kms "github.com/hyperledger/aries-framework-go/pkg/kms"
	secretlock "github.com/hyperledger/aries-framework-go/pkg/secretlock"
	storage "github.com/hyperledger/aries-framework-go/spi/storage"
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

// SecretLock mocks base method.
func (m *MockProvider) SecretLock() secretlock.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretLock")
	ret0, _ := ret[0].(secretlock.Service)
	return ret0
}

// SecretLock indicates an expected call of SecretLock.
func (mr *MockProviderMockRecorder) SecretLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretLock", reflect.TypeOf((*MockProvider)(nil).SecretLock))
}

// StorageProvider mocks base method.
func (m *MockProvider) StorageProvider() storage.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider")
	ret0, _ := ret[0].(storage.Provider)
	return ret0
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockProviderMockRecorder) StorageProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockProvider)(nil).StorageProvider))
}

// MockKeyManager is a mock of KeyManager interface.
type MockKeyManager struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerMockRecorder
}

// MockKeyManagerMockRecorder is the mock recorder for MockKeyManager.
type MockKeyManagerMockRecorder struct {
	mock *MockKeyManager
}

// NewMockKeyManager creates a new mock instance.
func NewMockKeyManager(ctrl *gomock.Controller) *MockKeyManager {
	mock := &MockKeyManager{ctrl: ctrl}
	mock.recorder = &MockKeyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyManager) EXPECT() *MockKeyManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockKeyManager) Create(arg0 kms.KeyType) (string, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockKeyManagerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKeyManager)(nil).Create), arg0)
}

// CreateAndExportPubKeyBytes mocks base method.
func (m *MockKeyManager) CreateAndExportPubKeyBytes(arg0 kms.KeyType) (string, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndExportPubKeyBytes", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateAndExportPubKeyBytes indicates an expected call of CreateAndExportPubKeyBytes.
func (mr *MockKeyManagerMockRecorder) CreateAndExportPubKeyBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndExportPubKeyBytes", reflect.TypeOf((*MockKeyManager)(nil).CreateAndExportPubKeyBytes), arg0)
}

// ExportPubKeyBytes mocks base method.
func (m *MockKeyManager) ExportPubKeyBytes(arg0 string) ([]byte, kms.KeyType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportPubKeyBytes", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(kms.KeyType)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportPubKeyBytes indicates an expected call of ExportPubKeyBytes.
func (mr *MockKeyManagerMockRecorder) ExportPubKeyBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportPubKeyBytes", reflect.TypeOf((*MockKeyManager)(nil).ExportPubKeyBytes), arg0)
}

// Get mocks base method.
func (m *MockKeyManager) Get(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyManagerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyManager)(nil).Get), arg0)
}

// ImportPrivateKey mocks base method.
func (m *MockKeyManager) ImportPrivateKey(arg0 interface{}, arg1 kms.KeyType, arg2 ...kms.PrivateKeyOpts) (string, interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportPrivateKey", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportPrivateKey indicates an expected call of ImportPrivateKey.
func (mr *MockKeyManagerMockRecorder) ImportPrivateKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportPrivateKey", reflect.TypeOf((*MockKeyManager)(nil).ImportPrivateKey), varargs...)
}

// PubKeyBytesToHandle mocks base method.
func (m *MockKeyManager) PubKeyBytesToHandle(arg0 []byte, arg1 kms.KeyType) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKeyBytesToHandle", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PubKeyBytesToHandle indicates an expected call of PubKeyBytesToHandle.
func (mr *MockKeyManagerMockRecorder) PubKeyBytesToHandle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKeyBytesToHandle", reflect.TypeOf((*MockKeyManager)(nil).PubKeyBytesToHandle), arg0, arg1)
}

// Rotate mocks base method.
func (m *MockKeyManager) Rotate(arg0 kms.KeyType, arg1 string) (string, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rotate", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Rotate indicates an expected call of Rotate.
func (mr *MockKeyManagerMockRecorder) Rotate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rotate", reflect.TypeOf((*MockKeyManager)(nil).Rotate), arg0, arg1)
}
