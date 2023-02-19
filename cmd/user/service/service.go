package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/Pinklr/tiktok_demo/cmd/user/dal/db"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
)

func CreateUser(ctx context.Context, username, password string) error {
	// 查询用户名是否已被使用
	users, err := db.QueryUser(ctx, username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// 计算密码md5
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return err
	}
	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))
	//写入数据库
	return db.CreateUser(ctx, []*db.User{
		&db.User{
			Username: username,
			Password: passwordMd5,
		},
	})
}

func CheckUser(ctx context.Context, username, password string) (int64, error) {
	users, err := db.QueryUser(ctx, username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return 0, err
	}
	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))
	if passwordMd5 != users[0].Password {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(users[0].ID), nil
}

func GetUserInfo(ctx context.Context, userID int64) (*user.User, error) {
	model, err := db.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}
	// TODO
	var a, b int64 = 100, 100
	var avatar string = "https://img1.baidu.com/it/u=1946634850,2532424654&fm=253&fmt=auto&app=138&f=JPEG?w=581&h=500"
	return &user.User{
		Id:              int64(model.Model.ID),
		Name:            model.Username,
		FollowCount:     &a,
		FollowerCount:   &b,
		IsFollow:        false,
		Avatar:          &avatar,
		BackgroundImage: nil,
		Signature:       nil,
		TotalFavorited:  nil,
		WorkCount:       nil,
		FavoriteCount:   nil,
	}, nil
}
