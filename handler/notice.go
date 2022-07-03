package handler

import (
	"context"
	"go-micro-v3/proto/notice"
	"log"
)

type NoticeService struct {
}

func (n NoticeService) Send(ctx context.Context, request *notice.SendRequest, response *notice.SendResponse) error {
	log.Println("this is NoticeService.Send")
	response.Msg = "NoticeService 接收到请求啦"
	return nil
}
