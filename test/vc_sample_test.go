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

func Test_VC_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.VC

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCMeeting(ctx, &lark.GetVCMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ListVCMeetingByNo(ctx, &lark.ListVCMeetingByNoReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.KickoutVCMeeting(ctx, &lark.KickoutVCMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCHostMeeting(ctx, &lark.SetVCHostMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCDailyReport(ctx, &lark.GetVCDailyReportReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCTopUserReport(ctx, &lark.GetVCTopUserReportReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCRoomConfig(ctx, &lark.GetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCRoomConfig(ctx, &lark.SetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.VC

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCGetVCMeeting(func(ctx context.Context, request *lark.GetVCMeetingReq, options ...lark.MethodOptionFunc) (*lark.GetVCMeetingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCGetVCMeeting()

			_, _, err := moduleCli.GetVCMeeting(ctx, &lark.GetVCMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCListVCMeetingByNo(func(ctx context.Context, request *lark.ListVCMeetingByNoReq, options ...lark.MethodOptionFunc) (*lark.ListVCMeetingByNoResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCListVCMeetingByNo()

			_, _, err := moduleCli.ListVCMeetingByNo(ctx, &lark.ListVCMeetingByNoReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCKickoutVCMeeting(func(ctx context.Context, request *lark.KickoutVCMeetingReq, options ...lark.MethodOptionFunc) (*lark.KickoutVCMeetingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCKickoutVCMeeting()

			_, _, err := moduleCli.KickoutVCMeeting(ctx, &lark.KickoutVCMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCSetVCHostMeeting(func(ctx context.Context, request *lark.SetVCHostMeetingReq, options ...lark.MethodOptionFunc) (*lark.SetVCHostMeetingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCSetVCHostMeeting()

			_, _, err := moduleCli.SetVCHostMeeting(ctx, &lark.SetVCHostMeetingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCGetVCDailyReport(func(ctx context.Context, request *lark.GetVCDailyReportReq, options ...lark.MethodOptionFunc) (*lark.GetVCDailyReportResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCGetVCDailyReport()

			_, _, err := moduleCli.GetVCDailyReport(ctx, &lark.GetVCDailyReportReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCGetVCTopUserReport(func(ctx context.Context, request *lark.GetVCTopUserReportReq, options ...lark.MethodOptionFunc) (*lark.GetVCTopUserReportResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCGetVCTopUserReport()

			_, _, err := moduleCli.GetVCTopUserReport(ctx, &lark.GetVCTopUserReportReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCGetVCRoomConfig(func(ctx context.Context, request *lark.GetVCRoomConfigReq, options ...lark.MethodOptionFunc) (*lark.GetVCRoomConfigResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCGetVCRoomConfig()

			_, _, err := moduleCli.GetVCRoomConfig(ctx, &lark.GetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockVCSetVCRoomConfig(func(ctx context.Context, request *lark.SetVCRoomConfigReq, options ...lark.MethodOptionFunc) (*lark.SetVCRoomConfigResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockVCSetVCRoomConfig()

			_, _, err := moduleCli.SetVCRoomConfig(ctx, &lark.SetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.VC

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCMeeting(ctx, &lark.GetVCMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ListVCMeetingByNo(ctx, &lark.ListVCMeetingByNoReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.KickoutVCMeeting(ctx, &lark.KickoutVCMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCHostMeeting(ctx, &lark.SetVCHostMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCDailyReport(ctx, &lark.GetVCDailyReportReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCTopUserReport(ctx, &lark.GetVCTopUserReportReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCRoomConfig(ctx, &lark.GetVCRoomConfigReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCRoomConfig(ctx, &lark.SetVCRoomConfigReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.VC
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCMeeting(ctx, &lark.GetVCMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ListVCMeetingByNo(ctx, &lark.ListVCMeetingByNoReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.KickoutVCMeeting(ctx, &lark.KickoutVCMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCHostMeeting(ctx, &lark.SetVCHostMeetingReq{
				MeetingID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCDailyReport(ctx, &lark.GetVCDailyReportReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCTopUserReport(ctx, &lark.GetVCTopUserReportReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetVCRoomConfig(ctx, &lark.GetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SetVCRoomConfig(ctx, &lark.SetVCRoomConfigReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
