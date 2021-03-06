// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cli

import (
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/olekukonko/tablewriter"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	namespacepb "go.temporal.io/api/namespace/v1"
	replicationpb "go.temporal.io/api/replication/v1"
	"go.temporal.io/api/serviceerror"
	taskqueuepb "go.temporal.io/api/taskqueue/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/api/workflowservicemock/v1"
	sdkclient "go.temporal.io/sdk/client"
	sdkmocks "go.temporal.io/sdk/mocks"

	"go.temporal.io/server/api/adminservice/v1"
	"go.temporal.io/server/api/adminservicemock/v1"
	"go.temporal.io/server/common/payload"
	"go.temporal.io/server/common/payloads"
)

type cliAppSuite struct {
	suite.Suite
	app               *cli.App
	mockCtrl          *gomock.Controller
	frontendClient    *workflowservicemock.MockWorkflowServiceClient
	serverAdminClient *adminservicemock.MockAdminServiceClient
	sdkClient         *sdkmocks.Client
}

type clientFactoryMock struct {
	frontendClient    workflowservice.WorkflowServiceClient
	serverAdminClient adminservice.AdminServiceClient
	sdkClient         *sdkmocks.Client
}

func (m *clientFactoryMock) FrontendClient(c *cli.Context) workflowservice.WorkflowServiceClient {
	return m.frontendClient
}

func (m *clientFactoryMock) AdminClient(c *cli.Context) adminservice.AdminServiceClient {
	return m.serverAdminClient
}

func (m *clientFactoryMock) SDKClient(c *cli.Context, namespace string) sdkclient.Client {
	return m.sdkClient
}

var commands = []string{
	"namespace", "n",
	"workflow", "wf",
	"taskqueue", "tq",
}

var cliTestNamespace = "cli-test-namespace"

func TestCLIAppSuite(t *testing.T) {
	s := new(cliAppSuite)
	suite.Run(t, s)
}

func (s *cliAppSuite) SetupSuite() {
	s.app = NewCliApp()
}

func (s *cliAppSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())

	s.frontendClient = workflowservicemock.NewMockWorkflowServiceClient(s.mockCtrl)
	s.serverAdminClient = adminservicemock.NewMockAdminServiceClient(s.mockCtrl)
	s.sdkClient = &sdkmocks.Client{}
	SetFactory(&clientFactoryMock{
		frontendClient:    s.frontendClient,
		serverAdminClient: s.serverAdminClient,
		sdkClient:         s.sdkClient,
	})
}

func (s *cliAppSuite) TearDownTest() {
	s.mockCtrl.Finish() // assert mock’s expectations
}

func (s *cliAppSuite) RunErrorExitCode(arguments []string) int {
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()
	var errorCode int
	osExit = func(code int) {
		errorCode = code
	}
	s.NoError(s.app.Run(arguments))
	return errorCode
}

func (s *cliAppSuite) TestAppCommands() {
	for _, test := range commands {
		cmd := s.app.Command(test)
		s.NotNil(cmd)
	}
}

func (s *cliAppSuite) TestNamespaceRegister_LocalNamespace() {
	s.frontendClient.EXPECT().RegisterNamespace(gomock.Any(), gomock.Any()).Return(nil, nil)
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "register", "--global_namespace", "false"})
	s.Equal(0, errorCode)
}

func (s *cliAppSuite) TestNamespaceRegister_GlobalNamespace() {
	s.frontendClient.EXPECT().RegisterNamespace(gomock.Any(), gomock.Any()).Return(nil, nil)
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "register", "--global_namespace", "true"})
	s.Equal(0, errorCode)
}

func (s *cliAppSuite) TestNamespaceRegister_NamespaceExist() {
	s.frontendClient.EXPECT().RegisterNamespace(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewNamespaceAlreadyExists(""))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "register", "--global_namespace", "true"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestNamespaceRegister_Failed() {
	s.frontendClient.EXPECT().RegisterNamespace(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "register", "--global_namespace", "true"})
	s.Equal(1, errorCode)
}

