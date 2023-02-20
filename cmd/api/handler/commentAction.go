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
	video_id, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)
	action_type, err2 := strconv.ParseInt(c.Query(constants.ActionTypeQueryKey), 10, 32)
	comment_text := c.Query(constants.CommentText)
	comment_id, err3 := strconv.ParseInt(c.Query(constants.CommentId), 10, 64)

	//println(video_id, action_type, comment_id)
	if len(comment_text) == 0 || err1 != nil || err2 != nil || len(comment_text) == 0 || err3 != nil { //
		SendResponse(c, errno.ParamErr, nil, "data")
		return

	}
	//println(video_id, action_type, comment_id)
	err := rpc.CommentAction(ctx, &interact.CommentActionRequest{

		UserID:      userId,
		VideoID:     video_id,
		ActionType:  action_type,
		CommentText: &comment_text,
		CommentID:   &comment_id,
	})
	if err != nil {
		SendResponse(c, err, nil, "video_list")
		return
	}

	SendResponse(c, errno.Success, nil, "data")
}
