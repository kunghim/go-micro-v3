package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	cons "go-micro-v3/constant"
	"go-micro-v3/proto/hello"
	"log"
)

func main() {
	// 创建 micro 服务
	service := micro.NewService(
		micro.Address(cons.ClientMicroAddress),
		// 设置 micro 服务名称
		micro.Name(cons.ClientMicroServer),
	)
	// 初始化 micro 服务
	service.Init()

	helloWorldService := hello.NewHelloWorldService(cons.ServerMicroServer, service.Client())
	callResponse, err := helloWorldService.Call(context.Background(), &hello.CallRequest{Name: "张三"})
	if err != nil {
		log.Fatal("调用 hello 服务的 send 接口失败 -> ", err)
		return
	}
	log.Println("执行成功 -> ", callResponse.Msg)
}
