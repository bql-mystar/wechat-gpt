// -----------------------------------------------------------------------
// Copyright (c) 2022 Akuvox Corporation and Akubela-Eevee Contributors.
// All rights reserved.
// -----------------------------------------------------------------------

package callback

import (
	"github.com/eatmoreapple/openwechat"
	"log"
)

// LoginCallback 登陆回调
type LoginCallback = func(openwechat.CheckLoginResponse)

// SuccessLoginCallback login successful
var SuccessLoginCallback LoginCallback = func(_ openwechat.CheckLoginResponse) {
	log.Println("login successful......")
}
