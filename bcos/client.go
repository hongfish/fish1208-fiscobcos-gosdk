package bcos

import (
	"fish1208-fiscobcos-gosdk/config"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/labstack/gommon/log"
)

func GetChainClient() *client.Client {
	configs, err := conf.ParseConfigFile(config.VConfig.GetString("BCOSConfig"))
	if err != nil {
		log.Fatal(err)
	}
	chainClient, err1 := client.Dial(&configs[0])

	if err1 != nil {
		log.Fatal(err1)
		return nil
	}
	return chainClient
}
