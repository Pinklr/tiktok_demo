package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"` //数据库中存储的是密码的md5
}

func (u *User) TableName() string {
	return "user"
}

func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func MGetUsers(ctx context.Context, userIds []int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id in ?", userIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserInfo(ctx context.Context, userId int64) (*User, error) {
	res := new(User)
	DB.WithContext(ctx).Where("id = ?", userId).First(res)
	return res, nil
}
