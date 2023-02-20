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

func UploadVideo(ctx context.Context, req *video.VideoActionRequest) error {
	resp, err := videoClient.VideoAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func Feed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, int64, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Videos, resp.NextTime, nil
}

func PublishList(ctx context.Context, req *video.ListRequest) ([]*video.Video, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Videos, nil
}
