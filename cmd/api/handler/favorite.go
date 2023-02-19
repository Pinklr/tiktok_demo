package handler

import (
	"context"
	"strconv"

	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func FavoriteHandler(ctx context.Context, c *app.RequestContext) {
	token := c.Query(constants.TokenQueryKey)
	videoId, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query(constants.ActionTypeQueryKey), 10, 32)
	if len(token) == 0 || err1 != nil || err2 != nil {
		SendResponse(c, errno.ParamErr, nil, "data")
		return
	}

	println(token, videoId, actionType)
	//需要rpc实现
	// err := rpc.Favorite(ctx, &interact.FavoriteRequest{
	// 	UserID:     token,
	// 	VideoID:    videoId,
	// 	ActionType: actionType,
	// })
	// if err != nil {
	// 	SendResponse(c, err, nil, "data")
	// 	return
	// }

	SendResponse(c, errno.Success, nil, "data")

}
