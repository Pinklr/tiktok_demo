package pack

import (
	"github.com/Pinklr/tiktok_demo/cmd/interact/dal/db"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
)

func Comment(model *db.Comment) *interact.Comment {
	comments := &interact.Comment{
		Id:          int64(model.ID),
		User:        &interact.User{Id: model.UserID},
		Content:     model.Content,
		CreatedData: model.CreatedAt.String(),
	}

	// TODO 获取用户信息

	return comments
}

func Comments(model []*db.Comment) []*interact.Comment {
	res := make([]*interact.Comment, 0, len(model))
	for _, item := range model {
		res = append(res, Comment(item))
	}
	return res
}
