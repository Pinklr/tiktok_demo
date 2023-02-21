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

func InitVideoRPC() {
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithHostPorts("0.0.0.0:8889"),
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
}

func MGetVideo(ctx context.Context, req *video.MGetVideoRequest) ([]*video.Video, error) {
	resp, err := videoClient.MGetVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Videos, nil
}
