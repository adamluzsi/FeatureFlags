// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../../usecases/Storage.go

// Package cachespecs is a generated GoMock package.
package cachespecs

import (
	context "context"
	frameless "github.com/adamluzsi/frameless"
	gomock "github.com/golang/mock/gomock"
	release "github.com/toggler-io/toggler/domains/release"
	security "github.com/toggler-io/toggler/domains/security"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStorage) Create(ctx context.Context, ptr interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, ptr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockStorageMockRecorder) Create(ctx, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorage)(nil).Create), ctx, ptr)
}

// FindByID mocks base method
func (m *MockStorage) FindByID(ctx context.Context, ptr interface{}, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ptr, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockStorageMockRecorder) FindByID(ctx, ptr, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockStorage)(nil).FindByID), ctx, ptr, id)
}

// FindAll mocks base method
func (m *MockStorage) FindAll(ctx context.Context, T interface{}) frameless.Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, T)
	ret0, _ := ret[0].(frameless.Iterator)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockStorageMockRecorder) FindAll(ctx, T interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStorage)(nil).FindAll), ctx, T)
}

// Update mocks base method
func (m *MockStorage) Update(ctx context.Context, ptr interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ptr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStorageMockRecorder) Update(ctx, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStorage)(nil).Update), ctx, ptr)
}

// DeleteByID mocks base method
func (m *MockStorage) DeleteByID(ctx context.Context, T interface{}, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, T, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockStorageMockRecorder) DeleteByID(ctx, T, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockStorage)(nil).DeleteByID), ctx, T, id)
}

// DeleteAll mocks base method
func (m *MockStorage) DeleteAll(ctx context.Context, T interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx, T)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockStorageMockRecorder) DeleteAll(ctx, T interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockStorage)(nil).DeleteAll), ctx, T)
}

// FindReleaseFlagByName mocks base method
func (m *MockStorage) FindReleaseFlagByName(ctx context.Context, name string) (*release.Flag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindReleaseFlagByName", ctx, name)
	ret0, _ := ret[0].(*release.Flag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReleaseFlagByName indicates an expected call of FindReleaseFlagByName
func (mr *MockStorageMockRecorder) FindReleaseFlagByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReleaseFlagByName", reflect.TypeOf((*MockStorage)(nil).FindReleaseFlagByName), ctx, name)
}

// FindReleaseFlagsByName mocks base method
func (m *MockStorage) FindReleaseFlagsByName(ctx context.Context, names ...string) release.FlagEntries {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range names {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindReleaseFlagsByName", varargs...)
	ret0, _ := ret[0].(release.FlagEntries)
	return ret0
}

// FindReleaseFlagsByName indicates an expected call of FindReleaseFlagsByName
func (mr *MockStorageMockRecorder) FindReleaseFlagsByName(ctx interface{}, names ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, names...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReleaseFlagsByName", reflect.TypeOf((*MockStorage)(nil).FindReleaseFlagsByName), varargs...)
}

// FindReleaseFlagPilotByPilotExternalID mocks base method
func (m *MockStorage) FindReleaseFlagPilotByPilotExternalID(ctx context.Context, flagID, pilotExtID string) (*release.Pilot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindReleaseFlagPilotByPilotExternalID", ctx, flagID, pilotExtID)
	ret0, _ := ret[0].(*release.Pilot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReleaseFlagPilotByPilotExternalID indicates an expected call of FindReleaseFlagPilotByPilotExternalID
func (mr *MockStorageMockRecorder) FindReleaseFlagPilotByPilotExternalID(ctx, flagID, pilotExtID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReleaseFlagPilotByPilotExternalID", reflect.TypeOf((*MockStorage)(nil).FindReleaseFlagPilotByPilotExternalID), ctx, flagID, pilotExtID)
}

// FindPilotsByFeatureFlag mocks base method
func (m *MockStorage) FindPilotsByFeatureFlag(ctx context.Context, ff *release.Flag) release.PilotEntries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPilotsByFeatureFlag", ctx, ff)
	ret0, _ := ret[0].(release.PilotEntries)
	return ret0
}

// FindPilotsByFeatureFlag indicates an expected call of FindPilotsByFeatureFlag
func (mr *MockStorageMockRecorder) FindPilotsByFeatureFlag(ctx, ff interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPilotsByFeatureFlag", reflect.TypeOf((*MockStorage)(nil).FindPilotsByFeatureFlag), ctx, ff)
}

// FindPilotEntriesByExtID mocks base method
func (m *MockStorage) FindPilotEntriesByExtID(ctx context.Context, pilotExtID string) release.PilotEntries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPilotEntriesByExtID", ctx, pilotExtID)
	ret0, _ := ret[0].(release.PilotEntries)
	return ret0
}

// FindPilotEntriesByExtID indicates an expected call of FindPilotEntriesByExtID
func (mr *MockStorageMockRecorder) FindPilotEntriesByExtID(ctx, pilotExtID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPilotEntriesByExtID", reflect.TypeOf((*MockStorage)(nil).FindPilotEntriesByExtID), ctx, pilotExtID)
}

// FindReleaseAllowsByReleaseFlags mocks base method
func (m *MockStorage) FindReleaseAllowsByReleaseFlags(ctx context.Context, flags ...release.Flag) release.AllowEntries {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range flags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindReleaseAllowsByReleaseFlags", varargs...)
	ret0, _ := ret[0].(release.AllowEntries)
	return ret0
}

// FindReleaseAllowsByReleaseFlags indicates an expected call of FindReleaseAllowsByReleaseFlags
func (mr *MockStorageMockRecorder) FindReleaseAllowsByReleaseFlags(ctx interface{}, flags ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, flags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReleaseAllowsByReleaseFlags", reflect.TypeOf((*MockStorage)(nil).FindReleaseAllowsByReleaseFlags), varargs...)
}

// FindTokenBySHA512Hex mocks base method
func (m *MockStorage) FindTokenBySHA512Hex(ctx context.Context, sha512hex string) (*security.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTokenBySHA512Hex", ctx, sha512hex)
	ret0, _ := ret[0].(*security.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTokenBySHA512Hex indicates an expected call of FindTokenBySHA512Hex
func (mr *MockStorageMockRecorder) FindTokenBySHA512Hex(ctx, sha512hex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTokenBySHA512Hex", reflect.TypeOf((*MockStorage)(nil).FindTokenBySHA512Hex), ctx, sha512hex)
}

// Close mocks base method
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}