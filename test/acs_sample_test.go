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

func Test_ACS_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.ACS

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordPhoto(ctx, &lark.GetACSAccessRecordPhotoReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordList(ctx, &lark.GetACSAccessRecordListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSDeviceList(ctx, &lark.GetACSDeviceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserFace(ctx, &lark.GetACSUserFaceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUserFace(ctx, &lark.UpdateACSUserFaceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUser(ctx, &lark.GetACSUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUser(ctx, &lark.UpdateACSUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserList(ctx, &lark.GetACSUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.ACS

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSAccessRecordPhoto(func(ctx context.Context, request *lark.GetACSAccessRecordPhotoReq, options ...lark.MethodOptionFunc) (*lark.GetACSAccessRecordPhotoResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSAccessRecordPhoto()

			_, _, err := moduleCli.GetACSAccessRecordPhoto(ctx, &lark.GetACSAccessRecordPhotoReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSAccessRecordList(func(ctx context.Context, request *lark.GetACSAccessRecordListReq, options ...lark.MethodOptionFunc) (*lark.GetACSAccessRecordListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSAccessRecordList()

			_, _, err := moduleCli.GetACSAccessRecordList(ctx, &lark.GetACSAccessRecordListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSDeviceList(func(ctx context.Context, request *lark.GetACSDeviceListReq, options ...lark.MethodOptionFunc) (*lark.GetACSDeviceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSDeviceList()

			_, _, err := moduleCli.GetACSDeviceList(ctx, &lark.GetACSDeviceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSUserFace(func(ctx context.Context, request *lark.GetACSUserFaceReq, options ...lark.MethodOptionFunc) (*lark.GetACSUserFaceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSUserFace()

			_, _, err := moduleCli.GetACSUserFace(ctx, &lark.GetACSUserFaceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSUpdateACSUserFace(func(ctx context.Context, request *lark.UpdateACSUserFaceReq, options ...lark.MethodOptionFunc) (*lark.UpdateACSUserFaceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSUpdateACSUserFace()

			_, _, err := moduleCli.UpdateACSUserFace(ctx, &lark.UpdateACSUserFaceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSUser(func(ctx context.Context, request *lark.GetACSUserReq, options ...lark.MethodOptionFunc) (*lark.GetACSUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSUser()

			_, _, err := moduleCli.GetACSUser(ctx, &lark.GetACSUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSUpdateACSUser(func(ctx context.Context, request *lark.UpdateACSUserReq, options ...lark.MethodOptionFunc) (*lark.UpdateACSUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSUpdateACSUser()

			_, _, err := moduleCli.UpdateACSUser(ctx, &lark.UpdateACSUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockACSGetACSUserList(func(ctx context.Context, request *lark.GetACSUserListReq, options ...lark.MethodOptionFunc) (*lark.GetACSUserListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockACSGetACSUserList()

			_, _, err := moduleCli.GetACSUserList(ctx, &lark.GetACSUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.ACS

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordPhoto(ctx, &lark.GetACSAccessRecordPhotoReq{
				AccessRecordID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordList(ctx, &lark.GetACSAccessRecordListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSDeviceList(ctx, &lark.GetACSDeviceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserFace(ctx, &lark.GetACSUserFaceReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUserFace(ctx, &lark.UpdateACSUserFaceReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUser(ctx, &lark.GetACSUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUser(ctx, &lark.UpdateACSUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserList(ctx, &lark.GetACSUserListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.ACS
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordPhoto(ctx, &lark.GetACSAccessRecordPhotoReq{
				AccessRecordID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSAccessRecordList(ctx, &lark.GetACSAccessRecordListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSDeviceList(ctx, &lark.GetACSDeviceListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserFace(ctx, &lark.GetACSUserFaceReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUserFace(ctx, &lark.UpdateACSUserFaceReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUser(ctx, &lark.GetACSUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateACSUser(ctx, &lark.UpdateACSUserReq{
				UserID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetACSUserList(ctx, &lark.GetACSUserListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
