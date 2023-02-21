package main

import (
	"github.com/Pinklr/tiktok_demo/cmd/user/dal"
	"github.com/Pinklr/tiktok_demo/cmd/user/rpc"
	user "github.com/Pinklr/tiktok_demo/kitex_gen/user/userservice"
	"log"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
