package main

import (
	"fmt"
	"live_room_watcher/biz"
)

func main() {
	fmt.Println("===欢迎使用房管助手===")
	fmt.Println("监听进入直播间用户")
	fmt.Println("")
	fmt.Println("输入'help'查看相关命令")
	fmt.Println("====================")
	biz.InitClient()
}
