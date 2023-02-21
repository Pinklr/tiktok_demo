package db

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	VideoID int64  `json:"video_id"`
	Content string `json:"content"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func CreateComment(ctx context.Context, comments *Comment) error {
	err := DB.WithContext(ctx).Create(comments).Error
	return err
}

func DeleteComment(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).Delete(&Comment{}, id).Error
}

func CountVideoComment(ctx context.Context, videoID int64) (int64, error) {
	var res int64
	err := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoID).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}

func VideoCommentList(ctx context.Context, videoID int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	err := DB.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at desc").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
