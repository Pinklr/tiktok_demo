package main

import (
	"context"
	interact "github.com/Pinklr/tiktok_demo/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Favorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Favorite(ctx context.Context, req *interact.FavoriteRequest) (resp *interact.FavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// CountVideoGetFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountVideoGetFavorite(ctx context.Context, req *interact.CountVideoGetFavoriteRequest) (resp *interact.CountResponse, err error) {
	// TODO: Your code here...
	return
}

// CountVideoGetComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountVideoGetComment(ctx context.Context, req *interact.CountVideoGetCommentRequest) (resp *interact.CountResponse, err error) {
	// TODO: Your code here...
	return
}

// CountUserGetFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountUserGetFavorite(ctx context.Context, req *interact.CountUserGetFavoriteRequest) (resp *interact.CountResponse, err error) {
	// TODO: Your code here...
	return
}

// CountUserFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CountUserFavorite(ctx context.Context, req *interact.CountUserFavoriteRequest) (resp *interact.CountUserFavoriteRequest, err error) {
	// TODO: Your code here...
	return
}
