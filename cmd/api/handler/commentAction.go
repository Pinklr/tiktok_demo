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

func CommentActionHandler(ctx context.Context, c *app.RequestContext) {

	//token := c.Query(constants.TokenQueryKey)
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	videoId, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query(constants.ActionTypeQueryKey), 10, 32)
	if err1 != nil {
		SendResponse(c, err1, nil, "data")
		return
	}
	if err2 != nil {
		SendResponse(c, err2, nil, "data")
		return
	}
	var commentText string
	var commentId int64
	var err error
	if actionType == 1 {
		commentText = c.Query(constants.CommentText)
		if len(commentText) == 0 {
			SendResponse(c, errno.ParamErr, nil, "data")
			return
		}
	} else if actionType == 2 {
		commentId, err = strconv.ParseInt(c.Query(constants.CommentId), 10, 64)
		if err != nil {
			SendResponse(c, err, nil, "data")
			return
		}
		if commentId <= 0 {
			SendResponse(c, errno.ParamErr, nil, "data")
			return
		}
	} else {
		SendResponse(c, errno.ParamErr, nil, "data")
		return
	}

	comment, err := rpc.CommentAction(ctx, &interact.CommentActionRequest{

		UserID:      userId,
		VideoID:     videoId,
		ActionType:  actionType,
		CommentText: &commentText,
		CommentID:   &commentId,
	})
	if err != nil {
		SendResponse(c, err, nil, "video_list")
		return
	}

	if actionType == 2 {
		SendResponse(c, errno.Success, nil, "data")
	} else if actionType == 1 {
		SendResponse(c, errno.Success, comment, "comment")
	}

}
