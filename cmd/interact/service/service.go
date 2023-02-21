package service

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/interact/dal/db"
	"github.com/Pinklr/tiktok_demo/cmd/interact/pack"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
)

// Favorite 点赞, 取消赞
func Favorite(ctx context.Context, userID, videoID, Type int64) error {
	if Type == 1 {
		return db.CreateFavorite(ctx, []*db.Favorite{&db.Favorite{
			UserID:  userID,
			VideoID: videoID,
		}})
	} else {
		return db.DeleteFavorite(ctx, userID, videoID)
	}
}

// FavoriteList 返回用户点赞过的视频列表
func FavoriteList(ctx context.Context, userID int64) ([]*interact.Video, error) {
	return nil, nil
}

func CreateComment(ctx context.Context, userID, videoID int64, content string) (*interact.Comment, error) {
	model := &db.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	err := db.CreateComment(ctx, model)
	if err != nil {
		return nil, err
	}
	return &interact.Comment{
		Id:          int64(model.ID),
		User:        &interact.User{Id: model.UserID},
		Content:     model.Content,
		CreatedData: model.CreatedAt.String(),
	}, nil
}

func DeleteComment(ctx context.Context, commentID int64) error {
	return db.DeleteComment(ctx, commentID)
}

func CommentList(ctx context.Context, videoID int64) ([]*interact.Comment, error) {
	model, err := db.VideoCommentList(ctx, videoID)
	if err != nil {
		return nil, err
	}
	if len(model) == 0 {
		return []*interact.Comment{}, err
	}
	return pack.Comments(model), nil
}

func CountVideoGetFavorite(ctx context.Context, videoID int64) (int64, error) {
	res, err := db.CountVideoFavorite(ctx, []int64{videoID})
	if err != nil {
		return 0, err
	}
	return res, nil
}

func CountVideoGetComment(ctx context.Context, videoID int64) (int64, error) {
	res, err := db.CountVideoComment(ctx, videoID)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func CountUserGetFavorite(ctx context.Context, userID int64) (int64, error) {
	return 0, nil
}

func CountUserFavorite(ctx context.Context, userID int64) (int64, error) {
	res, err := db.CountUserFavorite(ctx, userID)
	if err != nil {
		return 0, err
	}
	return res, nil
}
