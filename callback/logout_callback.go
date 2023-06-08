// -----------------------------------------------------------------------
// Copyright (c) 2022 Akuvox Corporation and Akubela-Eevee Contributors.
// All rights reserved.
// -----------------------------------------------------------------------

package callback

import (
	"github.com/eatmoreapple/openwechat"
	"log"
)

// LogoutCallBack 退出回调
type LogoutCallBack func(*openwechat.Bot)

// StopBlockLogoutCallBack context done to stop block
var StopBlockLogoutCallBack = func(bot *openwechat.Bot) {
	bot.Context().Done()
	log.Printf("logout successful......")
}
