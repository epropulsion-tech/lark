// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// QueryTenant 获取企业名称、企业编号等企业信息
//
// 如果ISV应用是预装的并且180天内企业未使用过此应用，则无法通过此接口获取到企业信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query
func (r *TenantAPI) QueryTenant(ctx context.Context, request *QueryTenantReq) (*QueryTenantResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/tenant/v2/tenant/query",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(queryTenantResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Tenant", "QueryTenant", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type QueryTenantReq struct{}

type queryTenantResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *QueryTenantResp `json:"data,omitempty"` //
}

type QueryTenantResp struct {
	Tenant *QueryTenantRespTenant `json:"tenant,omitempty"` // 企业信息
}

type QueryTenantRespTenant struct {
	Name      string                       `json:"name,omitempty"`       // 企业名称
	DisplayID string                       `json:"display_id,omitempty"` // 企业编号
	TenantTag int                          `json:"tenant_tag,omitempty"` // 个人版/团队版标志, 可选值有: `0`：团队版, `2`：个人版
	TenantKey string                       `json:"tenant_key,omitempty"` // 企业标识
	Avatar    *QueryTenantRespTenantAvatar `json:"avatar,omitempty"`     // 企业头像
}

type QueryTenantRespTenantAvatar struct {
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 企业头像
	Avatar72     string `json:"avatar_72,omitempty"`     // 企业头像 72x72
	Avatar240    string `json:"avatar_240,omitempty"`    // 企业头像 240x240
	Avatar640    string `json:"avatar_640,omitempty"`    // 企业头像 640x640
}