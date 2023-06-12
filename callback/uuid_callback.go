package callback

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"log"
	"runtime"
	"wechat-gpt/consts"
)

// UUIDCallback 获取UUID的回调函数
type UUIDCallback func(uuid string)

// GenerateQrcodeByOperationCallback Presentation according to different operating systems
var GenerateQrcodeByOperationCallback UUIDCallback = func(uuid string) {
	log.Println("login in " + runtime.GOOS)
	runEnv := runtime.GOOS
	if runEnv == consts.WindowsEnv {
		openwechat.PrintlnQrcodeUrl(uuid)
	} else if runEnv == consts.MacEnv || runEnv == consts.LinuxEnv {
		GenerateQrcodeCallback(uuid)
	} else {
		log.Fatalf("%s operation is not support\n", runEnv)
	}
}

// GenerateQrcodeCallback Generate CAPTCHA directly for display on the console
var GenerateQrcodeCallback UUIDCallback = func(uuid string) {
	q, _ := qrcode.New(consts.WechatLoginPrefixDomain+uuid, qrcode.High)
	fmt.Println(q.ToSmallString(true))
}
