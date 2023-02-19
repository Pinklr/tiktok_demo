package main

import (
	"context"
	video "github.com/Pinklr/tiktok_demo/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoAction(ctx context.Context, req *video.VideoActionRequest) (resp *video.VideoActionResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, req *video.ListRequest) (resp *video.ListResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.MGetVideoRequest) (resp *video.MGetVideoResponse, err error) {
	// TODO: Your code here...
	return
}
