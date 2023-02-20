package service

import (
	"context"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
)

// Favorite 点赞, 取消赞
func Favorite(ctx context.Context, userID, videoID, Type int64) error {

}

// FavoriteList 返回用户点赞过的视频列表
func FavoriteList(ctx context.Context, userID int64) ([]*interact.Video, error) {

}

func CreateComment(ctx context.Context, userID, videoID int64, content string) error {

}

func DeleteComment(ctx context.Context, commentID int64) error {

}

func CommentList(ctx context.Context, videoID int64) ([]*interact.Comment, error) {

}

func CountVideoGetFavorite(ctx context.Context, videoID int64) (int64, error) {

}

func CountVideoGetComment(ctx context.Context, videoID int64) (int64, error) {

}

func CountUserGetFavorite(ctx context.Context, userID int64) (int64, error) {

}

func CountUserFavorite(ctx context.Context, userID int64) (int64, error) {

}
