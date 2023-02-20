package handler

import (
	"context"
	"strconv"

	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func CommentListHandler(ctx context.Context, c *app.RequestContext) {

	token := c.Query(constants.TokenQueryKey)
	video_id, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)

	println(video_id)
	if len(token) == 0 || err1 != nil {
		SendResponse(c, errno.ParamErr, nil, "data")
		return

	}

	Comment_list, err := rpc.GetCommentList(ctx, &interact.CommentListRequest{

			//UserID:      token,
			VideoID:     video_id,

		})
		if err != nil {
			SendResponse(c, err, Comment_list, "video_list")
			return
		}

	SendResponse(c, errno.Success, nil, "data")
}
