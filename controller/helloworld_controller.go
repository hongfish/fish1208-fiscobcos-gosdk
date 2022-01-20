package controller

import (
	"encoding/json"
	"fish1208-fiscobcos-gosdk/bcos"
	"fish1208-fiscobcos-gosdk/common/response"
	"fish1208-fiscobcos-gosdk/config"
	"fish1208-fiscobcos-gosdk/contract"
	"fish1208-fiscobcos-gosdk/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

var helloWorldContractAddress = config.VConfig.GetString("HelloWorldContractAddress")

func (ctr *Controller) HelloWorldGet(ctx *gin.Context) {
	chainClient := bcos.GetChainClient()
	contractAddress := common.HexToAddress(helloWorldContractAddress)
	instance, err := helloworld.NewHelloWorld(contractAddress, chainClient)
	if err != nil {
		log.Fatal(err)
	}
	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *chainClient.GetCallOpts(), TransactOpts: *chainClient.GetTransactOpts()}
	value, err := helloworldSession.Get()
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx, "执行helloworld.get()失败！")
	}
	response.ResultData(ctx, value)
}

func (ctr *Controller) HelloWorldSet(ctx *gin.Context) {
	hello := new(entity.Hello)
	if err := ctx.ShouldBindJSON(hello); err != nil {
		response.ResultFail(ctx, "参数格式不对！")
		return
	}
	contractAddress := common.HexToAddress(helloWorldContractAddress)
	chainClient := bcos.GetChainClient()
	instance, err := helloworld.NewHelloWorld(contractAddress, chainClient)
	if err != nil {
		log.Fatal(err)
	}
	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *chainClient.GetCallOpts(), TransactOpts: *chainClient.GetTransactOpts()}
	tx, _, err := helloworldSession.Set(hello.V)
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx, "执行helloworld.set()失败！")
	}

	go addBlockTran(ctr, hello, tx.Hash().Hex())
	response.ResultData(ctx, tx)
}

func addBlockTran(ctr *Controller, hello *entity.Hello, hash string) {
	blockTran := new(entity.TxBlockTran)
	b, err1 := json.Marshal(hello)
	if err1 == nil {
		blockTran.Params = string(b)
	}
	blockTran.Contract = "HelloWorld"
	blockTran.Langue = CONTRACTLANGUE
	blockTran.Hash = hash
	blockTran.Status = 1
	ctr.blockTranService.Add(blockTran)
}
