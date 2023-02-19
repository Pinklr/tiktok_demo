package service

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/video/dal/db"
	"github.com/Pinklr/tiktok_demo/cmd/video/pack"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"log"
	"time"
)

func UploadVideo(ctx context.Context, authorID int64, playURL, coverURL, title string) error {
	return db.CreateVideo(ctx, []*db.Video{&db.Video{
		AuthorID: authorID,
		PlayURL:  playURL,
		CoverURL: coverURL,
		Title:    title,
	}})
}

func Feed(ctx context.Context, latest_time int64) ([]*video.Video, int64, error) {
	timeLatest := time.Unix(latest_time, 0)
	log.Println(timeLatest)
	model, err := db.LatestVideo(ctx, timeLatest)
	if err != nil {
		return nil, 0, err
	}
	nextTime := latest_time
	if len(model) > 0 {
		nextTime = model[0].CreatedAt.Unix()
	}
	return pack.Videos(model), nextTime, nil
}
