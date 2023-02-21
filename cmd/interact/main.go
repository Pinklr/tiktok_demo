package main

import (
	"github.com/Pinklr/tiktok_demo/cmd/interact/dal"
	interact "github.com/Pinklr/tiktok_demo/kitex_gen/interact/interactservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8890")
	if err != nil {
		panic(err)
	}
	svr := interact.NewServer(
		new(InteractServiceImpl),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
