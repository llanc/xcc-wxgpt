package actor

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"io"
	"log"
	"os"
	"strings"
	"xcc-wxgpt/chat"
)

type Actor struct {
	nickname        string
	workKeyInGroup  string
	dataPath        string
	weChatBot       *openwechat.Bot
	hotLoginStorage io.ReadWriteCloser
	chat            *chat.Chat
}

func (actor *Actor) Online() {
	actor.definition()
	actor.ready()
	actor.doJob()
}

func (actor *Actor) definition() {
	nickname := strings.TrimSpace(os.Getenv("BOT_NICKNAME"))
	if len(nickname) == 0 {
		panic("BOT_NICKNAME IS BLANK")
	}
	dataPath := strings.TrimSpace(os.Getenv("DATA_PATH"))
	if len(dataPath) == 0 {
		panic("DATA_PATH IS BLANK")
	}
	if !strings.HasSuffix(dataPath, "/") {
		dataPath += "/"
	}
	actor.nickname = nickname
	actor.workKeyInGroup = "@" + nickname + " "
	actor.dataPath = dataPath
}

func (actor *Actor) ready() {
	actor.chat = &chat.Chat{}
	actor.chat.Init()
	actor.weChatBot = openwechat.DefaultBot(openwechat.Desktop)
	workbenchHandlerRegister(actor.weChatBot, actor.makeMessageHandler())
	actor.hotLoginStorage = makeStorage(actor.dataPath + "hotLogin.json")
}

func (actor *Actor) doJob() {
	if err := actor.weChatBot.HotLogin(actor.hotLoginStorage); err != nil {
		log.Println("热登录失败：", err)
		err = actor.weChatBot.PushLogin(actor.hotLoginStorage, openwechat.NewRetryLoginOption())
		if err != nil {
			log.Println("登录失败", err)
			return
		}
	}
	err := actor.weChatBot.Block()
	if err != nil {
		log.Println("WeChatBot error:", err)
	}
}

func (actor *Actor) makeMessageHandler() openwechat.MessageHandler {
	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		if msg.IsSendByGroup() {
			msg.ReplyText(actor.chat.RawChat(msg.Content))
			return
		}
		msgContent := msg.Content
		if !GroupTextMsgMatch(msgContent, actor.workKeyInGroup) {
			return
		}
		groupMsgToBot := GetGroupTextMsg(msgContent, actor.workKeyInGroup)
		if len(groupMsgToBot) == 0 {
			msg.ReplyText("如果需要帮助请@我并告诉我需要做什么")
			return
		}
		//msg.ReplyText(groupMsgToBot)
		msg.ReplyText(actor.chat.RawChat(groupMsgToBot))
	})
	return dispatcher.AsMessageHandler()
}

func GroupTextMsgMatch(msgContent string, botKey string) bool {
	fmt.Println("Text: ", msgContent)
	return strings.HasPrefix(msgContent, botKey) || strings.HasSuffix(msgContent, botKey)
}

func GetGroupTextMsg(msgContent string, botKey string) string {
	gptMsg := ""
	if strings.HasPrefix(msgContent, botKey) {
		gptMsg, _ = strings.CutPrefix(msgContent, botKey)
	} else if strings.HasSuffix(msgContent, botKey) {
		gptMsg, _ = strings.CutSuffix(msgContent, botKey)
	}
	return gptMsg
}

func makeStorage(filename string) io.ReadWriteCloser {
	hotLoginStorage := openwechat.NewFileHotReloadStorage(filename)
	defer func(hotLoginStorage io.ReadWriteCloser) {
		err := hotLoginStorage.Close()
		if err != nil {
			log.Println("HotLoginStorage closed error:", err)
		}
	}(hotLoginStorage)
	return hotLoginStorage
}

func workbenchHandlerRegister(workbench *openwechat.Bot, messageHandler openwechat.MessageHandler) {
	workbench.MessageHandler = messageHandler
	workbench.UUIDCallback = func(uuid string) {
		qrcodeUrl := openwechat.GetQrcodeUrl(uuid)
		println("访问下面网址扫描二维码登录\n" + qrcodeUrl)
	}
}
