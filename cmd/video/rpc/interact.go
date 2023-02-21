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
		client.WithHostPorts("0.0.0.0:8890"),
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

func CommentAction(ctx context.Context, req *interact.CommentActionRequest) (*interact.Comment, error) {
	resp, err := interactClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Comment, nil
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

func GetVideoFavoriteCount(ctx context.Context, req *interact.CountVideoGetFavoriteRequest) (int64, error) {
	resp, err := interactClient.CountVideoGetFavorite(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Count, nil
}

func GetVideoCommentCount(ctx context.Context, req *interact.CountVideoGetCommentRequest) (int64, error) {
	resp, err := interactClient.CountVideoGetComment(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Count, nil
}

func GetUserFavoriteCount(ctx context.Context, req *interact.CountUserGetFavoriteRequest) (int64, error) {
	resp, err := interactClient.CountUserGetFavorite(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Count, nil
}
