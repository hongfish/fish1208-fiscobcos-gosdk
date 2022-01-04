package bcos

import (
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/labstack/gommon/log"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

//定义全局的client
var ChainClient *client.Client

func init(){
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	ChainClient, err = client.Dial(config)

	if err != nil {
		log.Fatal(err)
	}
}
