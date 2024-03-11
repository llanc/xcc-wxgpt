package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"strings"
)

func PrintlnQrcodeUrlNotOpen(uuid string) {
	println("访问下面网址扫描二维码登录")
	qrcodeUrl := openwechat.GetQrcodeUrl(uuid)
	println(qrcodeUrl)
}

func main() {
	nickname := strings.TrimSpace(os.Getenv("BOT_NICKNAME"))
	if len(nickname) == 0 {
		panic("BOT_NICKNAME IS BLANK")
		//nickname = "xcc"
	}
	botKey := "@" + nickname + " "

	// 桌面模式
	bot := openwechat.DefaultBot(openwechat.Desktop)

	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		msgContent := msg.Content
		fmt.Println("Text: ", msgContent)
		if !strings.HasPrefix(msgContent, botKey) && !strings.HasSuffix(msgContent, botKey) {
			return
		}
		gptMsg := ""
		if strings.HasPrefix(msgContent, botKey) {
			gptMsg, _ = strings.CutPrefix(msgContent, botKey)
		} else if strings.HasSuffix(msgContent, botKey) {
			gptMsg, _ = strings.CutSuffix(msgContent, botKey)
		}
		if len(gptMsg) == 0 {
			msg.ReplyText("要@我并且问问题才行哦")
		}
		//msg.ReplyText(gptMsg)
		msg.ReplyText(chat(gptMsg))
	})

	// 注册消息处理函数
	bot.MessageHandler = dispatcher.AsMessageHandler()

	bot.UUIDCallback = PrintlnQrcodeUrlNotOpen

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 热登录
	if err := bot.HotLogin(reloadStorage); err != nil {
		fmt.Println("热登录失败：", err)
		err = bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
		if err != nil {
			fmt.Println("登录失败", err)
			return
		}
	}
	//
	//self, err := bot.GetCurrentUser()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// 获取所有的好友
	//friends, err := self.Friends()
	//fmt.Println(friends, err)
	//
	//// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	err := bot.Block()
	defer fmt.Println(err)
}
