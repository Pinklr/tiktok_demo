package service

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/video/dal/db"
	"github.com/Pinklr/tiktok_demo/cmd/video/pack"
	"github.com/Pinklr/tiktok_demo/cmd/video/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
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
	timeLatest := time.Unix(latest_time/1000, 0)
	model, err := db.LatestVideo(ctx, timeLatest)
	if err != nil {
		return nil, 0, err
	}
	nextTime := latest_time
	if len(model) > 0 {
		nextTime = model[len(model)-1].CreatedAt.Unix() * 1000
	} else {
		// 没有视频，直接返回空列表
		return []*video.Video{}, nextTime, nil
	}

	videos := pack.Videos(model)

	//获取视频作者信息
	uidMap := make(map[int64]struct{})
	for _, item := range videos {
		uidMap[item.Author.Id] = struct{}{}
	}
	uids := make([]int64, 0)
	for i := range uidMap {
		uids = append(uids, i)
	}
	userMap, err := rpc.MGetUser(ctx, &user.MGetUserRequest{UserIds: uids})
	if err != nil {
		return nil, 0, err
	}

	for i := 0; i < len(videos); i++ {
		if u, ok := userMap[videos[i].Author.Id]; ok {
			videos[i].Author = &video.User{
				Id:              u.Id,
				Name:            u.Name,
				FollowCount:     u.FollowCount,
				FollowerCount:   u.FollowerCount,
				IsFollow:        u.IsFollow,
				Avatar:          u.Avatar,
				BackgroundImage: u.BackgroundImage,
				Signature:       u.Signature,
				TotalFavorited:  u.TotalFavorited,
				WorkCount:       u.WorkCount,
				FavoriteCount:   u.FavoriteCount,
			}
		}
	}

	return videos, nextTime, nil
}

func GetVideoByUserID(ctx context.Context, userID int64) ([]*video.Video, error) {
	model, err := db.GetVideoByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 没有视频，直接返回空列表
	if len(model) == 0 {
		return []*video.Video{}, nil
	}

	videos := pack.Videos(model)

	//获取视频作者信息
	uidMap := make(map[int64]struct{})
	for _, item := range videos {
		uidMap[item.Author.Id] = struct{}{}
	}
	uids := make([]int64, 0)
	for i := range uidMap {
		uids = append(uids, i)
	}
	userMap, err := rpc.MGetUser(ctx, &user.MGetUserRequest{UserIds: uids})
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(videos); i++ {
		if u, ok := userMap[videos[i].Author.Id]; ok {
			videos[i].Author = &video.User{
				Id:              u.Id,
				Name:            u.Name,
				FollowCount:     u.FollowCount,
				FollowerCount:   u.FollowerCount,
				IsFollow:        u.IsFollow,
				Avatar:          u.Avatar,
				BackgroundImage: u.BackgroundImage,
				Signature:       u.Signature,
				TotalFavorited:  u.TotalFavorited,
				WorkCount:       u.WorkCount,
				FavoriteCount:   u.FavoriteCount,
			}
		}
	}

	return videos, nil
}

func CountUserVideo(ctx context.Context, userID int64) (int64, error) {
	count, err := db.CountUserVideo(ctx, userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
