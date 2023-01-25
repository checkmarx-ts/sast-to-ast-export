// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/app/interfaces (interfaces: InstallationProvider)

// Package mock_installation_interfaces is a generated GoMock package.
package mock_installation_interfaces

import (
	reflect "reflect"

	soap "github.com/checkmarxDev/ast-sast-export/internal/integration/soap"
	gomock "github.com/golang/mock/gomock"
)

// MockInstallationProvider is a mock of InstallationProvider interface.
type MockInstallationProvider struct {
	ctrl     *gomock.Controller
	recorder *MockInstallationProviderMockRecorder
}

// MockInstallationProviderMockRecorder is the mock recorder for MockInstallationProvider.
type MockInstallationProviderMockRecorder struct {
	mock *MockInstallationProvider
}

// NewMockInstallationProvider creates a new mock instance.
func NewMockInstallationProvider(ctrl *gomock.Controller) *MockInstallationProvider {
	mock := &MockInstallationProvider{ctrl: ctrl}
	mock.recorder = &MockInstallationProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallationProvider) EXPECT() *MockInstallationProviderMockRecorder {
	return m.recorder
}

// GetInstallationSettings mocks base method.
func (m *MockInstallationProvider) GetInstallationSettings() (*soap.GetInstallationSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallationSettings")
	ret0, _ := ret[0].(*soap.GetInstallationSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallationSettings indicates an expected call of GetInstallationSettings.
func (mr *MockInstallationProviderMockRecorder) GetInstallationSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallationSettings", reflect.TypeOf((*MockInstallationProvider)(nil).GetInstallationSettings))
}
