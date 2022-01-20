package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Xorm struct {
	Config *MysqlConfig `yaml:"xorm"`
}

type MysqlConfig struct {
	Drivename string `yaml:"drivename"`
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Showsql   bool   `yaml:"showsql"`
	Maxidle   int    `yaml:"maxidle"`
	Maxopen   int    `yaml:"maxopen"`
}

func newXorm() *Xorm {
	return &Xorm{Config: &MysqlConfig{}}
}

func loadConfig(file string) *MysqlConfig {
	cfg, err := ioutil.ReadFile(file)
	if err != nil {
		log.Errorf("load dbconfig error: %s \n", err)
	}
	var xorm = newXorm()
	err = yaml.Unmarshal(cfg, xorm)
	if err != nil {
		log.Errorf("yaml Unmarshal error: %s \n", err)
	}
	return xorm.Config
}

func GetEngine(configFile string) *xorm.Engine {
	config := loadConfig(configFile)
	conn := config.User + ":" + config.Password + "@tcp(" + config.Ip + ":" + config.Port + ")/" + config.Database + "?charset=utf8"
	engine, err := xorm.NewEngine(config.Drivename, conn)
	if err != nil {
		log.Errorf("xorm NewEngine error: %s \n", err)
	}
	//打印sql
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(config.Maxidle)
	engine.SetMaxOpenConns(config.Maxopen)
	//连接生存时间半个小时
	engine.SetConnMaxLifetime(1800 * time.Second)
	return engine
}
