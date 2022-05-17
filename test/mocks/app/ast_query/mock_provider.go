// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/app/interfaces (interfaces: ASTQueryProvider)

// Package mock_interfaces_query_common is a generated GoMock package.
package mock_interfaces_query_common

import (
	reflect "reflect"

	soap "github.com/checkmarxDev/ast-sast-export/internal/integration/soap"
	gomock "github.com/golang/mock/gomock"
)

// MockASTQueryProvider is a mock of ASTQueryProvider interface.
type MockASTQueryProvider struct {
	ctrl     *gomock.Controller
	recorder *MockASTQueryProviderMockRecorder
}

// MockASTQueryProviderMockRecorder is the mock recorder for MockASTQueryProvider.
type MockASTQueryProviderMockRecorder struct {
	mock *MockASTQueryProvider
}

// NewMockASTQueryProvider creates a new mock instance.
func NewMockASTQueryProvider(ctrl *gomock.Controller) *MockASTQueryProvider {
	mock := &MockASTQueryProvider{ctrl: ctrl}
	mock.recorder = &MockASTQueryProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockASTQueryProvider) EXPECT() *MockASTQueryProviderMockRecorder {
	return m.recorder
}

// GetCustomQueriesList mocks base method.
func (m *MockASTQueryProvider) GetCustomQueriesList() (*soap.GetQueryCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomQueriesList")
	ret0, _ := ret[0].(*soap.GetQueryCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomQueriesList indicates an expected call of GetCustomQueriesList.
func (mr *MockASTQueryProviderMockRecorder) GetCustomQueriesList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomQueriesList", reflect.TypeOf((*MockASTQueryProvider)(nil).GetCustomQueriesList))
}

// GetQueryID mocks base method.
func (m *MockASTQueryProvider) GetQueryID(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryID", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryID indicates an expected call of GetQueryID.
func (mr *MockASTQueryProviderMockRecorder) GetQueryID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryID", reflect.TypeOf((*MockASTQueryProvider)(nil).GetQueryID), arg0, arg1, arg2)
}
