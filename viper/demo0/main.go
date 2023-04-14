package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	common "godemo/viper"
	"time"
)

var (
	config *common.Config
)

func init() {
	initConfig()
}

func main() {
	time.Sleep(1000 * time.Second)
}

func initConfig() {
	var (
		v *viper.Viper
	)
	config = &common.Config{}
	v = viper.New()
	v.SetConfigFile("./viper/config.yaml")
	//读取配置文件
	readConfig(v)
	//监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		//编辑器可能会触发两次事件
		fmt.Printf("监听到文件变化：%s：%s", e.Name, e.Op)
		//读取配置文件
		readConfig(v)
		//do something
		fmt.Println("do something")
	})
}

//读取配置文件
func readConfig(v *viper.Viper) {
	var (
		err error
	)
	//读取配置文件
	if err = v.ReadInConfig(); err != nil {
		panic(err)
	}
	//解析配置文件
	if err = v.Unmarshal(config); err != nil {
		panic(err)
	}
	fmt.Printf("config:%#v", config)
}
