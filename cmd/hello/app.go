package main

import (
	"github.com/asim/go-micro/v3"
	cons "go-micro-v3/constant"
	"go-micro-v3/handler"
	"go-micro-v3/proto/hello"
	"go-micro-v3/proto/notice"
	"log"
)

func main() {
	// 创建 micro 服务
	service := micro.NewService(
		// 设置 micro 服务名称
		micro.Name(cons.ServerMicroServer),
	)
	// 初始化 micro 服务
	service.Init()

	// 获取 micro-notice 服务的 noticeService，才能在 Call 中调用 notice send
	noticeService := notice.NewNoticeService(cons.NoticeMicroServer, service.Client())
	err := hello.RegisterHelloWorldHandler(service.Server(), handler.HelloService{NoticeServer: noticeService})
	if err != nil {
		log.Fatal("注册 server service 失败 -> ", err)
		return
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
