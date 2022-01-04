package main

import (
	"github.com/gin-gonic/gin"
	"fish1208-fiscobcos-gosdk/controller"
	"fish1208-fiscobcos-gosdk/config"
)

func main() {

	controller := &controller.Controller{}

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/chain/getBlockByNumber", controller.GetBlockByNumber)
		api.GET("/chain/getTransactionByHash", controller.GetTransactionByHash)
		api.GET("/helloWorld/get", controller.HelloWorldGet)
		api.POST("/helloWorld/set", controller.HelloWorldSet)
	}

	//router.Run(":8022")
	router.Run(":" + config.Config.GetString("ServerPort"))
}

