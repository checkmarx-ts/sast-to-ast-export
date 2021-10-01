// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/export (interfaces: Exporter)

// Package export is a generated GoMock package.
package export

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExporter is a mock of Exporter interface.
type MockExporter struct {
	ctrl     *gomock.Controller
	recorder *MockExporterMockRecorder
}

// MockExporterMockRecorder is the mock recorder for MockExporter.
type MockExporterMockRecorder struct {
	mock *MockExporter
}

// NewMockExporter creates a new mock instance.
func NewMockExporter(ctrl *gomock.Controller) *MockExporter {
	mock := &MockExporter{ctrl: ctrl}
	mock.recorder = &MockExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExporter) EXPECT() *MockExporterMockRecorder {
	return m.recorder
}

// AddFile mocks base method.
func (m *MockExporter) AddFile(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFile indicates an expected call of AddFile.
func (mr *MockExporterMockRecorder) AddFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFile", reflect.TypeOf((*MockExporter)(nil).AddFile), arg0, arg1)
}

// AddFileWithDataSource mocks base method.
func (m *MockExporter) AddFileWithDataSource(arg0 string, arg1 func() ([]byte, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFileWithDataSource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileWithDataSource indicates an expected call of AddFileWithDataSource.
func (mr *MockExporterMockRecorder) AddFileWithDataSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileWithDataSource", reflect.TypeOf((*MockExporter)(nil).AddFileWithDataSource), arg0, arg1)
}

// Clean mocks base method.
func (m *MockExporter) Clean() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clean indicates an expected call of Clean.
func (mr *MockExporterMockRecorder) Clean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockExporter)(nil).Clean))
}

// CreateExportPackage mocks base method.
func (m *MockExporter) CreateExportPackage(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExportPackage", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExportPackage indicates an expected call of CreateExportPackage.
func (mr *MockExporterMockRecorder) CreateExportPackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExportPackage", reflect.TypeOf((*MockExporter)(nil).CreateExportPackage), arg0, arg1)
}

// GetTmpDir mocks base method.
func (m *MockExporter) GetTmpDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTmpDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTmpDir indicates an expected call of GetTmpDir.
func (mr *MockExporterMockRecorder) GetTmpDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTmpDir", reflect.TypeOf((*MockExporter)(nil).GetTmpDir))
}
