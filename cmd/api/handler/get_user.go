package handler

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"strconv"
)

func GetUserHandler(ctx context.Context, c *app.RequestContext) {
	// 获取用户id
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 如果用户id为0，则从jwt token中获取用户自己的id
	if userId == 0 {
		claims := jwt.ExtractClaims(ctx, c)
		userId = int64(claims[constants.IdentityKey].(float64))
	}

	user, err := rpc.GetUserInfo(ctx, &user.GetUserInfoRequest{UserId: userId})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	//c.JSON(consts.StatusOK, map[string]interface{}{
	//	"status_code": 0,
	//	"status_msg":  "success",
	//	"user": &map[string]interface{}{
	//		"id":             userId,
	//		"name":           info["name"],
	//		"follow_count":   info["follow_count"],
	//		"follower_count": info["follower_count"],
	//		"is_follow":      false,
	//	},
	//})
	SendResponse(c, errno.Success, user, "user")
}
