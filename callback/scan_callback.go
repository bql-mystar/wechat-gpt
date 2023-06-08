// -----------------------------------------------------------------------
// Copyright (c) 2022 Akuvox Corporation and Akubela-Eevee Contributors.
// All rights reserved.
// -----------------------------------------------------------------------

package callback

import (
	"github.com/eatmoreapple/openwechat"
	"log"
)

// ScanCallBack 扫码回调,可获取扫码用户的头像
type ScanCallBack func(openwechat.CheckLoginResponse)

// WelcomeScanCallBack welcome
var WelcomeScanCallBack ScanCallBack = func(_ openwechat.CheckLoginResponse) {
	log.Printf("welcome......")
}
