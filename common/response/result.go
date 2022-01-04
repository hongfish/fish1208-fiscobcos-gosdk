package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success   = 200
	Failure = 400
)

type RespData struct {
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func ResultMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": Success, "data": nil, "msg": msg})
}

func ResultData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": Success, "data": data, "msg":""})
}

func ResultFail(ctx *gin.Context, msg string){
	ctx.JSON(http.StatusOK, gin.H{"code": Failure, "data": nil, "msg": msg})
}

func ResultSuccess(ctx *gin.Context, data interface{}, msg string){
	ctx.JSON(http.StatusOK, gin.H{"code": Success, "data": data, "msg": msg})
}

