// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Approval_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RollbackApprovalInstance(ctx, &lark.RollbackApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddApprovalInstanceSign(ctx, &lark.AddApprovalInstanceSignReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalTask(ctx, &lark.SearchApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalUserTaskList(ctx, &lark.GetApprovalUserTaskListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalCarbonCopy(ctx, &lark.SearchApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.PreviewApprovalInstance(ctx, &lark.PreviewApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApprovalMessage(ctx, &lark.UpdateApprovalMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SubscribeApprovalSubscription(ctx, &lark.SubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UnsubscribeApprovalSubscription(ctx, &lark.UnsubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalExternalList(ctx, &lark.GetApprovalExternalListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SendApprovalMessage(ctx, &lark.SendApprovalMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalGetApproval(func(ctx context.Context, request *lark.GetApprovalReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApproval()

			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalGetApprovalInstanceList(func(ctx context.Context, request *lark.GetApprovalInstanceListReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalInstanceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalInstanceList()

			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalGetApprovalInstance(func(ctx context.Context, request *lark.GetApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalInstance()

			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalCreateApprovalInstance(func(ctx context.Context, request *lark.CreateApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.CreateApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCreateApprovalInstance()

			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalApproveApprovalInstance(func(ctx context.Context, request *lark.ApproveApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.ApproveApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalApproveApprovalInstance()

			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalRejectApprovalInstance(func(ctx context.Context, request *lark.RejectApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.RejectApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalRejectApprovalInstance()

			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalTransferApprovalInstance(func(ctx context.Context, request *lark.TransferApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.TransferApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalTransferApprovalInstance()

			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalRollbackApprovalInstance(func(ctx context.Context, request *lark.RollbackApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.RollbackApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalRollbackApprovalInstance()

			_, _, err := moduleCli.RollbackApprovalInstance(ctx, &lark.RollbackApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalCancelApprovalInstance(func(ctx context.Context, request *lark.CancelApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.CancelApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCancelApprovalInstance()

			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalSearchApprovalInstance(func(ctx context.Context, request *lark.SearchApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.SearchApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSearchApprovalInstance()

			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalAddApprovalInstanceSign(func(ctx context.Context, request *lark.AddApprovalInstanceSignReq, options ...lark.MethodOptionFunc) (*lark.AddApprovalInstanceSignResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalAddApprovalInstanceSign()

			_, _, err := moduleCli.AddApprovalInstanceSign(ctx, &lark.AddApprovalInstanceSignReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalUploadApprovalFile(func(ctx context.Context, request *lark.UploadApprovalFileReq, options ...lark.MethodOptionFunc) (*lark.UploadApprovalFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalUploadApprovalFile()

			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalSearchApprovalTask(func(ctx context.Context, request *lark.SearchApprovalTaskReq, options ...lark.MethodOptionFunc) (*lark.SearchApprovalTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSearchApprovalTask()

			_, _, err := moduleCli.SearchApprovalTask(ctx, &lark.SearchApprovalTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalGetApprovalUserTaskList(func(ctx context.Context, request *lark.GetApprovalUserTaskListReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalUserTaskListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalUserTaskList()

			_, _, err := moduleCli.GetApprovalUserTaskList(ctx, &lark.GetApprovalUserTaskListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalSearchApprovalCarbonCopy(func(ctx context.Context, request *lark.SearchApprovalCarbonCopyReq, options ...lark.MethodOptionFunc) (*lark.SearchApprovalCarbonCopyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSearchApprovalCarbonCopy()

			_, _, err := moduleCli.SearchApprovalCarbonCopy(ctx, &lark.SearchApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalCreateApprovalCarbonCopy(func(ctx context.Context, request *lark.CreateApprovalCarbonCopyReq, options ...lark.MethodOptionFunc) (*lark.CreateApprovalCarbonCopyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCreateApprovalCarbonCopy()

			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalPreviewApprovalInstance(func(ctx context.Context, request *lark.PreviewApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.PreviewApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalPreviewApprovalInstance()

			_, _, err := moduleCli.PreviewApprovalInstance(ctx, &lark.PreviewApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalUpdateApprovalMessage(func(ctx context.Context, request *lark.UpdateApprovalMessageReq, options ...lark.MethodOptionFunc) (*lark.UpdateApprovalMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalUpdateApprovalMessage()

			_, _, err := moduleCli.UpdateApprovalMessage(ctx, &lark.UpdateApprovalMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalSubscribeApprovalSubscription(func(ctx context.Context, request *lark.SubscribeApprovalSubscriptionReq, options ...lark.MethodOptionFunc) (*lark.SubscribeApprovalSubscriptionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSubscribeApprovalSubscription()

			_, _, err := moduleCli.SubscribeApprovalSubscription(ctx, &lark.SubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalUnsubscribeApprovalSubscription(func(ctx context.Context, request *lark.UnsubscribeApprovalSubscriptionReq, options ...lark.MethodOptionFunc) (*lark.UnsubscribeApprovalSubscriptionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalUnsubscribeApprovalSubscription()

			_, _, err := moduleCli.UnsubscribeApprovalSubscription(ctx, &lark.UnsubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalGetApprovalExternalList(func(ctx context.Context, request *lark.GetApprovalExternalListReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalExternalListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalExternalList()

			_, _, err := moduleCli.GetApprovalExternalList(ctx, &lark.GetApprovalExternalListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApprovalSendApprovalMessage(func(ctx context.Context, request *lark.SendApprovalMessageReq, options ...lark.MethodOptionFunc) (*lark.SendApprovalMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSendApprovalMessage()

			_, _, err := moduleCli.SendApprovalMessage(ctx, &lark.SendApprovalMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RollbackApprovalInstance(ctx, &lark.RollbackApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddApprovalInstanceSign(ctx, &lark.AddApprovalInstanceSignReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalTask(ctx, &lark.SearchApprovalTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalUserTaskList(ctx, &lark.GetApprovalUserTaskListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalCarbonCopy(ctx, &lark.SearchApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.PreviewApprovalInstance(ctx, &lark.PreviewApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApprovalMessage(ctx, &lark.UpdateApprovalMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SubscribeApprovalSubscription(ctx, &lark.SubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UnsubscribeApprovalSubscription(ctx, &lark.UnsubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalExternalList(ctx, &lark.GetApprovalExternalListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SendApprovalMessage(ctx, &lark.SendApprovalMessageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Approval
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RollbackApprovalInstance(ctx, &lark.RollbackApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.AddApprovalInstanceSign(ctx, &lark.AddApprovalInstanceSignReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalTask(ctx, &lark.SearchApprovalTaskReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalUserTaskList(ctx, &lark.GetApprovalUserTaskListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchApprovalCarbonCopy(ctx, &lark.SearchApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.PreviewApprovalInstance(ctx, &lark.PreviewApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApprovalMessage(ctx, &lark.UpdateApprovalMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SubscribeApprovalSubscription(ctx, &lark.SubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UnsubscribeApprovalSubscription(ctx, &lark.UnsubscribeApprovalSubscriptionReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApprovalExternalList(ctx, &lark.GetApprovalExternalListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SendApprovalMessage(ctx, &lark.SendApprovalMessageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
