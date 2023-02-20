package main

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/interact/pack"
	"github.com/Pinklr/tiktok_demo/cmd/interact/service"
	interact "github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Favorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Favorite(ctx context.Context, req *interact.FavoriteRequest) (resp *interact.FavoriteResponse, err error) {
	resp = new(interact.FavoriteResponse)
	if req.UserID <= 0 || req.VideoID <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	err = service.Favorite(ctx, req.UserID, req.VideoID, req.ActionType)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)
	if req.UserID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	videos, err := service.FavoriteList(ctx, req.UserID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos
	return
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	resp = new(interact.CommentActionResponse)
	if req.UserID <= 0 || req.VideoID <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	if req.ActionType == 1 && req.CommentText == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	if req.ActionType == 2 && req.CommentID == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	if req.ActionType == 1 {
		err = service.CreateComment(ctx, req.UserID, req.VideoID, *req.CommentText)
	} else {
		err = service.DeleteComment(ctx, *req.CommentID)
	}
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	resp = new(interact.CommentListResponse)
	if req.VideoID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	comments, err := service.CommentList(ctx, req.VideoID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comments = comments
	return
}

// CountVideoGetFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountVideoGetFavorite(ctx context.Context, req *interact.CountVideoGetFavoriteRequest) (resp *interact.CountResponse, err error) {
	resp = new(interact.CountResponse)
	if req.VideoID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	count, err := service.CountVideoGetFavorite(ctx, req.VideoID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Count = count
	return
}

// CountVideoGetComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountVideoGetComment(ctx context.Context, req *interact.CountVideoGetCommentRequest) (resp *interact.CountResponse, err error) {
	resp = new(interact.CountResponse)
	if req.VideoID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	count, err := service.CountVideoGetComment(ctx, req.VideoID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Count = count
	return
}

// CountUserGetFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountUserGetFavorite(ctx context.Context, req *interact.CountUserGetFavoriteRequest) (resp *interact.CountResponse, err error) {
	resp = new(interact.CountResponse)
	if req.UserID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	count, err := service.CountUserGetFavorite(ctx, req.UserID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Count = count
	return
}

// CountUserFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountUserFavorite(ctx context.Context, req *interact.CountUserFavoriteRequest) (resp *interact.CountResponse, err error) {
	resp = new(interact.CountResponse)
	if req.UserID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	count, err := service.CountUserFavorite(ctx, req.UserID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Count = count
	return
}
