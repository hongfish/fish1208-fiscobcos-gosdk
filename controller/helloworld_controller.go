package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"
	"fish1208-fiscobcos-gosdk/contract"
	"fish1208-fiscobcos-gosdk/bcos"
	"github.com/labstack/gommon/log"
	"fish1208-fiscobcos-gosdk/common/response"
	"fish1208-fiscobcos-gosdk/entity"
	"fish1208-fiscobcos-gosdk/config"
)

var helloWorldContractAddress = config.Config.GetString("HelloWorldContractAddress")

func (ctr *Controller) HelloWorldGet(ctx *gin.Context){
	client := bcos.ChainClient
	contractAddress := common.HexToAddress(helloWorldContractAddress)
	instance, err := helloworld.NewHelloWorld(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	value, err := helloworldSession.Get()
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx,"执行helloworld.get()失败！")
	}
	response.ResultData(ctx, value)
}

func (ctr *Controller) HelloWorldSet(ctx *gin.Context){
	hello := new(entity.Hello)
	if err := ctx.ShouldBindJSON(hello); err != nil {
		response.ResultFail(ctx, "参数格式不对！")
		return
	}

	client := bcos.ChainClient
	contractAddress := common.HexToAddress(helloWorldContractAddress)
	instance, err := helloworld.NewHelloWorld(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	tx, _, err := helloworldSession.Set(hello.V)
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx,"执行helloworld.set()失败！")
	}

	response.ResultData(ctx, tx)
}