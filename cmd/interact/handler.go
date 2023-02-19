package main

import (
	"context"
	interact "github.com/Pinklr/tiktok_demo/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Favorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Favorite(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
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
