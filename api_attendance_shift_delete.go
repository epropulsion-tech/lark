// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteShift
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete
func (r *AttendanceAPI) DeleteShift(ctx context.Context, request *DeleteShiftReq) (*DeleteShiftResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "DeleteShift", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteShiftReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID，示例值："6919358778597097404"
}

type deleteShiftResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteShiftResp `json:"data,omitempty"`
}

type DeleteShiftResp struct{}