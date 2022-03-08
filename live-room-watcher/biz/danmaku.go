package biz

import (
	"live_room_watcher/libs/gobilibili"
	"log"
)

func NewDanmakuServer() *gobilibili.BiliBiliClient {
	BiliClient := gobilibili.NewBiliBiliClient()
	BiliClient.RegHandleFunc(gobilibili.CmdWelcome, func(c *gobilibili.Context) bool {
		info := c.GetWelcomeInfo()
		if ArrayIndexOf(info.UID, watcherConfigs.VIPUids) > -1 {
			log.Printf("发现 %s(%d) 进入了房间\r\n", info.Uname, info.UID)
		}
		return false
	})
	BiliClient.RegHandleFunc(gobilibili.CmdDanmuMsg, func(c *gobilibili.Context) bool {
		dinfo := c.GetDanmuInfo()
		if ArrayIndexOf(dinfo.UID, watcherConfigs.VIPUids) > -1 {
			log.Printf("%s[%d] 说: %s\r\n", dinfo.Uname, dinfo.UID, dinfo.Text)
		}
		return false
	})
	return BiliClient
}