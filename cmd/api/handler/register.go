package handler

import (
	"context"
	"fmt"
	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	username := c.Query(constants.UsernameQueryKey)
	password := c.Query(constants.PasswordQueryKey)
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, errno.ParamErr, nil, "data")
		return
	}
	err := rpc.CreateUser(ctx, &user.CreateUserRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}

	// 注册成功后自动登录
	uri := fmt.Sprintf("/douyin/user/login/?username=%v&password=%v", username, password)
	c.Redirect(consts.StatusTemporaryRedirect, []byte(uri))

}
