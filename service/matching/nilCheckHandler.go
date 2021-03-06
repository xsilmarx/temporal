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

package matching

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
)

// Disable lint due to missing comments.
// Code generated by generate-adapter. DO NOT EDIT.

var _ matchingservice.MatchingServiceServer = (*NilCheckHandler)(nil)

type (
	// NilCheckHandler - gRPC handler interface for matchingservice
	NilCheckHandler struct {
		parentHandler matchingservice.MatchingServiceServer
	}
)

// NewNilCheckHandler creates a gRPC handler for the temporal matchingservice
func NewNilCheckHandler(
	parentHandler matchingservice.MatchingServiceServer,
) *NilCheckHandler {
	handler := &NilCheckHandler{
		parentHandler: parentHandler,
	}

	return handler
}

func (h *NilCheckHandler) PollForDecisionTask(ctx context.Context, request *matchingservice.PollForDecisionTaskRequest) (*matchingservice.PollForDecisionTaskResponse, error) {
	resp, err := h.parentHandler.PollForDecisionTask(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.PollForDecisionTaskResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) PollForActivityTask(ctx context.Context, request *matchingservice.PollForActivityTaskRequest) (*matchingservice.PollForActivityTaskResponse, error) {
	resp, err := h.parentHandler.PollForActivityTask(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.PollForActivityTaskResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) AddDecisionTask(ctx context.Context, request *matchingservice.AddDecisionTaskRequest) (*matchingservice.AddDecisionTaskResponse, error) {
	resp, err := h.parentHandler.AddDecisionTask(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.AddDecisionTaskResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) AddActivityTask(ctx context.Context, request *matchingservice.AddActivityTaskRequest) (*matchingservice.AddActivityTaskResponse, error) {
	resp, err := h.parentHandler.AddActivityTask(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.AddActivityTaskResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) QueryWorkflow(ctx context.Context, request *matchingservice.QueryWorkflowRequest) (*matchingservice.QueryWorkflowResponse, error) {
	resp, err := h.parentHandler.QueryWorkflow(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.QueryWorkflowResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) RespondQueryTaskCompleted(ctx context.Context, request *matchingservice.RespondQueryTaskCompletedRequest) (*matchingservice.RespondQueryTaskCompletedResponse, error) {
	resp, err := h.parentHandler.RespondQueryTaskCompleted(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.RespondQueryTaskCompletedResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) CancelOutstandingPoll(ctx context.Context, request *matchingservice.CancelOutstandingPollRequest) (*matchingservice.CancelOutstandingPollResponse, error) {
	resp, err := h.parentHandler.CancelOutstandingPoll(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.CancelOutstandingPollResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) DescribeTaskQueue(ctx context.Context, request *matchingservice.DescribeTaskQueueRequest) (*matchingservice.DescribeTaskQueueResponse, error) {
	resp, err := h.parentHandler.DescribeTaskQueue(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.DescribeTaskQueueResponse{}
	}
	return resp, err
}

func (h *NilCheckHandler) ListTaskQueuePartitions(ctx context.Context, request *matchingservice.ListTaskQueuePartitionsRequest) (*matchingservice.ListTaskQueuePartitionsResponse, error) {
	resp, err := h.parentHandler.ListTaskQueuePartitions(ctx, request)
	if resp == nil && err == nil {
		resp = &matchingservice.ListTaskQueuePartitionsResponse{}
	}
	return resp, err
}
