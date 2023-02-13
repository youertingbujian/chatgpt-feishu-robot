package biz

import (
	"chatgpt/config"
	"encoding/json"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
	"strconv"
)

// ReceiveMessageEventHandle ...
func ReceiveMessageEventHandle(event *larkim.P2MessageReceiveV1) error {
	// 处理消息 event
	text := larkcore.Prettify(event.Event.Message.Content)
	openID := larkcore.Prettify(event.Event.Sender.SenderId.OpenId)
	openID = openID[1 : len(openID)-1]

	// 反序列化用户发送的内容
	var ct ContentText
	s, _ := strconv.Unquote(text)
	err := json.Unmarshal([]byte(s), &ct)
	if err != nil {
		logrus.WithError(err).Errorf("unmarshal new text")
		return err
	}

	// 创建 Client
	client := lark.NewClient(config.Conf.AppID, config.Conf.AppSecret)

	// 发送消息
	err = SendMessageHandler(client, ct.Text, openID)
	if err != nil {
		logrus.WithError(err).Errorf("failed to send msg")
		return err
	}

	return nil
}
