package handler

import (
	"context"
	"go-micro-v3/proto/hello"
	"go-micro-v3/proto/notice"
	"log"
)

type HelloService struct {
	NoticeServer notice.NoticeService
}

func (h HelloService) Call(ctx context.Context, request *hello.CallRequest, response *hello.CallResponse) error {
	log.Println("this is HelloService.Call")
	response.Msg = "Hello, " + request.Name
	return nil
}
