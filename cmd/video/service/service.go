package service

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/video/dal/db"
)

func UploadVideo(ctx context.Context, authorID int64, playURL, coverURL, title string) error {
	return db.CreateVideo(ctx, []*db.Video{&db.Video{
		AuthorID: authorID,
		PlayURL:  playURL,
		CoverURL: coverURL,
		Title:    title,
	}})
}
