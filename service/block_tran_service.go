package service

import (
	"fish1208-fiscobcos-gosdk/entity"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
)

type BlockTranService struct {
	DbEngine *xorm.Engine
}

func (bt *BlockTranService) Add(blockTran *entity.TxBlockTran) (bool, string) {
	i, err := bt.DbEngine.Insert(blockTran)
	if err != nil {
		log.Errorf("BlockTranService Add error: s%\n", err)
	}
	if i > 0 {
		return true, "新增成功！"
	}
	return false, "新增失败！"
}

func NewBlockTranService(engine *xorm.Engine) *BlockTranService {
	return &BlockTranService{
		DbEngine: engine,
	}
}
