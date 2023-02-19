package pack

import (
	"errors"
	"github.com/Pinklr/tiktok_demo/cmd/video/dal/db"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *video.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *video.BaseResp {
	return &video.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

func Video(m *db.Video) *video.Video {
	return &video.Video{
		Id: int64(m.ID),
		Author: &video.User{
			Id:              m.AuthorID,
			Name:            "",
			FollowCount:     nil,
			FollowerCount:   nil,
			IsFollow:        false,
			Avatar:          nil,
			BackgroundImage: nil,
			Signature:       nil,
			TotalFavorited:  nil,
			WorkCount:       nil,
			FavoriteCount:   nil,
		},
		PlayUrl:       m.PlayURL,
		CoverUrl:      m.CoverURL,
		FavoriteCount: 11,
		CommentCount:  11,
		IsFavorite:    false,
		Title:         m.Title,
	}
}

func Videos(model []*db.Video) []*video.Video {
	res := make([]*video.Video, 0, len(model))
	for _, item := range model {
		res = append(res, Video(item))
	}
	return res
}
