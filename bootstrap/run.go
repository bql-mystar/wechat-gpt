package bootstrap

import (
	"context"
	"github.com/eatmoreapple/openwechat"
	"wechat-gpt/callback"
	"wechat-gpt/consts"
	"wechat-gpt/handler"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 用于阻塞管理
	bot := openwechat.DefaultBot(openwechat.WithContextOption(ctx), openwechat.Desktop)

	bot.UUIDCallback = callback.GenerateQrcodeByOperationCallback

	bot.ScanCallBack = callback.WelcomeScanCallBack

	bot.LoginCallBack = callback.SuccessLoginCallback

	bot.LogoutCallBack = callback.StopBlockLogoutCallBack

	bot.MessageHandler = handler.NewHandler()

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage(consts.HotReloadStorageFilePath + consts.HotReloadStorageFilename)

	defer reloadStorage.Close()

	// 执行热登录
	//if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		panic(err.Error())
	}

	bot.Block()
}
