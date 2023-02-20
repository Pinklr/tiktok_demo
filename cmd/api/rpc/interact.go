package rpc

import (
	"context"

	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact/interactservice"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/kitex/client"
)

var interactClient interactservice.Client

func initInteractRpc() {
	c, err := interactservice.NewClient(
		constants.InteractServiceName,
		client.WithHostPorts("0.0.0.0:8889"),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}

func FavoriteAction(ctx context.Context, req *interact.FavoriteRequest) error {
	resp, err := interactClient.Favorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetFavorateList(ctx context.Context, req *interact.FavoriteListRequest) ([]*interact.Video, error) {
	resp, err := interactClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Videos, nil
}

func CommentAction(ctx context.Context, req *interact.CommentActionRequest) error {
	resp, err := interactClient.CommentAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetCommentList(ctx context.Context, req *interact.CommentListRequest) ([]*interact.Comment, error) {
	resp, err := interactClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Comments, nil
}
