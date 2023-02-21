package pack

import (
	"errors"
	"github.com/Pinklr/tiktok_demo/cmd/user/dal/db"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *user.BaseResp {
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

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

func User(model *db.User) *user.User {
	// TODO
	var a, b int64 = 100, 100
	var avatar string = "http://" + constants.ServerAddr + ":9002/static/image/avatar.jpeg"
	var background string = "http://" + "ServerAddr" + ":9002/static/image/background.jpeg"
	var signature string = "这个用户很懒，什么都没有留下"
	return &user.User{
		Id:              int64(model.Model.ID),
		Name:            model.Username,
		FollowCount:     &a,
		FollowerCount:   &b,
		IsFollow:        false,
		Avatar:          &avatar,
		BackgroundImage: &background,
		Signature:       &signature,
		TotalFavorited:  nil,
		WorkCount:       nil,
		FavoriteCount:   nil,
	}
}

func Users(model []*db.User) []*user.User {
	res := make([]*user.User, 0, len(model))
	for _, item := range model {
		res = append(res, User(item))
	}
	return res
}
