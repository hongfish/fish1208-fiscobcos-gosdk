package controller

import (
	"fish1208-fiscobcos-gosdk/service"
)

const CONTRACTLANGUE string = "GoLang"

type Controller struct {
	blockLogService  *service.BlockLogService
	blockTranService *service.BlockTranService
}

func NewController(blockLogService *service.BlockLogService, blockTranService *service.BlockTranService) *Controller {
	return &Controller{
		blockLogService:  blockLogService,
		blockTranService: blockTranService,
	}
}
