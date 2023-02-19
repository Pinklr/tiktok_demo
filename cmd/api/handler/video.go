package handler

import (
	"context"
	"fmt"
	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"log"
	"strconv"
	"time"
)

func UploadVideoHandler(ctx context.Context, c *app.RequestContext) {
	fileHeader, err := c.FormFile("data")
	title := c.PostForm("title")
	// 获取用户id
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))
	log.Println(title)
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	log.Println(fileHeader.Header)
	filename := fmt.Sprintf("%v%v", time.Now().Unix(), fileHeader.Filename)
	err = c.SaveUploadedFile(fileHeader, constants.VideoSaveDirectory+filename)
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}

	playurl := constants.PlayURLPrefix + filename
	err = rpc.UploadVideo(ctx, &video.VideoActionRequest{Video: &video.Video{
		Author: &video.User{
			Id: userId,
		},
		PlayUrl:       playurl,
		CoverUrl:      playurl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	SendResponse(c, errno.Success, nil, "data")
}

func FeedHandler(ctx context.Context, c *app.RequestContext) {
	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	var req video.FeedRequest
	req.LatestTime = latestTime

	videos, nextTime, err := rpc.Feed(ctx, &req)
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	//SendResponse(c, errno.Success, videos, "video_list")
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": errno.Success.ErrCode,
		"status_msg":  errno.Success.ErrMsg,
		"next_time":   nextTime,
		"video_list":  videos,
	})
}
