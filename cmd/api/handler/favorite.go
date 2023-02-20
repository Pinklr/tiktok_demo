package handler

import (
	"context"
	"strconv"

	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

func FavoriteHandler(ctx context.Context, c *app.RequestContext) {

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	videoId, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query(constants.ActionTypeQueryKey), 10, 32)
	if err1 != nil || err2 != nil {
		SendResponse(c, errno.ParamErr, nil, "data")
		return
	}

	//需要rpc实现
	err := rpc.FavoriteAction(ctx, &interact.FavoriteRequest{
		UserID:     userId,
		VideoID:    videoId,
		ActionType: actionType,
	})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}

	SendResponse(c, errno.Success, nil, "data")

}
