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

func VideoAction(ctx context.Context, req *video.VideoActionRequest) error {
	resp, err := videoClient.VideoAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func List(ctx context.Context, req *video.ListRequest) ([]*video.Video, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Videos, nil
}

func Feed(ctx context.Context, req *video.FeedRequest) (int64, []*video.Video, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return 0, nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.NextTime, resp.Videos, nil
}
