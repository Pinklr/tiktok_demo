package main

import (
	"github.com/Pinklr/tiktok_demo/cmd/video/dal"
	"github.com/Pinklr/tiktok_demo/cmd/video/rpc"
	video "github.com/Pinklr/tiktok_demo/kitex_gen/video/videoservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func Init() {
	rpc.Init()
	dal.Init()
}

func main() {
	Init()
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8889")
	if err != nil {
		panic(err)
	}
	svr := video.NewServer(
		new(VideoServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
