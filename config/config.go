package config

import (
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"os"
)

var Config *viper.Viper

func init() {
	//监听改变动态跟新配置
	go watchConfig()
	//加载配置
	loadConfig()
}

//监听配置改变
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", e.Name)
		//改变重新加载
		loadConfig()
	})
}

//加载配置
func loadConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Fatal error config file: %s \n", err)
		os.Exit(-1)
	}
	//全局配置
	Config = viper.GetViper()
	log.Infof("%v", Config.AllSettings())
}