package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/Pinklr/tiktok_demo/cmd/user/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"io"

	"github.com/Pinklr/tiktok_demo/cmd/user/dal/db"
	"github.com/Pinklr/tiktok_demo/cmd/user/pack"
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
	user := pack.User(model)

	// TODO 获取用户作品数、点赞数、评论数
	count, err := rpc.CountUserVideo(ctx, &video.CountUserVideoRequest{UserID: userID})
	if err != nil {
		return nil, err
	}
	user.WorkCount = &count
	return user, nil
}

func MGetUser(ctx context.Context, userIDs []int64) ([]*user.User, error) {
	model, err := db.MGetUsers(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	users := pack.Users(model)

	// TODO 获取用户作品数、点赞数、评论数
	for i := 0; i < len(users); i += 1 {
		id := users[i].Id
		count, err := rpc.CountUserVideo(ctx, &video.CountUserVideoRequest{UserID: id})
		if err != nil {
			return nil, err
		}
		users[i].WorkCount = &count
	}
	return users, nil
}
