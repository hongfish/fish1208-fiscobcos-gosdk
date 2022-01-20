package service

import (
	"fish1208-fiscobcos-gosdk/common/gintool"
	"fish1208-fiscobcos-gosdk/entity"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
)

type BlockLogService struct {
	DbEngine *xorm.Engine
}

func (bl *BlockLogService) Add(blockLog *entity.BlockLog) (bool, string) {
	i, err := bl.DbEngine.Insert(blockLog)
	if err != nil {
		log.Errorf("BlockLogService Add error: s%\n", err)
	}
	if i > 0 {
		return true, "新增成功！"
	}
	return false, "新增失败！"
}

func (bl *BlockLogService) Edit(blockLog *entity.BlockLog) (bool, string) {
	i, err := bl.DbEngine.Id(blockLog.Id).Update(blockLog)
	if err != nil {
		log.Errorf("BlockLogService Edit error: s%\n", err)
	}
	if i > 0 {
		return true, "更新成功！"
	}
	return false, "更新失败！"
}

func (bl *BlockLogService) Del(blockLog *entity.BlockLog) (bool, string) {
	i, err := bl.DbEngine.Where("id = ?", blockLog.Id).Delete(&entity.BlockLog{})
	if err != nil {
		log.Errorf("BlockLogService Del error: s%\n", err)
	}
	if i > 0 {
		return true, "删除成功！"
	}
	return false, "删除失败！"
}

func (bl *BlockLogService) Get(blockLog *entity.BlockLog) (bool, *entity.BlockLog) {
	has, err := bl.DbEngine.Where("id = ?", blockLog.Id).Get(blockLog)
	if err != nil {
		log.Errorf("BlockLogService Get error: s%\n", err)
	}
	return has, blockLog
}

func (bl *BlockLogService) Page(blockLog *entity.BlockLog, page, size int) (bool, []*entity.BlockLog, int64) {
	pager := gintool.CreatePager(page, size)
	blockLogs := make([]*entity.BlockLog, 0)
	values := make([]interface{}, 0)
	where := "1 = 1"
	if blockLog.ModuleName != "" {
		where += " and module_name like ?"
		values = append(values, "%"+blockLog.ModuleName+"%")
	}
	err := bl.DbEngine.Where(where, values...).Limit(pager.PageSize, pager.NumStart).Find(&blockLogs)
	if err != nil {
		log.Errorf("BlockLogService Page error: s%\n", err)
	}
	total, err := bl.DbEngine.Where(where, values...).Count(new(entity.BlockLog))
	if err != nil {
		log.Errorf("BlockLogService Count error: s%\n", err)
	}
	return true, blockLogs, total
}

func (bl *BlockLogService) List(blockLog *entity.BlockLog) (bool, []*entity.BlockLog) {
	blockLogs := make([]*entity.BlockLog, 0)
	values := make([]interface{}, 0)
	where := "1 = 1"
	if blockLog.ModuleName != "" {
		where += " and module_name like ?"
		values = append(values, "%"+blockLog.ModuleName+"%")
	}
	err := bl.DbEngine.Where(where, values...).Find(&blockLogs)
	if err != nil {
		log.Errorf("BlockLogService List error: s%\n", err)
	}
	return true, blockLogs
}

func NewBlockLogService(engine *xorm.Engine) *BlockLogService {
	return &BlockLogService{
		DbEngine: engine,
	}
}
