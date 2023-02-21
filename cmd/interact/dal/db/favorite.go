package db

import (
	"context"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  int64 `json:"user_id"`
	VideoID int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}

func CreateFavorite(ctx context.Context, favorite []*Favorite) error {
	return DB.WithContext(ctx).Create(favorite).Error
}

func DeleteFavorite(ctx context.Context, userID, videoID int64) error {
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&Favorite{}).Error
}

func CountUserFavorite(ctx context.Context, userID int64) (int64, error) {
	var res int64
	err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ?", userID).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}

func CountVideoFavorite(ctx context.Context, videoIDs []int64) (int64, error) {
	var res int64
	err := DB.WithContext(ctx).Model(&Favorite{}).Where("video_id in ?", videoIDs).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}

func UserFavoriteList(ctx context.Context, userID int64) ([]int64, error) {
	model := make([]*Favorite, 0)
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&model).Error
	if err != nil {
		return nil, err
	}
	res := make([]int64, 0, len(model))
	for _, item := range model {
		res = append(res, item.VideoID)
	}
	return res, err
}

func IsFavorited(ctx context.Context, userID, videoID int64) (bool, error) {
	model := make([]*Favorite, 0)
	err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userID, videoID).Find(&model).Error
	if err != nil {
		return false, err
	}
	return len(model) > 0, nil
}