var describeNamespaceResponseServer = &workflowservice.DescribeNamespaceResponse{
	NamespaceInfo: &namespacepb.NamespaceInfo{
		Name:        "test-namespace",
		Description: "a test namespace",
		OwnerEmail:  "test@uber.com",
	},
	Config: &namespacepb.NamespaceConfig{
		WorkflowExecutionRetentionPeriodInDays: 3,
		EmitMetric:                             &types.BoolValue{Value: true},
	},
	ReplicationConfig: &replicationpb.NamespaceReplicationConfig{
		ActiveClusterName: "active",
		Clusters: []*replicationpb.ClusterReplicationConfig{
			{
				ClusterName: "active",
			},
			{
				ClusterName: "standby",
			},
		},
	},
}

func (s *cliAppSuite) TestNamespaceUpdate() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, nil).Times(2)
	s.frontendClient.EXPECT().UpdateNamespace(gomock.Any(), gomock.Any()).Return(nil, nil).Times(2)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "namespace", "update"})
	s.Nil(err)
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "namespace", "update", "--desc", "another desc", "--oe", "another@uber.com", "--rd", "1"})
	s.Nil(err)
}

func (s *cliAppSuite) TestNamespaceUpdate_NamespaceNotExist() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, nil)
	s.frontendClient.EXPECT().UpdateNamespace(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewNotFound(""))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "update"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestNamespaceUpdate_ActiveClusterFlagNotSet_NamespaceNotExist() {
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewNotFound(""))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "update"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestNamespaceUpdate_Failed() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, nil)
	s.frontendClient.EXPECT().UpdateNamespace(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "update"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestNamespaceDescribe() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, nil)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "namespace", "describe"})
	s.Nil(err)
}

func (s *cliAppSuite) TestNamespaceDescribe_NamespaceNotExist() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, serviceerror.NewNotFound(""))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "describe"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestNamespaceDescribe_Failed() {
	resp := describeNamespaceResponseServer
	s.frontendClient.EXPECT().DescribeNamespace(gomock.Any(), gomock.Any()).Return(resp, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "namespace", "describe"})
	s.Equal(1, errorCode)
}

var (
	eventType = enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED

	getWorkflowExecutionHistoryResponse = &workflowservice.GetWorkflowExecutionHistoryResponse{
		History: &historypb.History{
			Events: []*historypb.HistoryEvent{
				{
					EventType: eventType,
					Attributes: &historypb.HistoryEvent_WorkflowExecutionStartedEventAttributes{WorkflowExecutionStartedEventAttributes: &historypb.WorkflowExecutionStartedEventAttributes{
						WorkflowType:               &commonpb.WorkflowType{Name: "TestWorkflow"},
						TaskQueue:                  &taskqueuepb.TaskQueue{Name: "taskQueue"},
						WorkflowRunTimeoutSeconds:  60,
						WorkflowTaskTimeoutSeconds: 10,
						Identity:                   "tester",
					}},
				},
			},
		},
		NextPageToken: nil,
	}
)

func (s *cliAppSuite) TestShowHistory() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "show", "-w", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestShowHistoryWithID() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "showid", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestShowHistory_PrintRawTime() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "show", "-w", "wid", "-prt"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestShowHistory_PrintDateTime() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "show", "-w", "wid", "-pdt"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestStartWorkflow() {
	resp := &workflowservice.StartWorkflowExecutionResponse{RunId: uuid.New()}
	s.frontendClient.EXPECT().StartWorkflowExecution(gomock.Any(), gomock.Any()).Return(resp, nil).Times(2)
	// start with wid
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "start", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "-w", "wid", "-wrp", "AllowDuplicateFailedOnly"})
	s.Nil(err)
	// start without wid
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "start", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "-wrp", "AllowDuplicateFailedOnly"})
	s.Nil(err)
}

