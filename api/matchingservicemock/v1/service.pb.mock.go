// Code generated by MockGen. DO NOT EDIT.
// Source: matchingservice/v1/service.pb.go

// Package matchingservicemock is a generated GoMock package.
package matchingservicemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	matchingservice "go.temporal.io/server/api/matchingservice/v1"
	grpc "google.golang.org/grpc"
)

// MockMatchingServiceClient is a mock of MatchingServiceClient interface.
type MockMatchingServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMatchingServiceClientMockRecorder
}

// MockMatchingServiceClientMockRecorder is the mock recorder for MockMatchingServiceClient.
type MockMatchingServiceClientMockRecorder struct {
	mock *MockMatchingServiceClient
}

// NewMockMatchingServiceClient creates a new mock instance.
func NewMockMatchingServiceClient(ctrl *gomock.Controller) *MockMatchingServiceClient {
	mock := &MockMatchingServiceClient{ctrl: ctrl}
	mock.recorder = &MockMatchingServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchingServiceClient) EXPECT() *MockMatchingServiceClientMockRecorder {
	return m.recorder
}

// PollForDecisionTask mocks base method.
func (m *MockMatchingServiceClient) PollForDecisionTask(ctx context.Context, in *matchingservice.PollForDecisionTaskRequest, opts ...grpc.CallOption) (*matchingservice.PollForDecisionTaskResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PollForDecisionTask", varargs...)
	ret0, _ := ret[0].(*matchingservice.PollForDecisionTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForDecisionTask indicates an expected call of PollForDecisionTask.
func (mr *MockMatchingServiceClientMockRecorder) PollForDecisionTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForDecisionTask", reflect.TypeOf((*MockMatchingServiceClient)(nil).PollForDecisionTask), varargs...)
}

// PollForActivityTask mocks base method.
func (m *MockMatchingServiceClient) PollForActivityTask(ctx context.Context, in *matchingservice.PollForActivityTaskRequest, opts ...grpc.CallOption) (*matchingservice.PollForActivityTaskResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PollForActivityTask", varargs...)
	ret0, _ := ret[0].(*matchingservice.PollForActivityTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForActivityTask indicates an expected call of PollForActivityTask.
func (mr *MockMatchingServiceClientMockRecorder) PollForActivityTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForActivityTask", reflect.TypeOf((*MockMatchingServiceClient)(nil).PollForActivityTask), varargs...)
}

// AddDecisionTask mocks base method.
func (m *MockMatchingServiceClient) AddDecisionTask(ctx context.Context, in *matchingservice.AddDecisionTaskRequest, opts ...grpc.CallOption) (*matchingservice.AddDecisionTaskResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddDecisionTask", varargs...)
	ret0, _ := ret[0].(*matchingservice.AddDecisionTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDecisionTask indicates an expected call of AddDecisionTask.
func (mr *MockMatchingServiceClientMockRecorder) AddDecisionTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDecisionTask", reflect.TypeOf((*MockMatchingServiceClient)(nil).AddDecisionTask), varargs...)
}

// AddActivityTask mocks base method.
func (m *MockMatchingServiceClient) AddActivityTask(ctx context.Context, in *matchingservice.AddActivityTaskRequest, opts ...grpc.CallOption) (*matchingservice.AddActivityTaskResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddActivityTask", varargs...)
	ret0, _ := ret[0].(*matchingservice.AddActivityTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddActivityTask indicates an expected call of AddActivityTask.
func (mr *MockMatchingServiceClientMockRecorder) AddActivityTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddActivityTask", reflect.TypeOf((*MockMatchingServiceClient)(nil).AddActivityTask), varargs...)
}

// QueryWorkflow mocks base method.
func (m *MockMatchingServiceClient) QueryWorkflow(ctx context.Context, in *matchingservice.QueryWorkflowRequest, opts ...grpc.CallOption) (*matchingservice.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryWorkflow", varargs...)
	ret0, _ := ret[0].(*matchingservice.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow.
func (mr *MockMatchingServiceClientMockRecorder) QueryWorkflow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockMatchingServiceClient)(nil).QueryWorkflow), varargs...)
}

// RespondQueryTaskCompleted mocks base method.
func (m *MockMatchingServiceClient) RespondQueryTaskCompleted(ctx context.Context, in *matchingservice.RespondQueryTaskCompletedRequest, opts ...grpc.CallOption) (*matchingservice.RespondQueryTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondQueryTaskCompleted", varargs...)
	ret0, _ := ret[0].(*matchingservice.RespondQueryTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondQueryTaskCompleted indicates an expected call of RespondQueryTaskCompleted.
func (mr *MockMatchingServiceClientMockRecorder) RespondQueryTaskCompleted(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondQueryTaskCompleted", reflect.TypeOf((*MockMatchingServiceClient)(nil).RespondQueryTaskCompleted), varargs...)
}

// CancelOutstandingPoll mocks base method.
func (m *MockMatchingServiceClient) CancelOutstandingPoll(ctx context.Context, in *matchingservice.CancelOutstandingPollRequest, opts ...grpc.CallOption) (*matchingservice.CancelOutstandingPollResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOutstandingPoll", varargs...)
	ret0, _ := ret[0].(*matchingservice.CancelOutstandingPollResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOutstandingPoll indicates an expected call of CancelOutstandingPoll.
func (mr *MockMatchingServiceClientMockRecorder) CancelOutstandingPoll(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOutstandingPoll", reflect.TypeOf((*MockMatchingServiceClient)(nil).CancelOutstandingPoll), varargs...)
}

// DescribeTaskQueue mocks base method.
func (m *MockMatchingServiceClient) DescribeTaskQueue(ctx context.Context, in *matchingservice.DescribeTaskQueueRequest, opts ...grpc.CallOption) (*matchingservice.DescribeTaskQueueResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTaskQueue", varargs...)
	ret0, _ := ret[0].(*matchingservice.DescribeTaskQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTaskQueue indicates an expected call of DescribeTaskQueue.
func (mr *MockMatchingServiceClientMockRecorder) DescribeTaskQueue(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskQueue", reflect.TypeOf((*MockMatchingServiceClient)(nil).DescribeTaskQueue), varargs...)
}

// ListTaskQueuePartitions mocks base method.
func (m *MockMatchingServiceClient) ListTaskQueuePartitions(ctx context.Context, in *matchingservice.ListTaskQueuePartitionsRequest, opts ...grpc.CallOption) (*matchingservice.ListTaskQueuePartitionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTaskQueuePartitions", varargs...)
	ret0, _ := ret[0].(*matchingservice.ListTaskQueuePartitionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTaskQueuePartitions indicates an expected call of ListTaskQueuePartitions.
func (mr *MockMatchingServiceClientMockRecorder) ListTaskQueuePartitions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTaskQueuePartitions", reflect.TypeOf((*MockMatchingServiceClient)(nil).ListTaskQueuePartitions), varargs...)
}

// MockMatchingServiceServer is a mock of MatchingServiceServer interface.
type MockMatchingServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMatchingServiceServerMockRecorder
}

// MockMatchingServiceServerMockRecorder is the mock recorder for MockMatchingServiceServer.
type MockMatchingServiceServerMockRecorder struct {
	mock *MockMatchingServiceServer
}

// NewMockMatchingServiceServer creates a new mock instance.
func NewMockMatchingServiceServer(ctrl *gomock.Controller) *MockMatchingServiceServer {
	mock := &MockMatchingServiceServer{ctrl: ctrl}
	mock.recorder = &MockMatchingServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchingServiceServer) EXPECT() *MockMatchingServiceServerMockRecorder {
	return m.recorder
}

// PollForDecisionTask mocks base method.
func (m *MockMatchingServiceServer) PollForDecisionTask(arg0 context.Context, arg1 *matchingservice.PollForDecisionTaskRequest) (*matchingservice.PollForDecisionTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollForDecisionTask", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.PollForDecisionTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForDecisionTask indicates an expected call of PollForDecisionTask.
func (mr *MockMatchingServiceServerMockRecorder) PollForDecisionTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForDecisionTask", reflect.TypeOf((*MockMatchingServiceServer)(nil).PollForDecisionTask), arg0, arg1)
}

// PollForActivityTask mocks base method.
func (m *MockMatchingServiceServer) PollForActivityTask(arg0 context.Context, arg1 *matchingservice.PollForActivityTaskRequest) (*matchingservice.PollForActivityTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollForActivityTask", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.PollForActivityTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollForActivityTask indicates an expected call of PollForActivityTask.
func (mr *MockMatchingServiceServerMockRecorder) PollForActivityTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollForActivityTask", reflect.TypeOf((*MockMatchingServiceServer)(nil).PollForActivityTask), arg0, arg1)
}

// AddDecisionTask mocks base method.
func (m *MockMatchingServiceServer) AddDecisionTask(arg0 context.Context, arg1 *matchingservice.AddDecisionTaskRequest) (*matchingservice.AddDecisionTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDecisionTask", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.AddDecisionTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDecisionTask indicates an expected call of AddDecisionTask.
func (mr *MockMatchingServiceServerMockRecorder) AddDecisionTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDecisionTask", reflect.TypeOf((*MockMatchingServiceServer)(nil).AddDecisionTask), arg0, arg1)
}

// AddActivityTask mocks base method.
func (m *MockMatchingServiceServer) AddActivityTask(arg0 context.Context, arg1 *matchingservice.AddActivityTaskRequest) (*matchingservice.AddActivityTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddActivityTask", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.AddActivityTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddActivityTask indicates an expected call of AddActivityTask.
func (mr *MockMatchingServiceServerMockRecorder) AddActivityTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddActivityTask", reflect.TypeOf((*MockMatchingServiceServer)(nil).AddActivityTask), arg0, arg1)
}

// QueryWorkflow mocks base method.
func (m *MockMatchingServiceServer) QueryWorkflow(arg0 context.Context, arg1 *matchingservice.QueryWorkflowRequest) (*matchingservice.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow.
func (mr *MockMatchingServiceServerMockRecorder) QueryWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockMatchingServiceServer)(nil).QueryWorkflow), arg0, arg1)
}

// RespondQueryTaskCompleted mocks base method.
func (m *MockMatchingServiceServer) RespondQueryTaskCompleted(arg0 context.Context, arg1 *matchingservice.RespondQueryTaskCompletedRequest) (*matchingservice.RespondQueryTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondQueryTaskCompleted", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.RespondQueryTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondQueryTaskCompleted indicates an expected call of RespondQueryTaskCompleted.
func (mr *MockMatchingServiceServerMockRecorder) RespondQueryTaskCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondQueryTaskCompleted", reflect.TypeOf((*MockMatchingServiceServer)(nil).RespondQueryTaskCompleted), arg0, arg1)
}

// CancelOutstandingPoll mocks base method.
func (m *MockMatchingServiceServer) CancelOutstandingPoll(arg0 context.Context, arg1 *matchingservice.CancelOutstandingPollRequest) (*matchingservice.CancelOutstandingPollResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOutstandingPoll", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.CancelOutstandingPollResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOutstandingPoll indicates an expected call of CancelOutstandingPoll.
func (mr *MockMatchingServiceServerMockRecorder) CancelOutstandingPoll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOutstandingPoll", reflect.TypeOf((*MockMatchingServiceServer)(nil).CancelOutstandingPoll), arg0, arg1)
}

// DescribeTaskQueue mocks base method.
func (m *MockMatchingServiceServer) DescribeTaskQueue(arg0 context.Context, arg1 *matchingservice.DescribeTaskQueueRequest) (*matchingservice.DescribeTaskQueueResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTaskQueue", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.DescribeTaskQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTaskQueue indicates an expected call of DescribeTaskQueue.
func (mr *MockMatchingServiceServerMockRecorder) DescribeTaskQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskQueue", reflect.TypeOf((*MockMatchingServiceServer)(nil).DescribeTaskQueue), arg0, arg1)
}

// ListTaskQueuePartitions mocks base method.
func (m *MockMatchingServiceServer) ListTaskQueuePartitions(arg0 context.Context, arg1 *matchingservice.ListTaskQueuePartitionsRequest) (*matchingservice.ListTaskQueuePartitionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTaskQueuePartitions", arg0, arg1)
	ret0, _ := ret[0].(*matchingservice.ListTaskQueuePartitionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTaskQueuePartitions indicates an expected call of ListTaskQueuePartitions.
func (mr *MockMatchingServiceServerMockRecorder) ListTaskQueuePartitions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTaskQueuePartitions", reflect.TypeOf((*MockMatchingServiceServer)(nil).ListTaskQueuePartitions), arg0, arg1)
}
