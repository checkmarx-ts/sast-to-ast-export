// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/app/interfaces (interfaces: QueryMappingRepo)
//
// Generated by this command:
//
//	mockgen -package mock_interfaces -destination test/mocks/app/ast_query_mapping/mock_provider.go github.com/checkmarxDev/ast-sast-export/internal/app/interfaces QueryMappingRepo
//
// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	querymapping "github.com/checkmarxDev/ast-sast-export/internal/app/querymapping"
	gomock "go.uber.org/mock/gomock"
)

// MockQueryMappingRepo is a mock of QueryMappingRepo interface.
type MockQueryMappingRepo struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMappingRepoMockRecorder
}

// MockQueryMappingRepoMockRecorder is the mock recorder for MockQueryMappingRepo.
type MockQueryMappingRepoMockRecorder struct {
	mock *MockQueryMappingRepo
}

// NewMockQueryMappingRepo creates a new mock instance.
func NewMockQueryMappingRepo(ctrl *gomock.Controller) *MockQueryMappingRepo {
	mock := &MockQueryMappingRepo{ctrl: ctrl}
	mock.recorder = &MockQueryMappingRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryMappingRepo) EXPECT() *MockQueryMappingRepoMockRecorder {
	return m.recorder
}

// AddQueryMapping mocks base method.
func (m *MockQueryMappingRepo) AddQueryMapping(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddQueryMapping", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddQueryMapping indicates an expected call of AddQueryMapping.
func (mr *MockQueryMappingRepoMockRecorder) AddQueryMapping(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQueryMapping", reflect.TypeOf((*MockQueryMappingRepo)(nil).AddQueryMapping), arg0, arg1, arg2, arg3)
}

// GetMapping mocks base method.
func (m *MockQueryMappingRepo) GetMapping() []querymapping.QueryMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapping")
	ret0, _ := ret[0].([]querymapping.QueryMap)
	return ret0
}

// GetMapping indicates an expected call of GetMapping.
func (mr *MockQueryMappingRepoMockRecorder) GetMapping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapping", reflect.TypeOf((*MockQueryMappingRepo)(nil).GetMapping))
}