func (s *cliAppSuite) TestStartWorkflow_Failed() {
	resp := &workflowservice.StartWorkflowExecutionResponse{RunId: uuid.New()}
	s.frontendClient.EXPECT().StartWorkflowExecution(gomock.Any(), gomock.Any()).Return(resp, serviceerror.NewInvalidArgument("faked error"))
	// start with wid
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "start", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "-w", "wid"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestRunWorkflow() {
	resp := &workflowservice.StartWorkflowExecutionResponse{RunId: uuid.New()}
	s.frontendClient.EXPECT().StartWorkflowExecution(gomock.Any(), gomock.Any()).Return(resp, nil).Times(2)
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", mock.Anything, mock.Anything, mock.Anything).Return(historyEventIterator()).Once()

	// start with wid
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "run", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "-w", "wid", "wrp", "2"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())

	s.sdkClient.On("GetWorkflowHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	// start without wid
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "run", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "wrp", "2"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestRunWorkflow_Failed() {
	resp := &workflowservice.StartWorkflowExecutionResponse{RunId: uuid.New()}
	s.frontendClient.EXPECT().StartWorkflowExecution(gomock.Any(), gomock.Any()).Return(resp, serviceerror.NewInvalidArgument("faked error"))
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	// start with wid
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "run", "-tq", "testTaskQueue", "-wt", "testWorkflowType", "-et", "60", "-w", "wid"})
	s.Equal(1, errorCode)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestTerminateWorkflow() {
	s.sdkClient.On("TerminateWorkflow", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "terminate", "-w", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestTerminateWorkflow_Failed() {
	s.sdkClient.On("TerminateWorkflow", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(serviceerror.NewInvalidArgument("faked error")).Once()
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "terminate", "-w", "wid"})
	s.Equal(1, errorCode)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestCancelWorkflow() {
	s.sdkClient.On("CancelWorkflow", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "cancel", "-w", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestCancelWorkflow_Failed() {
	s.sdkClient.On("CancelWorkflow", mock.Anything, mock.Anything, mock.Anything).Return(serviceerror.NewInvalidArgument("faked error")).Once()
	//s.frontendClient.EXPECT().RequestCancelWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "cancel", "-w", "wid"})
	s.Equal(1, errorCode)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestSignalWorkflow() {
	s.frontendClient.EXPECT().SignalWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil, nil)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "signal", "-w", "wid", "-n", "signal-name"})
	s.Nil(err)
}

func (s *cliAppSuite) TestSignalWorkflow_Failed() {
	s.frontendClient.EXPECT().SignalWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "signal", "-w", "wid", "-n", "signal-name"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestQueryWorkflow() {
	resp := &workflowservice.QueryWorkflowResponse{
		QueryResult: payloads.EncodeString("query-result"),
	}
	s.frontendClient.EXPECT().QueryWorkflow(gomock.Any(), gomock.Any()).Return(resp, nil)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "query", "-w", "wid", "-qt", "query-type-test"})
	s.Nil(err)
}

func (s *cliAppSuite) TestQueryWorkflowUsingStackTrace() {
	resp := &workflowservice.QueryWorkflowResponse{
		QueryResult: payloads.EncodeString("query-result"),
	}
	s.frontendClient.EXPECT().QueryWorkflow(gomock.Any(), gomock.Any()).Return(resp, nil)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "stack", "-w", "wid"})
	s.Nil(err)
}

func (s *cliAppSuite) TestQueryWorkflow_Failed() {
	resp := &workflowservice.QueryWorkflowResponse{
		QueryResult: payloads.EncodeString("query-result"),
	}
	s.frontendClient.EXPECT().QueryWorkflow(gomock.Any(), gomock.Any()).Return(resp, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "workflow", "query", "-w", "wid", "-qt", "query-type-test"})
	s.Equal(1, errorCode)
}

var (
	status = enumspb.WORKFLOW_EXECUTION_STATUS_COMPLETED

	listClosedWorkflowExecutionsResponse = &workflowservice.ListClosedWorkflowExecutionsResponse{
		Executions: []*workflowpb.WorkflowExecutionInfo{
			{
				Execution: &commonpb.WorkflowExecution{
					WorkflowId: "test-list-workflow-id",
					RunId:      uuid.New(),
				},
				Type: &commonpb.WorkflowType{
					Name: "test-list-workflow-type",
				},
				StartTime:     &types.Int64Value{Value: time.Now().UnixNano()},
				CloseTime:     &types.Int64Value{Value: time.Now().Add(time.Hour).UnixNano()},
				Status:        status,
				HistoryLength: 12,
			},
		},
	}

	listOpenWorkflowExecutionsResponse = &workflowservice.ListOpenWorkflowExecutionsResponse{
		Executions: []*workflowpb.WorkflowExecutionInfo{
			{
				Execution: &commonpb.WorkflowExecution{
					WorkflowId: "test-list-open-workflow-id",
					RunId:      uuid.New(),
				},
				Type: &commonpb.WorkflowType{
					Name: "test-list-open-workflow-type",
				},
				StartTime:     &types.Int64Value{Value: time.Now().UnixNano()},
				CloseTime:     &types.Int64Value{Value: time.Now().Add(time.Hour).UnixNano()},
				HistoryLength: 12,
			},
		},
	}
)

