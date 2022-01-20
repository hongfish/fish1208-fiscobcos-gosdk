package main

import (
	"fish1208-fiscobcos-gosdk/common/xorm"
	"fish1208-fiscobcos-gosdk/config"
	"fish1208-fiscobcos-gosdk/controller"
	"fish1208-fiscobcos-gosdk/service"
	"github.com/gin-gonic/gin"
)

func main() {
	dbengine := xorm.GetEngine(config.VConfig.GetString("DBConfig"))
	controller := controller.NewController(
		service.NewBlockLogService(dbengine),
		service.NewBlockTranService(dbengine),
	)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/chain/getBlockByNumber", controller.GetBlockByNumber)
		api.GET("/chain/getTransactionByHash", controller.GetTransactionByHash)
		api.GET("/helloWorld/get", controller.HelloWorldGet)
		api.POST("/helloWorld/set", controller.HelloWorldSet)
	}

	//router.Run(":8022")
	router.Run(":" + config.VConfig.GetString("ServerPort"))
}
