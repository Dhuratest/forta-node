// Code generated by MockGen. DO NOT EDIT.
// Source: clients/agentgrpc/client.go

// Package mock_agentgrpc is a generated GoMock package.
package mock_agentgrpc

import (
	context "context"
	reflect "reflect"

	protocol "github.com/forta-network/forta-core-go/protocol"
	agentgrpc "github.com/forta-network/forta-node/clients/agentgrpc"
	config "github.com/forta-network/forta-node/config"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// DialWithRetry mocks base method.
func (m *MockClient) DialWithRetry(arg0 config.AgentConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialWithRetry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DialWithRetry indicates an expected call of DialWithRetry.
func (mr *MockClientMockRecorder) DialWithRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialWithRetry", reflect.TypeOf((*MockClient)(nil).DialWithRetry), arg0)
}

// EvaluateAlert mocks base method.
func (m *MockClient) EvaluateAlert(ctx context.Context, in *protocol.EvaluateAlertRequest, opts ...grpc.CallOption) (*protocol.EvaluateAlertResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvaluateAlert", varargs...)
	ret0, _ := ret[0].(*protocol.EvaluateAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateAlert indicates an expected call of EvaluateAlert.
func (mr *MockClientMockRecorder) EvaluateAlert(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateAlert", reflect.TypeOf((*MockClient)(nil).EvaluateAlert), varargs...)
}

// EvaluateBlock mocks base method.
func (m *MockClient) EvaluateBlock(ctx context.Context, in *protocol.EvaluateBlockRequest, opts ...grpc.CallOption) (*protocol.EvaluateBlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvaluateBlock", varargs...)
	ret0, _ := ret[0].(*protocol.EvaluateBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateBlock indicates an expected call of EvaluateBlock.
func (mr *MockClientMockRecorder) EvaluateBlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateBlock", reflect.TypeOf((*MockClient)(nil).EvaluateBlock), varargs...)
}

// EvaluateTx mocks base method.
func (m *MockClient) EvaluateTx(ctx context.Context, in *protocol.EvaluateTxRequest, opts ...grpc.CallOption) (*protocol.EvaluateTxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvaluateTx", varargs...)
	ret0, _ := ret[0].(*protocol.EvaluateTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateTx indicates an expected call of EvaluateTx.
func (mr *MockClientMockRecorder) EvaluateTx(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateTx", reflect.TypeOf((*MockClient)(nil).EvaluateTx), varargs...)
}

// Initialize mocks base method.
func (m *MockClient) Initialize(ctx context.Context, in *protocol.InitializeRequest, opts ...grpc.CallOption) (*protocol.InitializeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Initialize", varargs...)
	ret0, _ := ret[0].(*protocol.InitializeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Initialize indicates an expected call of Initialize.
func (mr *MockClientMockRecorder) Initialize(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize), varargs...)
}

// Invoke mocks base method.
func (m *MockClient) Invoke(ctx context.Context, method agentgrpc.Method, in, out interface{}, opts ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, method, in, out}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Invoke", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invoke indicates an expected call of Invoke.
func (mr *MockClientMockRecorder) Invoke(ctx, method, in, out interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, method, in, out}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockClient)(nil).Invoke), varargs...)
}
