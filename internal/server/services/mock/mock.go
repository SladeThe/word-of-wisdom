// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SladeThe/word-of-wisdom/internal/server/services (interfaces: Challenge,WordOfWisdom)
//
// Generated by this command:
//
//	mockgen -package mock -destination mock/mock.go . Challenge,WordOfWisdom
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entities "github.com/SladeThe/word-of-wisdom/internal/common/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockChallenge is a mock of Challenge interface.
type MockChallenge struct {
	ctrl     *gomock.Controller
	recorder *MockChallengeMockRecorder
}

// MockChallengeMockRecorder is the mock recorder for MockChallenge.
type MockChallengeMockRecorder struct {
	mock *MockChallenge
}

// NewMockChallenge creates a new mock instance.
func NewMockChallenge(ctrl *gomock.Controller) *MockChallenge {
	mock := &MockChallenge{ctrl: ctrl}
	mock.recorder = &MockChallengeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChallenge) EXPECT() *MockChallengeMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockChallenge) Accept(arg0 context.Context, arg1 entities.ClientID) (entities.Challenge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0, arg1)
	ret0, _ := ret[0].(entities.Challenge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept.
func (mr *MockChallengeMockRecorder) Accept(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockChallenge)(nil).Accept), arg0, arg1)
}

// Solve mocks base method.
func (m *MockChallenge) Solve(arg0 context.Context, arg1 entities.ClientID, arg2 entities.Challenge, arg3 entities.Solution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Solve", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Solve indicates an expected call of Solve.
func (mr *MockChallengeMockRecorder) Solve(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Solve", reflect.TypeOf((*MockChallenge)(nil).Solve), arg0, arg1, arg2, arg3)
}

// MockWordOfWisdom is a mock of WordOfWisdom interface.
type MockWordOfWisdom struct {
	ctrl     *gomock.Controller
	recorder *MockWordOfWisdomMockRecorder
}

// MockWordOfWisdomMockRecorder is the mock recorder for MockWordOfWisdom.
type MockWordOfWisdomMockRecorder struct {
	mock *MockWordOfWisdom
}

// NewMockWordOfWisdom creates a new mock instance.
func NewMockWordOfWisdom(ctrl *gomock.Controller) *MockWordOfWisdom {
	mock := &MockWordOfWisdom{ctrl: ctrl}
	mock.recorder = &MockWordOfWisdomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWordOfWisdom) EXPECT() *MockWordOfWisdomMockRecorder {
	return m.recorder
}

// OneRandom mocks base method.
func (m *MockWordOfWisdom) OneRandom(arg0 context.Context) (entities.WordOfWisdom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneRandom", arg0)
	ret0, _ := ret[0].(entities.WordOfWisdom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OneRandom indicates an expected call of OneRandom.
func (mr *MockWordOfWisdomMockRecorder) OneRandom(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneRandom", reflect.TypeOf((*MockWordOfWisdom)(nil).OneRandom), arg0)
}
