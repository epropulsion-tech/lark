// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Admin_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Admin()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAdminDeptStats(ctx, &lark.GetAdminDeptStatsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAdminUserStats(ctx, &lark.GetAdminUserStatsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Admin()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAdminDeptStats(ctx, &lark.GetAdminDeptStatsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAdminUserStats(ctx, &lark.GetAdminUserStatsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}