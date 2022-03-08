package biz

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func ArrayIndexOf(obj interface{}, target interface{}) int {
	t := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < t.Len(); i++ {
			if t.Index(i).Interface() == obj {
				return i
			}
		}
	}
	return -1
}


func initWatchingConfigs() {
	watcherConfigs = &WatcherConfigs{}
	raw, err := os.OpenFile("./configs.json", os.O_RDONLY, 0777)
	if err == nil {
		err = json.NewDecoder(raw).Decode(watcherConfigs)
		if err != nil {
			watcherConfigs = &WatcherConfigs{}
		}
	}
}

func SaveWatcherConfig() {
	file, err := os.OpenFile("./configs.json", os.O_RDWR | os.O_CREATE, 0755)
	if err != nil {
		log.Printf("保存配置信息失败，请使用'save'命令重试。错误：%s\n", err.Error())
		return
	} else {
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(watcherConfigs)
		if err != nil {
			log.Printf("保存配置信息失败，请使用'save'命令重试。错误：%s\n", err.Error())
			return
		}
	}
}