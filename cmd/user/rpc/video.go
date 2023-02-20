package rpc

import (
	"context"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video/videoservice"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/kitex/client"
)

var videoClient videoservice.Client

func initVideoRPC() {
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithHostPorts("0.0.0.0:8889"),
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
}

func CountUserVideo(ctx context.Context, req *video.CountUserVideoRequest) (int64, error) {
	resp, err := videoClient.CountUserVideo(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Count, nil
}
