package handler

import (
	"context"
	"strconv"

	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func FavoriteListHandler(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.ParseInt(c.Query(constants.UserIdQueryKey), 10, 64)
	token := c.Query(constants.TokenQueryKey)
	println(userId)
	if len(token) == 0 { //|| err1 != nil
		SendResponse(c, errno.ParamErr, nil, "data111")
		return
	}
	println(token, userId)
	/*video_list, err := rpc.FavoriteList(ctx, &interact.FavoriteListRequest{
		UserID: userId,
	})
	if err != nil {
		SendResponse(c, err, nil, "video_list")
		return
	}*/

	SendResponse(c, errno.Success, nil, "video_list")

}
