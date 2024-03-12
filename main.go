package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"strings"
)

func printlnQrcodeUrlNotOpen(uuid string) {
	println("访问下面网址扫描二维码登录")
	qrcodeUrl := openwechat.GetQrcodeUrl(uuid)
	println(qrcodeUrl)
}
func groupMsgToBotMatch(msgContent string, botKey string) bool {
	fmt.Println("Text: ", msgContent)
	return strings.HasPrefix(msgContent, botKey) || strings.HasSuffix(msgContent, botKey)
}
func getGroupMsgToBot(msgContent string, botKey string) string {
	gptMsg := ""
	if strings.HasPrefix(msgContent, botKey) {
		gptMsg, _ = strings.CutPrefix(msgContent, botKey)
	} else if strings.HasSuffix(msgContent, botKey) {
		gptMsg, _ = strings.CutSuffix(msgContent, botKey)
	}
	return gptMsg
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
		if msg.IsSendByFriend() {
			msg.ReplyText(chat(msg.Content))
			return
		}
		msgContent := msg.Content
		if !groupMsgToBotMatch(msgContent, botKey) {
			return
		}
		groupMsgToBot := getGroupMsgToBot(msgContent, botKey)
		if len(groupMsgToBot) == 0 {
			msg.ReplyText("要@我并且问问题才行哦")
			return
		}
		//msg.ReplyText(groupMsgToBot)
		msg.ReplyText(chat(groupMsgToBot))
	})

	// 注册消息处理函数
	bot.MessageHandler = dispatcher.AsMessageHandler()

	bot.UUIDCallback = printlnQrcodeUrlNotOpen

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
