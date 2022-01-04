package controller

import (
	"github.com/gin-gonic/gin"
	"fish1208-fiscobcos-gosdk/bcos"
	"fish1208-fiscobcos-gosdk/common/response"
	"context"
	"github.com/labstack/gommon/log"
	"strconv"
	"github.com/ethereum/go-ethereum/common"
)

func (ctr *Controller) GetBlockByNumber(ctx *gin.Context){

	blockNumber := ctx.Query("blockNumber")
	number, err := strconv.ParseInt(blockNumber, 10, 64)
	if err != nil{
		log.Fatal(err,"请传入数值！")
		return
	}
	c, cancel := context.WithCancel(context.Background())
	b, err := bcos.ChainClient.GetBlockByNumber(c, number, false )
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx,"获取区块信息失败！")
	}
	cancel()
	response.ResultSuccess(ctx, string(b),"获取区块信息成功！")
}

func (ctr *Controller) GetTransactionByHash(ctx *gin.Context){

	hash := ctx.Query("txHash")
	txHash := common.Hash{}
	txHash.Scan(hash)
	c, cancel := context.WithCancel(context.Background())
	b, err := bcos.ChainClient.GetTransactionByHash(c, txHash)
	if err != nil {
		log.Fatal(err)
		response.ResultFail(ctx,"获取交易信息失败！")
	}
	cancel()
	response.ResultSuccess(ctx, string(b),"获取交易信息成功！")
}