package pack

import (
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *interact.BaseResp {
	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *interact.BaseResp {
	return &interact.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}
