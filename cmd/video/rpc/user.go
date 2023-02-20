package rpc

import (
	"context"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user"
	"github.com/Pinklr/tiktok_demo/kitex_gen/user/userservice"
	"github.com/Pinklr/tiktok_demo/pkg/constants"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUserRPC() {
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) (map[int64]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*user.User)
	for _, u := range resp.Users {
		res[u.Id] = u
	}
	return res, nil
}
