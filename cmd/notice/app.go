package main

import (
	"github.com/asim/go-micro/v3"
	cons "go-micro-v3/constant"
	"go-micro-v3/handler"
	"go-micro-v3/proto/notice"
	"log"
)

func main() {
	// 创建 micro 服务
	service := micro.NewService(
		micro.Address(cons.NoticeMicroAddress),
		// 设置 micro 服务名称
		micro.Name(cons.NoticeMicroServer),
	)

	// 初始化 micro 服务
	service.Init()

	err := notice.RegisterNoticeServiceHandler(service.Server(), new(handler.NoticeService))
	if err != nil {
		log.Fatal("注册 notice service 失败 -> ", err)
		return
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