func (s *cliAppSuite) TestListWorkflow() {
	s.sdkClient.On("ListClosedWorkflow", mock.Anything, mock.Anything).Return(listClosedWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_WithWorkflowID() {
	s.sdkClient.On("ListClosedWorkflow", mock.Anything, mock.Anything).Return(listClosedWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-wid", "nothing"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_WithWorkflowType() {
	s.sdkClient.On("ListClosedWorkflow", mock.Anything, mock.Anything).Return(listClosedWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-wt", "no-type"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_PrintDateTime() {
	s.sdkClient.On("ListClosedWorkflow", mock.Anything, mock.Anything).Return(listClosedWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-pdt"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_PrintRawTime() {
	s.sdkClient.On("ListClosedWorkflow", mock.Anything, mock.Anything).Return(listClosedWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-prt"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_Open() {
	s.sdkClient.On("ListOpenWorkflow", mock.Anything, mock.Anything).Return(listOpenWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-op"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_Open_WithWorkflowID() {
	s.sdkClient.On("ListOpenWorkflow", mock.Anything, mock.Anything).Return(listOpenWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-op", "-wid", "nothing"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListWorkflow_Open_WithWorkflowType() {
	s.sdkClient.On("ListOpenWorkflow", mock.Anything, mock.Anything).Return(listOpenWorkflowExecutionsResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "list", "-op", "-wt", "no-type"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestListArchivedWorkflow() {
	s.sdkClient.On("ListArchivedWorkflow", mock.Anything, mock.Anything).Return(&workflowservice.ListArchivedWorkflowExecutionsResponse{}, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "listarchived", "-q", "some query string", "--ps", "200", "--all"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestCountWorkflow() {
	s.sdkClient.On("CountWorkflow", mock.Anything, mock.Anything).Return(&workflowservice.CountWorkflowExecutionsResponse{}, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "count"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())

	s.sdkClient.On("CountWorkflow", mock.Anything, mock.Anything).Return(&workflowservice.CountWorkflowExecutionsResponse{}, nil).Once()
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "count", "-q", "'CloseTime = missing'"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

var describeTaskQueueResponse = &workflowservice.DescribeTaskQueueResponse{
	Pollers: []*taskqueuepb.PollerInfo{
		{
			LastAccessTime: time.Now().UnixNano(),
			Identity:       "tester",
		},
	},
}

func (s *cliAppSuite) TestAdminDescribeWorkflow() {
	resp := &adminservice.DescribeWorkflowExecutionResponse{
		ShardId:                "test-shard-id",
		HistoryAddr:            "ip:port",
		MutableStateInDatabase: `{"ExecutionInfo":{"BranchToken":"ChBNWvyipehOuYvioA1u+suwEhDyawZ9XsdN6Liiof+Novu5"}}`,
	}

	s.serverAdminClient.EXPECT().DescribeWorkflowExecution(gomock.Any(), gomock.Any()).Return(resp, nil)
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "admin", "wf", "describe", "-w", "test-wf-id"})
	s.Nil(err)
}

func (s *cliAppSuite) TestAdminDescribeWorkflow_Failed() {
	s.serverAdminClient.EXPECT().DescribeWorkflowExecution(gomock.Any(), gomock.Any()).Return(nil, serviceerror.NewInvalidArgument("faked error"))
	errorCode := s.RunErrorExitCode([]string{"", "--ns", cliTestNamespace, "admin", "wf", "describe", "-w", "test-wf-id"})
	s.Equal(1, errorCode)
}

func (s *cliAppSuite) TestAdminAddSearchAttribute() {
	request := &adminservice.AddSearchAttributeRequest{
		SearchAttribute: map[string]enumspb.IndexedValueType{
			"testKey": enumspb.IndexedValueType(2),
		},
	}
	s.serverAdminClient.EXPECT().AddSearchAttribute(gomock.Any(), request).Times(1)

	err := s.app.Run([]string{"", "--auto_confirm", "--ns", cliTestNamespace, "admin", "cl", "asa", "--search_attr_key", "testKey", "--search_attr_type", "keyword"})
	s.Nil(err)
}

func (s *cliAppSuite) TestDescribeTaskQueue() {
	s.sdkClient.On("DescribeTaskQueue", mock.Anything, mock.Anything, mock.Anything).Return(describeTaskQueueResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "taskqueue", "describe", "-tq", "test-taskQueue"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestDescribeTaskQueue_Activity() {
	s.sdkClient.On("DescribeTaskQueue", mock.Anything, mock.Anything, mock.Anything).Return(describeTaskQueueResponse, nil).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "taskqueue", "describe", "-tq", "test-taskQueue", "-tqt", "activity"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestObserveWorkflow() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "observe", "-w", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())

	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "observe", "-w", "wid", "-sd"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestObserveWorkflowWithID() {
	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err := s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "observeid", "wid"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())

	s.sdkClient.On("GetWorkflowHistory", mock.Anything, "wid", "", mock.Anything, mock.Anything).Return(historyEventIterator()).Once()
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "workflow", "observeid", "wid", "-sd"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

// TestParseTime tests the parsing of date argument in UTC and UnixNano formats
func (s *cliAppSuite) TestParseTime() {
	s.Equal(int64(100), parseTime("", 100, time.Now()))
	s.Equal(int64(1528383845000000000), parseTime("2018-06-07T15:04:05+00:00", 0, time.Now()))
	s.Equal(int64(1528383845000000000), parseTime("1528383845000000000", 0, time.Now()))
}

// TestParseTimeDateRange tests the parsing of date argument in time range format, N<duration>
// where N is the integral multiplier, and duration can be second/minute/hour/day/week/month/year
func (s *cliAppSuite) TestParseTimeDateRange() {
	now := time.Now()
	tests := []struct {
		timeStr  string // input
		defVal   int64  // input
		expected int64  // expected unix nano (approx)
	}{
		{
			timeStr:  "1s",
			defVal:   int64(0),
			expected: now.Add(-time.Second).UnixNano(),
		},
		{
			timeStr:  "100second",
			defVal:   int64(0),
			expected: now.Add(-100 * time.Second).UnixNano(),
		},
		{
			timeStr:  "2m",
			defVal:   int64(0),
			expected: now.Add(-2 * time.Minute).UnixNano(),
		},
		{
			timeStr:  "200minute",
			defVal:   int64(0),
			expected: now.Add(-200 * time.Minute).UnixNano(),
		},
		{
			timeStr:  "3h",
			defVal:   int64(0),
			expected: now.Add(-3 * time.Hour).UnixNano(),
		},
		{
			timeStr:  "1000hour",
			defVal:   int64(0),
			expected: now.Add(-1000 * time.Hour).UnixNano(),
		},
		{
			timeStr:  "5d",
			defVal:   int64(0),
			expected: now.Add(-5 * day).UnixNano(),
		},
		{
			timeStr:  "25day",
			defVal:   int64(0),
			expected: now.Add(-25 * day).UnixNano(),
		},
		{
			timeStr:  "5w",
			defVal:   int64(0),
			expected: now.Add(-5 * week).UnixNano(),
		},
		{
			timeStr:  "52week",
			defVal:   int64(0),
			expected: now.Add(-52 * week).UnixNano(),
		},
		{
			timeStr:  "3M",
			defVal:   int64(0),
			expected: now.Add(-3 * month).UnixNano(),
		},
		{
			timeStr:  "6month",
			defVal:   int64(0),
			expected: now.Add(-6 * month).UnixNano(),
		},
		{
			timeStr:  "1y",
			defVal:   int64(0),
			expected: now.Add(-year).UnixNano(),
		},
		{
			timeStr:  "7year",
			defVal:   int64(0),
			expected: now.Add(-7 * year).UnixNano(),
		},
		{
			timeStr:  "100y", // epoch time will be returned as that's the minimum unix timestamp possible
			defVal:   int64(0),
			expected: time.Unix(0, 0).UnixNano(),
		},
	}
	const delta = int64(5 * time.Millisecond)
	for _, te := range tests {
		parsedTime := parseTime(te.timeStr, te.defVal, now)
		s.True(te.expected <= parsedTime, "Case: %s. %d must be less or equal than parsed %d", te.timeStr, te.expected, parsedTime)
		s.True(te.expected+delta >= parsedTime, "Case: %s. %d must be greater or equal than parsed %d", te.timeStr, te.expected, parsedTime)
	}
}

func (s *cliAppSuite) TestBreakLongWords() {
	s.Equal("111 222 333 4", breakLongWords("1112223334", 3))
	s.Equal("111 2 223", breakLongWords("1112 223", 3))
	s.Equal("11 122 23", breakLongWords("11 12223", 3))
	s.Equal("111", breakLongWords("111", 3))
	s.Equal("", breakLongWords("", 3))
	s.Equal("111  222", breakLongWords("111 222", 3))
}

func (s *cliAppSuite) TestAnyToString() {
	arg := strings.Repeat("LongText", 80)
	event := &historypb.HistoryEvent{
		EventId:   1,
		EventType: eventType,
		Attributes: &historypb.HistoryEvent_WorkflowExecutionStartedEventAttributes{WorkflowExecutionStartedEventAttributes: &historypb.WorkflowExecutionStartedEventAttributes{
			WorkflowType:               &commonpb.WorkflowType{Name: "helloworldWorkflow"},
			TaskQueue:                  &taskqueuepb.TaskQueue{Name: "taskQueue"},
			WorkflowRunTimeoutSeconds:  60,
			WorkflowTaskTimeoutSeconds: 10,
			Identity:                   "tester",
			Input:                      payloads.EncodeString(arg),
		}},
	}
	res := anyToString(event, false, defaultMaxFieldLength)
	ss, l := tablewriter.WrapString(res, 10)
	s.Equal(6, len(ss))
	s.Equal(131, l)
}

func (s *cliAppSuite) TestAnyToString_DecodeMapValues() {
	fields := map[string]*commonpb.Payload{
		"TestKey": payload.EncodeString("testValue"),
	}
	execution := &workflowpb.WorkflowExecutionInfo{
		Status: enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		Memo:   &commonpb.Memo{Fields: fields},
	}
	s.Equal("{Status:Running, HistoryLength:0, ExecutionTime:0, Memo:{Fields:map{TestKey:\"testValue\"}}}", anyToString(execution, true, 0))

	fields["TestKey2"] = payload.EncodeString(`anotherTestValue`)
	execution.Memo = &commonpb.Memo{Fields: fields}
	got := anyToString(execution, true, 0)
	expected := got == "{Status:Running, HistoryLength:0, ExecutionTime:0, Memo:{Fields:map{TestKey2:\"anotherTestValue\", TestKey:\"testValue\"}}}" ||
		got == "{Status:Running, HistoryLength:0, ExecutionTime:0, Memo:{Fields:map{TestKey:\"testValue\", TestKey2:\"anotherTestValue\"}}}"
	s.True(expected)
}

func (s *cliAppSuite) TestIsAttributeName() {
	s.True(isAttributeName("WorkflowExecutionStartedEventAttributes"))
	s.False(isAttributeName("workflowExecutionStartedEventAttributes"))
}

func (s *cliAppSuite) TestGetSearchAttributes() {
	s.sdkClient.On("GetSearchAttributes", mock.Anything).Return(&workflowservice.GetSearchAttributesResponse{}, nil).Once()
	err := s.app.Run([]string{"", "cluster", "get-search-attr"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())

	s.sdkClient.On("GetSearchAttributes", mock.Anything).Return(&workflowservice.GetSearchAttributesResponse{}, nil).Once()
	err = s.app.Run([]string{"", "--ns", cliTestNamespace, "cluster", "get-search-attr"})
	s.Nil(err)
	s.sdkClient.AssertExpectations(s.T())
}

func (s *cliAppSuite) TestParseBool() {
	res, err := parseBool("true")
	s.NoError(err)
	s.True(res)

	res, err = parseBool("false")
	s.NoError(err)
	s.False(res)

	for _, v := range []string{"True, TRUE, False, FALSE, T, F"} {
		res, err = parseBool(v)
		s.Error(err)
		s.False(res)
	}
}

func (s *cliAppSuite) TestConvertStringToRealType() {
	var res interface{}

	// int
	res = convertStringToRealType("1")
	s.Equal(int64(1), res)

	// bool
	res = convertStringToRealType("true")
	s.Equal(true, res)
	res = convertStringToRealType("false")
	s.Equal(false, res)

	// double
	res = convertStringToRealType("1.0")
	s.Equal(float64(1.0), res)

	// datetime
	res = convertStringToRealType("2019-01-01T01:01:01Z")
	s.Equal(time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC), res)

	// array
	res = convertStringToRealType(`["a", "b", "c"]`)
	s.Equal([]interface{}{"a", "b", "c"}, res)

	// string
	res = convertStringToRealType("test string")
	s.Equal("test string", res)
}

func (s *cliAppSuite) TestConvertArray() {
	t1, _ := time.Parse(defaultDateTimeFormat, "2019-06-07T16:16:34-08:00")
	t2, _ := time.Parse(defaultDateTimeFormat, "2019-06-07T17:16:34-08:00")
	testCases := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{
			name:     "string",
			input:    `["a", "b", "c"]`,
			expected: []interface{}{"a", "b", "c"},
		},
		{
			name:     "int",
			input:    `[1, 2, 3]`,
			expected: []interface{}{"1", "2", "3"},
		},
		{
			name:     "double",
			input:    `[1.1, 2.2, 3.3]`,
			expected: []interface{}{"1.1", "2.2", "3.3"},
		},
		{
			name:     "bool",
			input:    `["true", "false"]`,
			expected: []interface{}{"true", "false"},
		},
		{
			name:     "datetime",
			input:    `["2019-06-07T16:16:34-08:00", "2019-06-07T17:16:34-08:00"]`,
			expected: []interface{}{t1, t2},
		},
	}
	for _, testCase := range testCases {
		res, err := parseArray(testCase.input)
		s.Nil(err)
		s.Equal(testCase.expected, res)
	}

	testCases2 := []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:  "not array",
			input: "normal string",
		},
		{
			name:  "empty string",
			input: "",
		},
		{
			name:  "not json array",
			input: "[a, b, c]",
		},
	}
	for _, testCase := range testCases2 {
		res, err := parseArray(testCase.input)
		s.NotNil(err)
		s.Nil(res)
	}
}

func historyEventIterator() sdkclient.HistoryEventIterator {
	iteratorMock := &sdkmocks.HistoryEventIterator{}

	counter := 0
	hasNextFn := func() bool {
		if counter == 0 {
			return true
		} else {
			return false
		}
	}

	nextFn := func() *historypb.HistoryEvent {
		if counter == 0 {
			event := &historypb.HistoryEvent{
				EventType: eventType,
				Attributes: &historypb.HistoryEvent_WorkflowExecutionStartedEventAttributes{WorkflowExecutionStartedEventAttributes: &historypb.WorkflowExecutionStartedEventAttributes{
					WorkflowType:               &commonpb.WorkflowType{Name: "TestWorkflow"},
					TaskQueue:                  &taskqueuepb.TaskQueue{Name: "taskQueue"},
					WorkflowRunTimeoutSeconds:  60,
					WorkflowTaskTimeoutSeconds: 10,
					Identity:                   "tester",
				}},
			}
			counter++
			return event
		} else {
			return nil
		}
	}

	iteratorMock.On("HasNext").Return(hasNextFn).Twice()
	iteratorMock.On("Next").Return(nextFn, nil).Once()

	return iteratorMock
}

func workflowRun() sdkclient.WorkflowRun {
	workflowRunMock := &sdkmocks.WorkflowRun{}

	workflowRunMock.On("GetRunID").Return(uuid.New()).Maybe()
	workflowRunMock.On("GetID").Return(uuid.New()).Maybe()

	return workflowRunMock
}
