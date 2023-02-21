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
	videoId, err1 := strconv.ParseInt(c.Query(constants.VodeoIdQueryKey), 10, 64)

	if len(token) == 0 || err1 != nil {
		SendResponse(c, errno.ParamErr, nil, "data")
		return

	}

	CommentList, err := rpc.GetCommentList(ctx, &interact.CommentListRequest{

		//UserID:      token,
		VideoID: videoId,
	})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}

	SendResponse(c, errno.Success, CommentList, "comment_list")
}
