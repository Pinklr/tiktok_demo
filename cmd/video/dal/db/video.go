package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AuthorID int64  `json:"author_id"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
	Title    string `json:"title"`
}

func (v *Video) TableName() string {
	return "video"
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

func LatestVideo(ctx context.Context, latest time.Time) ([]*Video, error) {
	res := make([]*Video, 0)
	err := DB.WithContext(ctx).Where("created_at < ?", latest).Order("created_at desc").Limit(30).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
