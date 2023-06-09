// -----------------------------------------------------------------------
// Copyright (c) 2022 Akuvox Corporation and Akubela-Eevee Contributors.
// All rights reserved.
// -----------------------------------------------------------------------

package handler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"wechat-gpt/gpt"
)

var userMatchFunc openwechat.MatchFunc = func(message *openwechat.Message) bool {
	return message.IsSendByFriend() || message.IsSendBySelf()
}

var userMessageContextHandler openwechat.MessageContextHandler = func(ctx *openwechat.MessageContext) {
	user, _ := ctx.Message.Sender()
	log.Println(user)
	content := ctx.Message.Content
	reply, err := gpt.Gpt(content)
	if err != nil {
		log.Printf("request gpt response err, error is: %s\n", err.Error())
	}
	_, err = ctx.ReplyText("机器人回复：" + reply)
	if err != nil {
		fmt.Println(err.Error())
	}

}
