package handler

import "github.com/eatmoreapple/openwechat"

type MessageHandler interface {
	match() error
	handle() error
	ReplyText() error
}

func NewHandler() openwechat.MessageHandler {
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// todo 加入相应的消息处理
	dispatcher.RegisterHandler(userMatchFunc, userMessageContextHandler)

	return dispatcher.AsMessageHandler()
}
