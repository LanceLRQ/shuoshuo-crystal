package biz

import (
	"fmt"
	"github.com/peterh/liner"
	"live_room_watcher/libs/gobilibili"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	cmdHistoryFilePath = filepath.Join(os.TempDir(), ".live_room_watcher_cmd_history")
	bilibiliClient *gobilibili.BiliBiliClient
	exitDanmakuFlag = false
	commandNames = []string{"conn ", "watch ", "unwatch ", "exit", "stop", "help"}
	watcherConfigs *WatcherConfigs
)

func connectDanmaku(roomId int)  {
	for !exitDanmakuFlag {
		err := bilibiliClient.ConnectServer(roomId)
		if err != nil {
			fmt.Printf("[danmaku] Error: %s\n", err.Error())
		}
		// 等待3秒后重连
		time.Sleep(3000)
	}
	exitDanmakuFlag = false
}

func CommandRouter(line *liner.State, cmd string, params []string)  {
	switch cmd {
		case "conn":
			if bilibiliClient != nil && bilibiliClient.IsConnected() {
				log.Println("房间监控正在运行中，请结束后重试！")
				return
			}
			if len(params) < 1 {
				if watcherConfigs.RoomId <= 0 {
					log.Println("请输入房间号")
					return
				}
			}
			roomId := 0
			if len(params) < 1 || params[0] == "" {
				roomId = watcherConfigs.RoomId
			} else {
				rid, err := strconv.ParseInt(params[0], 10, 32)
				if err != nil {
					log.Println("房间号格式错误，必须是一个数字！")
					return
				}
				roomId = int(rid)
			}
			bilibiliClient = NewDanmakuServer()
			go connectDanmaku(roomId)
			wcnt := 0
			for !bilibiliClient.IsConnected() && wcnt < 100 {
				time.Sleep(100 * time.Millisecond)
				wcnt++
			}
			log.Printf("连接成功，当前房间[%d]", roomId)
			watcherConfigs.RoomId = roomId
			SaveWatcherConfig()
			return
		case "stop":
			if bilibiliClient == nil {
				log.Println("房间监控未运行。")
				return
			}
			exitDanmakuFlag = true
			bilibiliClient.Close()
			wcnt := 0
			for bilibiliClient.IsConnected() && wcnt < 100 {
				time.Sleep(100 * time.Millisecond)
				wcnt++
			}
			bilibiliClient = nil
			log.Println("连接已关闭")
			return
		case "watch":
			if len(params) < 1 {
				log.Println("请输入B站UID")
				return
			}
			uid, err := strconv.ParseInt(params[0], 10, 32)
			if err != nil {
				log.Println("UID格式错误，必须是一个数字！")
				return
			}
			if ArrayIndexOf(int(uid), watcherConfigs.VIPUids) > -1 {
				log.Println("UID已存在")
				return
			}
			watcherConfigs.VIPUids = append(watcherConfigs.VIPUids, int(uid))
			log.Printf("添加成功，当前有%d个被监听用户。", len(watcherConfigs.VIPUids))
			SaveWatcherConfig()
			break
		case "unwatch":
			if len(params) < 1 {
				log.Println("请输入B站UID")
				return
			}
			uid, err := strconv.ParseInt(params[0], 10, 32)
			if err != nil {
				log.Println("UID格式错误，必须是一个数字！")
				return
			}
			index := ArrayIndexOf(int(uid), watcherConfigs.VIPUids)
			if index == -1 {
				log.Println("UID不存在")
				return
			}
			watcherConfigs.VIPUids = append(watcherConfigs.VIPUids[0:index], watcherConfigs.VIPUids[index+1:]...)
			log.Printf("移除成功，当前有%d个被监听用户。", len(watcherConfigs.VIPUids))
			SaveWatcherConfig()
			break
		case "exit":
			commonExit(line)
			break
		case "save":
			SaveWatcherConfig()
			break
		case "":
			return
		case "help":
			fmt.Println("支持以下命令：")
			fmt.Println("conn [房间号]\t建立和B站直播服务器的连接并开始工作，如果不输入房间号则会使用上次用过的。")
			fmt.Println("stop\t\t停止监听服务。")
			fmt.Println("watch [uid]\t监听指定的用户进入直播的消息，UID为B站用户UID。")
			fmt.Println("unwatch [uid]\t取消监听指定的用户进入直播的消息，UID为B站用户UID。")
			fmt.Println("exit\t\t退出整个程序。")
			return
		default:
			log.Println("不支持的命令")
	}

}


func commonExit(line *liner.State)  {
	if f, err := os.Create(cmdHistoryFilePath); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	os.Exit(0)
}


func InitClient() {
	initWatchingConfigs()

	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commandNames {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(cmdHistoryFilePath); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		if cmdRaw, err := line.Prompt("> "); err == nil {
			commands := strings.Split(cmdRaw, " ")
			if len(commands) < 1 {
				log.Println("请输入正确的命令")
				continue
			}
			mainCmd := commands[0]
			CommandRouter(line, mainCmd, commands[1:])
			line.AppendHistory(cmdRaw)
		} else if err == liner.ErrPromptAborted {
			commonExit(line)
		}  else {
			commonExit(line)
		}
	}
}
