package biz

import (
	"context"
	"github.com/google/uuid"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

// SendMessageHandler 发送消息
func SendMessageHandler(client *lark.Client, msg, openID string) error {
	uid := uuid.New()
	key := uid.String()

	// 接收用户输入的消息传入ChatGPT
	chatGPTMessage := ChatGPT(msg)

	// 富文本
	zhCnPostText := &larkim.MessagePostText{Text: chatGPTMessage, UnEscape: false}

	// 中文
	zhCn := larkim.NewMessagePostContent().
		ContentTitle(msg).
		AppendContent([]larkim.MessagePostElement{zhCnPostText}).
		Build()

	// 构建消息体
	content, err := larkim.NewMessagePost().ZhCn(zhCn).Build()
	if err != nil {
		logrus.WithError(err).Errorf("构建消息体失败")
		panic(err)
	}

	logrus.Infof("content: %v", content)

	// 发起请求
	resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeOpenId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(larkim.MsgTypePost).
			ReceiveId(openID).
			Content(content).
			Uuid(key).
			Build()).
		Build())

	// 处理错误
	if err != nil {
		logrus.WithError(err).Errorf("发送消息-处理错误")
		panic(err)
	}

	// 服务端错误处理
	if !resp.Success() {
		logrus.WithError(err).Errorf("发送消息-服务端处理错误")
		panic(resp.Msg)
	}

	return nil
}
