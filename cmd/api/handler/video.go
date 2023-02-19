package handler

import (
	"context"
	"fmt"
	"github.com/Pinklr/tiktok_demo/cmd/api/rpc"
	"github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strconv"
	"time"
)

func UploadVideoHandler(ctx context.Context, c *app.RequestContext) {
	fileHeader, err := c.FormFile("data")
	title := c.PostForm("title")
	// 获取用户id
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
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

	playurl := constants.PlayURLPrefix + "http://localhost:9902/static/video/"
	err = rpc.UploadVideo(ctx, &video.VideoActionRequest{Video: &video.Video{
		Author: &video.User{
			Id: userId,
		},
		PlayUrl:       playurl,
		CoverUrl:      playurl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "",
	}})
	if err != nil {
		SendResponse(c, err, nil, "data")
		return
	}
	SendResponse(c, errno.Success, nil, "data")
}
